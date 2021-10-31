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
	Init InitCmd `command:"init" help:"First-time setup." group:"Setup commands"`

	ListUsers  ListUsersCmd  `command:"listusers" help:"List all users." aliases:"lu" group:"User commands"`
	CreateUser CreateUserCmd `command:"createuser" help:"Create a new user." aliases:"cu" group:"User commands"`
	DeleteUser DeleteUserCmd `command:"deleteuser" help:"Delete a user." aliases:"du" group:"User commands"`
	// SetPassword SetUPasswordCmd `command:"setpasswd" help:"Set a user's password." aliases:"sup" group:"User commands"`

	ListSites  ListSitesCmd  `command:"listsites" help:"List all sites." aliases:"ls" group:"Site commands"`
	CreateSite CreateSiteCmd `command:"createsite" help:"Create a new site." aliases:"cs" group:"Site commands"`
	DeleteSite DeleteSiteCmd `command:"deletesite" help:"Delete a site." aliases:"ds" group:"Site commands"`

	// ListGroups  ListGroupsCmd  `command:"listgroups" help:"List all groups." aliases:"lg" group:"Group commands"`
	CreateGroup CreateGroupCmd `command:"creategroup" help:"Create a new group." aliases:"cg" group:"Group commands"`
	// DeleteGroup DeleteGroupCmd `command:"deletegroup" help:"Delete a group." aliases:"dg" group:"Group commands"`
	// AddMember AddMemberCmd `command:"addmember" help:"Add a user to a group." aliases:"am" group:"Group commands"`
	// RemoveMember RemoveMemberCmd `command:"removemember" help:"Remove a user from a group." aliases:"rm" group:"Group commands"`
	// ListMembers ListMembersCmd `command:"listmembers" help:"List all members of a group." aliases:"lm" group:"Group commands"`
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
