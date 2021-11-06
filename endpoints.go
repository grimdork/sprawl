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
	EPListUsers   = "/listusers"
	EPCreateUser  = "/createuser"
	EPDeleteUser  = "/deleteuser"
	EPSetPassword = "/setpassword"

	EPListGroups       = "/listgroups"
	EPCreateGroup      = "/creategroup"
	EPDeleteGroup      = "/deletegroup"
	EPListGroupMembers = "/listgroupmembers"
	EPAddGroupMember   = "/addgroupmember"
	EPRemoveMember     = "/removemember"

	EPListSites       = "/listsites"
	EPCreateSite      = "/createsite"
	EPDeleteSite      = "/deletesite"
	EPListSiteMembers = "/listsitemembers"

	EPListPermissions  = "/listpermissions"
	EPCreatePermission = "/createpermission"
	EPDeletePermission = "/deletepermission"

	EPListRoles  = "/listroles"
	EPCreateRole = "/createrole"
	EPDeleteRole = "/deleterole"
)

// Request contains the variables passed to endpoints.
type Request map[string]string

// GroupList is returned from the list groups endpoint.
type GroupList struct {
	// Groups on the site.
	Groups []Group `json:"groups"`
}
