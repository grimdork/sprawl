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
	EPAuth = "auth"

	// EPListUser lists all users.
	EPListUsers = "/listusers"
	// EPCreateUser creates a new user and sets the password.
	EPCreateUser = "/createuser"
	// EPDeleteUser deletes a user.
	EPDeleteUser = "/deleteuser"
	// EPSetPassword sets the password for a user.
	EPSetPassword = "/setpassword"

	// EPListSites lists all sites.
	EPListSites = "/listsites"
	// EPCreateSite creates a new site.
	EPCreateSite = "/createsite"
	// EPDeleteSite deletes a site.
	EPDeleteSite = "/deletesite"
	// EPListSiteMembers lists all users with profiles on a site.
	EPListSiteMembers = "/listsitemembers"

	// EPCreateProfile creates a new profile for a user, incidentally making the user a member.
	EPCreateProfile = "/createprofile"
	// EPDeleteProfile deletes a profile, effectively removing a user from a site.
	EPDeleteProfile = "/deleteprofile"
	// EPUpdateProfile updates profike data for a user on a site.
	EPUpdateProfile = "/updateprofile"

	// EPListGroups lists all groups, optionally just for one site.
	EPListGroups = "/listgroups"
	// EPCreateGroup creates a new group on a site.
	EPCreateGroup = "/creategroup"
	// EPDeleteGroup deletes a group from a site.
	EPDeleteGroup = "/deletegroup"
	// EPListGroupMembers lists all users in a site's group.
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

	// EPListPermissions lists all permission keywords.
	EPListPermissions = "/listpermissions"
	// EPCreatePermission creates a new permission keyword.
	EPCreatePermission = "/createpermission"
	// EPDeletePermission deletes a permission keyword.
	EPDeletePermission = "/deletepermission"
)

// Request contains the variables passed to endpoints.
type Request map[string]string

// GroupList is returned from the list groups endpoint.
type GroupList struct {
	// Groups on the site.
	Groups []Group `json:"groups"`
}
