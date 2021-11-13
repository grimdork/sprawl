//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package sprawl

const (
	// EPAuth authenticates a user and returns a token.
	EPAuth = "/auth"

	// EPUsers is the bulk user endpoint.
	// GET lists all users.
	// POST creates multiple users.
	// PUT updates multiple users.
	// DELETE deletes multiple users.
	EPUsers = "/users"

	// EPUser is the single user endpoint.
	// GET returns a single user.
	// POST creates a single user.
	// PUT updates a single user.
	// DELETE deletes a single user.
	EPUser = "/user"

	// EPSetPassword sets the password for a user.
	EPSetPassword = "/password"

	// EPSites is the bulk site endpoint.
	// GET lists all sites.
	// POST creates multiple sites.
	// PUT updates multiple sites.
	// DELETE removes multiple sites.
	EPSites = "/sites"

	// EPSite is the single site endpoint.
	// POST creates a single site.
	// PUT updates a single site.
	// DELETE removes a single site.
	EPSite = "/site"

	// EPSetAdmin sets a user as an admin on a site.
	// PUT updates a single member.
	EPSetAdmin = "/setadmin"

	// EPSetProfile updates profike data for a user on a site.
	// PUT updates a single profile.
	EPSetProfile = "/updateprofile"

	//EPGroups is the bulk group endpoint.
	//GET lists all groups.
	//POST creates multiple groups.
	//DELETE removes multiple groups.
	EPGroups = "/groups"

	//EPGroup is the single group endpoint.
	//POST creates a single group.
	//DELETE removes a single group.
	EPGroup = "/group"

	EPListGroupMembers = "/listgroupmembers"
	// EPAddGroupMember adds a user to a group on a site.
	EPAddGroupMember = "/addgroupmember"
	// EPRemoveGroupMember removes a user from a group on a site.
	EPRemoveGroupMember = "/removegroupmember"
	// EPAddGroupPermission adds a permission to a group on a site.
	EPAddGroupPermission = "/addgrouppermission"
	// EPRemoveGroupPermission removes a permission from a group on a site.
	EPRemoveGroupPermission = "/removegrouppermission"
	// EPListGroupPermissions lists all permissions for a group on a site.
	EPListGroupPermissions = "/listgrouppermissions"

	// EPPermissions is the bulk permission endpoint.
	// GET lists all permissions.
	// POST creates multiple permissions.
	// DELETE removes multiple permissions.
	EPPermissions = "/permissions"

	// EPPermission is the single permission endpoint.
	// GET returns a single permission.
	// POST creates a single permission.
	// DELETE removes a single permission.
	EPPermission = "/permission"

	// EPMembers is the bulk member endpoint.
	// GET lists all members.
	// POST creates multiple members.
	// PUT updates multiple members.
	// DELETE removes multiple members.
	EPMembers = "/members"

	// EPMember is the single member endpoint.
	// POST creates a single member.
	// PUT updates a single member.
	// DELETE removes a single member.
	EPMember = "/member"

	// EPAdmin enables and disables admin status for a user on a site.
	// PUT enables admin status.
	// DELETE disables admin status.
	EPAdmin = "/admin"
)

// Request contains the variables passed to endpoints.
type Request map[string]string

// GroupList is returned from the list groups endpoint.
type GroupList struct {
	// Groups on the site.
	Groups []Group `json:"groups"`
}
