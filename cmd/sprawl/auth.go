// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
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
		t := srv.GenerateToken(username)
		s := fmt.Sprintf("{\"token\":\"%s\"}", t)
		w.Write([]byte(s))
	}
}

// tokencheck middleware to insert before privileged endpoints.
func (srv *Server) tokencheck(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if !srv.VerifyToken(
			r.Header.Get("username"),
			r.Header.Get("token"),
		) {
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
