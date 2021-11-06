// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/grimdork/sprawl"
)

func (srv *Server) createGroup(w http.ResponseWriter, r *http.Request) {
	err := srv.CreateGroup(
		r.Header.Get("name"),
		r.Header.Get("site"),
	)

	if err != nil {
		srv.E("Failed to create group: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (srv *Server) deleteGroup(w http.ResponseWriter, r *http.Request) {
	err := srv.DeleteGroup(r.Header.Get("name"), r.Header.Get("site"))
	if err != nil {
		println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (srv *Server) listGroups(w http.ResponseWriter, r *http.Request) {
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

	list, err := srv.GetGroups(start, max, r.Header.Get("site"))
	if err != nil {
		srv.E("Failed to get groups: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(sprawl.GroupList{Groups: list})
	if err != nil {
		srv.E("Failed to marshal groups: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (srv *Server) listGroupMembers(w http.ResponseWriter, r *http.Request) {
}

func (srv *Server) addGroupMember(w http.ResponseWriter, r *http.Request) {
}
