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
	"os"
	"strconv"
	"time"

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

func getenv(key, alt string) string {
	s := os.Getenv(key)
	if s == "" {
		return alt
	}

	return s
}

// NewServer creates a web server and background tasks.
func NewServer() (*Server, error) {
	c := getenv("PW_ITERATIONS", "10")
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
	srv.Database, err = sprawl.NewDatabase(getenv("DATABASE_URL", "localhost"))
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
	err = srv.InitDatabase(getenv("ADMIN_PASSWORD", "potrzebie"))
	if err != nil {
		return nil, err
	}

	_, err = srv.GetUser("admin")
	if err != nil {
		return nil, err
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
			r.Put("/", srv.updateUser)
			r.Get("/", srv.getUser)
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

			// Toggle admin status.
			r.Put(sprawl.EPAdmin, srv.setSiteAdmin)

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

		// Group operations.
		r.Route(sprawl.EPGroup, func(r chi.Router) {
			r.Use(
				srv.tokencheck,
				srv.siteadmincheck,
			)

			r.Get(sprawl.EPMembers, srv.listGroupMembers)

			// Single member operations.
			r.Route(sprawl.EPMember, func(r chi.Router) {
				r.Post("/", srv.addGroupMember)
				r.Delete("/", srv.removeGroupMember)
			})

			// Bulk permission operations.
			r.Get(sprawl.EPPermissions, srv.listGroupPermissions)

			// Single permission operations.
			r.Route(sprawl.EPPermission, func(r chi.Router) {
				r.Post("/", srv.addGroupPermission)
				r.Delete("/", srv.removeGroupPermission)
			})

			// Single group operations.
			r.Post("/", srv.createGroup)
			r.Delete("/", srv.deleteGroup)
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
			r.Get("/", srv.getPermission)
			r.Put("/", srv.updatePermission)
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

func truth(s string) bool {
	switch s {
	case "true", "1", "yes", "on":
		return true
	}

	return false
}
