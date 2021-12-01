//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"github.com/grimdork/opt"
	"github.com/grimdork/sprawl/client"
)

// DeleteGroupCmd options.
type DeleteGroupCmd struct {
	opt.DefaultHelp
	Site string `placeholder:"SITE" help:"Site of the group."`
	Name string `placeholder:"NAME" help:"Name of group to delete."`
}

// Run command.
func (cmd *DeleteGroupCmd) Run(args []string) error {
	if cmd.Help || cmd.Name == "" {
		return opt.ErrUsage
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	return c.DeleteGroup(cmd.Site, cmd.Name)
}
