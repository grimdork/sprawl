package main

import (
	"encoding/json"
	"net/http"
)

func (srv *Server) createPermission(w http.ResponseWriter, r *http.Request) {
	err := srv.CreatePermission(
		r.Header.Get("name"),
		r.Header.Get("description"),
	)

	if err != nil {
		srv.E("Failed to create permission: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (srv *Server) deletePermission(w http.ResponseWriter, r *http.Request) {
	err := srv.DeletePermission(r.Header.Get("name"))
	if err != nil {
		println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (srv *Server) listPermissions(w http.ResponseWriter, r *http.Request) {
	list, err := srv.GetPermissions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(list)
	if err != nil {
		srv.E("Failed to marshal permissions: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}
