//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl/client"
)

// CreateGroupCmd options.
type CreateGroupCmd struct {
	opt.DefaultHelp
	Site string `placeholder:"SITE" help:"The site in which to create the group."`
	Name string `placeholder:"NAME" help:"An alphanumeric group name to create."`
}

// Run command.
func (cmd *CreateGroupCmd) Run(args []string) error {
	if cmd.Help || cmd.Site == "" {
		return opt.ErrUsage
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	return c.CreateGroup(cmd.Site, cmd.Name)
}
