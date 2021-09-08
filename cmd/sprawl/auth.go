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
	u := srv.GetUser(username)
	if u == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if u.CheckPassword(password) {
		t := srv.generateToken(username)
		s := fmt.Sprintf("{\"token\":\"%s\"}", t)
		w.Write([]byte(s))
		return
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
	u := srv.GetUser(name)
	sql := "delete from tokens where uid=$1"
	_, err := srv.Exec(context.Background(), sql, u.ID)
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
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
