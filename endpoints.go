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
)

// CreateRequest is used for create endpoints.
type CreateRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
