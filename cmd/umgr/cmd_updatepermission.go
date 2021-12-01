// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"github.com/grimdork/opt"
	"github.com/grimdork/sprawl/client"
)

// UpdatePermCmd options.
type UpdatePermCmd struct {
	opt.DefaultHelp
	Name        string `placeholder:"KEYWORD" help:"Permission to change the description for."`
	Description string `placeholder:"DESCRIPTION" help:"New description."`
}

// Run command.
func (cmd *UpdatePermCmd) Run(args []string) error {
	if cmd.Help || cmd.Name == "" {
		return opt.ErrUsage
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	err = c.UpdatePermission(cmd.Name, cmd.Description)
	return err
}
