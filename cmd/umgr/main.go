//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"os"

	"github.com/Urethramancer/signor/opt"
)

var o struct {
	opt.DefaultHelp
	Init   InitCmd   `command:"init" help:"First-time setup." group:"Setup commands"`
	Import ImportCmd `command:"import" help:"Import user, group, site, permission and membership settings from a file." group:"Setup commands"`

	ListUsers   ListUsersCmd   `command:"listusers" help:"List all users." aliases:"lu" group:"User commands"`
	CreateUser  CreateUserCmd  `command:"createuser" help:"Create a new user." aliases:"cu" group:"User commands"`
	DeleteUser  DeleteUserCmd  `command:"deleteuser" help:"Delete a user." aliases:"du" group:"User commands"`
	SetPassword SetPasswordCmd `command:"setpasswd" help:"Set the password for a user." aliases:"sup,setpw" group:"User commands"`

	ListSites  ListSitesCmd  `command:"listsites" help:"List all sites." aliases:"ls" group:"Site commands"`
	CreateSite CreateSiteCmd `command:"createsite" help:"Create a new site." aliases:"cs" group:"Site commands"`
	DeleteSite DeleteSiteCmd `command:"deletesite" help:"Delete a site." aliases:"ds" group:"Site commands"`

	AddSiteMember    AddSiteMemberCmd    `command:"addsitemember" help:"Add a user to a site." aliases:"asm" group:"Site member commands"`
	RemoveSiteMember RemoveSiteMemberCmd `command:"removesitemember" help:"Remove a user from a site." aliases:"rsm" group:"Site member commands"`
	ListSiteMembers  ListSiteMembersCmd  `command:"listsitemembers" help:"List all users in a site." aliases:"lsm" group:"Site member commands"`
	SetSideAdmin     SetSiteAdminCmd     `command:"setsideadmin" help:"Set a user's admin status on a site." aliases:"ssa" group:"Site commands"`

	ListGroups  ListGroupsCmd  `command:"listgroups" help:"List all groups." aliases:"lg" group:"Group commands"`
	CreateGroup CreateGroupCmd `command:"creategroup" help:"Create a new group." aliases:"cg" group:"Group commands"`
	DeleteGroup DeleteGroupCmd `command:"deletegroup" help:"Delete a group." aliases:"dg" group:"Group commands"`
	// AddGroupPermission AddGroupPermissionCmd `command:"addgrouppermission" help:"Add a permission to a group." aliases:"agp" group:"Group commands"`
	// RemoveGroupPermission RemoveGroupPermissionCmd `command:"removegrouppermission" help:"Remove a permission from a group." aliases:"rgp" group:"Group commands"`
	// ListGroupPermissions ListGroupPermissionsCmd `command:"listgrouppermissions" help:"List all permissions of a group." aliases:"lgp" group:"Group commands"`

	AddGroupMember    AddGroupMemberCmd    `command:"addgroupmember" help:"Add a user to a group." aliases:"agm" group:"Group member commands"`
	RemoveGroupMember RemoveGroupMemberCmd `command:"removegroupmember" help:"Remove a user from a group." aliases:"rgm" group:"Group member commands"`
	ListGroupMembers  ListGroupMembersCmd  `command:"listgroupmembers" help:"List all users in a group." aliases:"lgm" group:"Group member commands"`

	ListPerms  ListPermsCmd  `command:"listperms" help:"List all permissions." aliases:"lp" group:"Permission commands"`
	CreatePerm CreatePermCmd `command:"createperm" help:"Create a new permission." aliases:"cp" group:"Permission commands"`
	DeletePerm DeletePermCmd `command:"deleteperm" help:"Delete a permission." aliases:"dp" group:"Permission commands"`
}

func main() {
	a := opt.Parse(&o)
	if o.Help {
		a.Usage()
		return
	}

	err := a.RunCommand(false)
	if err != nil {
		if err == opt.ErrNoCommand {
			a.Usage()
			return
		}

		pr("Error running command: %s", err.Error())
		os.Exit(2)
	}
}

func pr(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
}
