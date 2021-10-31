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
	"io"
	"net/http"
	"strconv"

	"github.com/grimdork/sprawl"
)

func (srv *Server) createUser(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		srv.E("Failed to read body: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var req sprawl.CreateRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		srv.E("Failed to unmarshal body: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = srv.CreateUser(req.Name, req.Password)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	srv.L("Added user '%s'", req.Name)
}

func (srv *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		srv.E("Failed to read body: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var req sprawl.CreateRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		srv.E("Failed to unmarshal body: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = srv.DeleteUser(req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (srv *Server) setPassword(w http.ResponseWriter, r *http.Request) {
	u := srv.GetUser(r.Header.Get("username"))
	srv.L("%+v", u)
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
