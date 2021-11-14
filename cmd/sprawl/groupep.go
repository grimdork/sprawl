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

	list, err := srv.GetGroups(
		start, max,
		r.Header.Get("site"),
	)
	if err != nil {
		srv.E("Failed to get group members: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(list)
	if err != nil {
		srv.E("Failed to marshal users: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (srv *Server) listGroupMembers(w http.ResponseWriter, r *http.Request) {
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

	users, err := srv.GetGroupMembers(
		r.Header.Get("site"),
		r.Header.Get("group"),
	)
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

func (srv *Server) addGroupMember(w http.ResponseWriter, r *http.Request) {
	err := srv.AddGroupMember(
		r.Header.Get("site"),
		r.Header.Get("group"),
		r.Header.Get("name"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (srv *Server) removeGroupMember(w http.ResponseWriter, r *http.Request) {
	err := srv.RemoveGroupMember(
		r.Header.Get("site"),
		r.Header.Get("group"),
		r.Header.Get("name"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (srv *Server) listGroupPermissions(w http.ResponseWriter, r *http.Request) {
	perms, err := srv.GetGroupPermissions(
		r.Header.Get("site"),
		r.Header.Get("group"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(perms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func (srv *Server) addGroupPermission(w http.ResponseWriter, r *http.Request) {
	err := srv.AddGroupPermission(
		r.Header.Get("site"),
		r.Header.Get("group"),
		r.Header.Get("permission"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (srv *Server) removeGroupPermission(w http.ResponseWriter, r *http.Request) {
	err := srv.RemoveGroupPermission(
		r.Header.Get("site"),
		r.Header.Get("group"),
		r.Header.Get("permission"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
