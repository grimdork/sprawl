// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
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
		http.Error(w, "User/password combination not found", http.StatusForbidden)
		return
	}

	if u.CheckPassword(password) {
		t := srv.GenerateToken(username)
		w.Write([]byte(fmt.Sprintf(
			"{\"token\":\"%s\"}",
			t,
		)))
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
		if !srv.IsSiteAdmin(username, site) {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
