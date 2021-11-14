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
	"strconv"
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
	pwiter int
}

// NewServer creates a web server and background tasks.
func NewServer() (*Server, error) {
	c := env.Get("PW_ITERATIONS", "10")
	x, _ := strconv.Atoi(c)
	if x < 1 {
		x = 10
	}
	srv := &Server{pwiter: x}
	srv.Init()

	//
	// Database
	//
	var err error
	srv.L("Opening database.")
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

		err = srv.CreateProfile("admin", "system", "", true)
		if err != nil {
			return nil, err
		}
	}

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

		// Bulk user operations.
		r.Route(sprawl.EPUsers, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.admincheck,
			)
			r.Get("/", srv.listUsers)
		})

		// Single user operations.
		r.Route(sprawl.EPUser, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.admincheck,
			)
			r.Post("/", srv.createUser)
			r.Delete("/", srv.deleteUser)
			// r.Post("/", srv.updateUser)
			// r.Get("/", srv.getUser)
			r.Post(sprawl.EPSetPassword, srv.setPassword)
		})

		//
		// Sites
		//

		// Bulk site operations.
		r.Route(sprawl.EPSites, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.admincheck,
			)
			r.Get("/", srv.listSites)
		})

		// Site operations.
		r.Route(sprawl.EPSite, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.admincheck,
			)

			// Bulk member operations.
			r.Get(sprawl.EPMembers, srv.listSiteMembers)

			// Single member operations.
			r.Route(sprawl.EPMember, func(r chi.Router) {
				r.Post("/", srv.addSiteMember)
				r.Put("/", srv.updateSiteMember)
				r.Delete("/", srv.removeSiteMember)
			})

			// Admin operations.
			r.Route(sprawl.EPAdmin, func(r chi.Router) {
				r.Put("/", srv.enableSiteAdmin)
				r.Delete("/", srv.disableSiteAdmin)
			})

			// Single site operations.
			r.Post("/", srv.createSite)
			r.Delete("/", srv.deleteSite)
		})

		//
		// Groups
		//

		// Bulk group operations.
		r.Route(sprawl.EPGroups, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.siteadmincheck,
			)
			r.Get("/", srv.listGroups)
		})

		// Single group operations.
		r.Route(sprawl.EPGroup, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.siteadmincheck,
			)
			r.Post("/", srv.createGroup)
			r.Delete("/", srv.deleteGroup)
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
		// Permissions
		//

		// Bulk permission operations.
		r.Route(sprawl.EPPermissions, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.admincheck,
			)
			r.Get("/", srv.listPermissions)
		})

		// Single permission operations.
		r.Route(sprawl.EPPermission, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.admincheck,
			)
			// r.Get("/", srv.getPermission)
			r.Post("/", srv.createPermission)
			r.Delete("/", srv.deletePermission)
		})

		// Default route for "/".
		r.Get("/", srv.index)
	})

	return srv, nil
}

func (srv *Server) index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("v1"))
}
