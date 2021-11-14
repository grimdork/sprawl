//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
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

	cfg, err := sprawl.LoadConfig(configPath)
	if err != nil {
		return err
	}

	_, err = cfg.Post(sprawl.EPGroup+sprawl.EPMember, sprawl.Request{
		"site":  cmd.Site,
		"group": cmd.Group,
		"name":  cmd.Name,
	})
	return err
}
