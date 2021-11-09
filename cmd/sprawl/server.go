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
	"net/http"
	"time"

	"github.com/Urethramancer/signor/env"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/grimdork/sprawl"
	"github.com/grimdork/sweb"
)

type Server struct {
	sweb.Server
	*sprawl.Database
}

// NewServer creates a web server and background tasks.
func NewServer() (*Server, error) {
	srv := &Server{}
	srv.Init()

	//
	// Endpoints
	//
	srv.Route("/", func(r chi.Router) {
		r.Use(
			middleware.NoCache,
			middleware.RealIP,
			sweb.AddCORS,
			middleware.Timeout(time.Second*10),
			sweb.AddJSONHeaders,
		)
		r.Options("/", sweb.Preflight)
		// Add route for login token.
		r.Post("/auth", srv.auth)

		//
		// Users
		//

		// List users
		r.Route(sprawl.EPListUsers, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.admincheck,
			)
			r.Get("/", srv.listUsers)
		})

		// Create users
		r.Route(sprawl.EPCreateUser, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.admincheck,
			)
			r.Post("/", srv.createUser)
		})

		// Delete users
		r.Route(sprawl.EPDeleteUser, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.admincheck,
			)
			r.Post("/", srv.deleteUser)
		})

		// Set user passwords
		r.Route(sprawl.EPSetPassword, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.siteadmincheck,
			)
			r.Post("/", srv.setPassword)
		})

		//
		// Groups
		//

		// List groups
		r.Route(sprawl.EPListGroups, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.siteadmincheck,
			)
			r.Get("/", srv.listGroups)
		})

		// Create groups
		r.Route(sprawl.EPCreateGroup, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.siteadmincheck,
			)
			r.Post("/", srv.createGroup)
		})

		// Delete groupa
		r.Route(sprawl.EPDeleteGroup, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.siteadmincheck,
			)
			r.Post("/", srv.deleteGroup)
		})

		// List users in groups
		r.Route(sprawl.EPListGroupMembers, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.siteadmincheck,
			)
			r.Get("/", srv.listGroupMembers)
		})

		// Add users to groups
		r.Route(sprawl.EPAddGroupMember, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.siteadmincheck,
			)
			r.Post("/", srv.addGroupMember)
		})

		//
		// Sites
		//

		// List sites (domains)
		r.Route(sprawl.EPListSites, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.admincheck,
			)
			r.Get("/", srv.listSites)
		})

		// Create a new site/domain
		r.Route(sprawl.EPCreateSite, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.admincheck,
			)
			r.Post("/", srv.createSite)
		})

		// Delete a site/domain
		r.Route(sprawl.EPDeleteSite, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.admincheck,
			)
			r.Post("/", srv.deleteSite)
		})

		// List users in all groups in a site/domain
		r.Route(sprawl.EPListSiteMembers, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.siteadmincheck,
			)
			r.Get("/", srv.listSiteMembers)
		})

		//
		// Permissions
		//

		// List permissions
		r.Route(sprawl.EPListPermissions, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.admincheck,
			)
			r.Get("/", srv.listPermissions)
		})

		// Create permission
		r.Route(sprawl.EPCreatePermission, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.admincheck,
			)
			r.Post("/", srv.createPermission)
		})

		// Delete permission
		r.Route(sprawl.EPDeletePermission, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.admincheck,
			)
			r.Post("/", srv.deletePermission)
		})

		// Default route for "/".
		r.Get("/", srv.index)
	})

	//
	// Database
	//
	var err error
	srv.Database, err = sprawl.NewDatabase(env.Get("DATABASE_URL", "localhost"))
	if err != nil {
		return nil, err
	}

	srv.AddStopHook(func() {
		srv.L("Closing database.")
		srv.Database.Close()
	})

	err = srv.CreateTables()
	if err != nil {
		return nil, err
	}

	// Create admin if it doesn't exist.
	u := srv.GetUser("admin")
	if u == nil {
		srv.L("No admin user - creating.")
		err = srv.CreateUser("admin", env.Get("ADMIN_PASSWORD", "potrzebie"))
		if err != nil {
			println("FAIL")
			return nil, err
		}
	}

	_, err = srv.GetSiteID("system")
	if err != nil {
		srv.L("No system site - creating.")
		err = srv.CreateSite("system")
		if err != nil {
			return nil, err
		}
	}

	return srv, nil
}

func (srv *Server) index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("v1"))
}
