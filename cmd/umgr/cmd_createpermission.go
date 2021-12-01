// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"github.com/grimdork/opt"
	"github.com/grimdork/sprawl/client"
)

// CreatePermCmd options.
type CreatePermCmd struct {
	opt.DefaultHelp
	Name        string `placeholder:"KEYWORD" help:"Permission keyword to create."`
	Description string `placeholder:"DESCRIPTION" help:"Optional description of the permission."`
}

// Run command.
func (cmd *CreatePermCmd) Run(args []string) error {
	if cmd.Help || cmd.Name == "" {
		return opt.ErrUsage
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	err = c.CreatePermission(cmd.Name, cmd.Description)
	return err
}
