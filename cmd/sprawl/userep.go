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
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	var req sprawl.CreateRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		srv.E("Failed to unmarshal body: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	err = srv.CreateUser(req.Name, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(http.StatusText(http.StatusConflict)))
		return
	}

	w.WriteHeader(http.StatusCreated)
	srv.L("Added user '%s'", req.Name)
}

func (srv *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		srv.E("Failed to read body: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	var req sprawl.CreateRequest
	err = json.Unmarshal(data, &req)
	if err != nil {
		srv.E("Failed to unmarshal body: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	err = srv.DeleteUser(req.Name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
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

	users := srv.GetUsers(start, max)
	data, err := json.Marshal(users)
	if err != nil {
		// Write a not found HTTP error
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
		return
	}

	w.Write(data)
}
