//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (srv *Server) createUser(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("name")
	err := srv.CreateUser(
		name,
		r.Header.Get("password"),
	)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	srv.L("Added user '%s'", name)
}

func (srv *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("name")
	err := srv.DeleteUser(name)
	if err != nil {
		println("Fuck: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	srv.L("Deleted user '%s'", name)
}

func (srv *Server) updateUser(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("name")
	newname := r.Header.Get("newname")
	fullname := r.Header.Get("fullname")
	email := r.Header.Get("email")
	err := srv.UpdateUser(
		name,
		newname,
		fullname,
		email,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (srv *Server) getUser(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("name")
	user, err := srv.GetUser(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (srv *Server) setPassword(w http.ResponseWriter, r *http.Request) {
	u := r.Header.Get("name")
	pw := r.Header.Get("password")
	err := srv.SetPassword(u, pw, srv.pwiter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (srv *Server) listUsers(w http.ResponseWriter, r *http.Request) {
	start, _ := strconv.ParseInt(r.Header.Get("start"), 10, 64)
	max, _ := strconv.ParseInt(r.Header.Get("max"), 10, 64)
	if start < 0 {
		start = 0
	}
	if max < 1 {
		max = 10
	}
	if max > 100 {
		max = 100
	}

	users, err := srv.GetUsers(start, max)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	data, err := json.Marshal(users)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Write(data)
}
