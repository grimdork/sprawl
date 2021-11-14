// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"context"
	"crypto/sha512"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/grimdork/sprawl"
)

func (srv *Server) auth(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("username")
	username = strings.TrimSpace(username)
	password := r.Header.Get("password")
	password = strings.TrimSpace(password)
	u, err := srv.GetUser(username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if u.CheckPassword(password) {
		t := srv.generateToken(username)
		s := fmt.Sprintf("{\"token\":\"%s\"}", t)
		w.Write([]byte(s))
	}
}

func (srv *Server) getToken(username string) string {
	sql := "select hash,expires from tokens inner join users on users.id=tokens.uid where users.name=$1"
	var token string
	var expires int64
	err := srv.QueryRow(context.Background(), sql, username).Scan(&token)
	ex := time.Unix(expires, 0)
	if err != nil && ex.After(time.Now()) {
		return token
	}

	return ""
}

// Generate a new token if needed.
func (srv *Server) generateToken(name string) string {
	token := srv.getToken(name)
	if token != "" {
		return token
	}

	// Delete all expired tokens for this user.
	u, err := srv.GetUser(name)
	if err != nil {
		srv.E("Error getting user: %s", err.Error())
		return ""
	}

	sql := "delete from tokens where uid=$1"
	_, err = srv.Exec(context.Background(), sql, u.ID)
	if err != nil {
		srv.L("Error deleting expired tokens: %s", err)
	}

	// And finally generate a new one.
	s := sprawl.RandString(32)
	h := sha512.New()
	h.Write([]byte(s))
	token = fmt.Sprintf("%x", h.Sum(nil))
	t := time.Now().Add(time.Hour * 8)
	expires := t.Unix()
	_, err = srv.Exec(context.Background(), "insert into tokens (uid,hash,expires) values ($1,$2,$3)", u.ID, token, expires)
	if err != nil {
		srv.E("Insert error on tokens table: %s", err.Error())
		return ""
	}

	return token
}

// Verify that the token is valid.
func (srv *Server) verifyToken(username, token string) bool {
	sql := "select expires from tokens inner join users on users.id=tokens.uid where users.name=$1 and hash=$2"
	var expires int64
	err := srv.QueryRow(context.Background(), sql, username, token).Scan(&expires)
	ex := time.Unix(expires, 0)
	if err != nil || ex.Before(time.Now()) {
		srv.E("Token verification failed: %s", err.Error())
		return false
	}

	return true
}

// tokencheck middleware to insert before privileged endpoints.
func (srv *Server) tokencheck(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if !srv.verifyToken(r.Header.Get("username"), r.Header.Get("token")) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// admincheck middleware to insert before system endpoints.
func (srv *Server) admincheck(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("username") != "admin" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// siteadmincheck middleware to insert before site endpoints.
func (srv *Server) siteadmincheck(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		site := r.Header.Get("site")
		if site == "" {
			h.ServeHTTP(w, r)
			return
		}

		username := r.Header.Get("username")
		if !srv.isSiteAdmin(username, site) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (srv *Server) isSiteAdmin(name, site string) bool {
	if name == "admin" {
		return true
	}

	sql := `select count(profiles.uid) from profiles
		inner join users on profiles.uid=users.id
		inner join sites on profiles.sid=sites.id
		where users.name=$1 and sites.name=$2;`
	var count int64
	err := srv.QueryRow(context.Background(), sql, name, site).Scan(&count)
	if err != nil {
		srv.E("Error checking site admin: %s", err.Error())
		return false
	}

	return count > 0
}
