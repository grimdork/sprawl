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

func (srv *Server) listSites(w http.ResponseWriter, r *http.Request) {
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

	sites, err := srv.GetSites(start, max)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
		return
	}

	data, err := json.Marshal(sites)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(http.StatusText(http.StatusNotFound)))
		return
	}

	w.Write(data)
}

func (srv *Server) createSite(w http.ResponseWriter, r *http.Request) {
	name, err := getName(r)
	if err != nil {
		srv.E("Failed to unmarshal body: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = srv.CreateSite(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (srv *Server) deleteSite(w http.ResponseWriter, r *http.Request) {
	name, err := getName(r)
	if err != nil {
		srv.E("Failed to unmarshal body: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if name == "1" || name == "system" {
		http.Error(w, "Cannot delete system site.", http.StatusInternalServerError)
		return
	}

	err = srv.DeleteSite(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (srv *Server) listSiteMembers(w http.ResponseWriter, r *http.Request) {
}
