//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"github.com/grimdork/opt"
	"github.com/grimdork/sprawl/client"
)

// AddGroupMemberCmd options.
type AddGroupMemberCmd struct {
	opt.DefaultHelp
	Site  string `placeholder:"SITE" help:"Site of the group."`
	Group string `placeholder:"GROUP" help:"Group to add the member to."`
	Name  string `placeholder:"USERNAME" help:"User to add to the group."`
}

// Run command.
func (cmd AddGroupMemberCmd) Run(args []string) error {
	if cmd.Name == "" {
		return opt.ErrUsage
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	return c.AddGroupMember(cmd.Site, cmd.Group, cmd.Name)
}
