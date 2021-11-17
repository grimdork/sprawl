//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl/client"
)

// AddGroupPermissionCmd options.
type AddGroupPermissionCmd struct {
	opt.DefaultHelp
	Site    string `placeholder:"SITE" help:"Site of the group."`
	Group   string `placeholder:"GROUP" help:"Group to add the permission to."`
	Keyword string `placeholder:"KEYWORD" help:"Permission to add to the group."`
}

// Run command.
func (cmd AddGroupPermissionCmd) Run(args []string) error {
	if cmd.Keyword == "" {
		return opt.ErrUsage
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	return c.AddGroupPermission(cmd.Site, cmd.Group, cmd.Keyword)
}
