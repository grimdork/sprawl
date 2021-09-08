//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import "net/http"

func (srv *Server) listSites(w http.ResponseWriter, r *http.Request) {

}

func (srv *Server) createSite(w http.ResponseWriter, r *http.Request) {
	srv.AddSite(r.Header.Get("site"))
}

func (srv *Server) deleteSite(w http.ResponseWriter, r *http.Request) {
	srv.RemoveSite(r.Header.Get("site"))
}

func (srv *Server) listSiteMembers(w http.ResponseWriter, r *http.Request) {

}
