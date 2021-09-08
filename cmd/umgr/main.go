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
	Init InitCmd `command:"init" help:"First-time setup."`

	ListUsers  ListUsersCmd  `command:"listusers" help:"List all users." aliases:"lu"`
	CreateUser CreateUserCmd `command:"createuser" help:"Create a new user." aliases:"cu"`
	DeleteUser DeleteUserCmd `command:"deleteuser" help:"Delete a user." aliases:"du"`
	// SetPassword SetUPasswordCmd `command:"setpasswd" help:"Set a user's password." aliases:"sup"`

	// ListSites ListSitesCmd `command:"listsites" help:"List all sites." aliases:"ls"`
	// CreateSite CreateSiteCmd `command:"createsite" help:"Create a new site." aliases:"cs"`
	// DeleteSite DeleteSiteCmd `command:"deletesite" help:"Delete a site." aliases:"ds"`
	// AddToSite AddToSiteCmd `command:"addtosite" help:"Add a user to a site." aliases:"at"`
	// RemoveFromSite RemoveFromSiteCmd `command:"removefromsite" help:"Remove a user from a site." aliases:"rfs"`

	// ListGroups  ListGroupsCmd  `command:"listgroups" help:"List all groups." aliases:"lg"`
	CreateGroup CreateGroupCmd `command:"creategroup" help:"Create a new group." aliases:"cg"`
	// DeleteGroup DeleteGroupCmd `command:"deletegroup" help:"Delete a group." aliases:"dg"`
	// AddMember AddMemberCmd `command:"addmember" help:"Add a user to a group." aliases:"am"`
	// RemoveMember RemoveMemberCmd `command:"removemember" help:"Remove a user from a group." aliases:"rm"`
	// ListMembers ListMembersCmd `command:"listmembers" help:"List all members of a group." aliases:"lm"`
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
