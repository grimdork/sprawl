// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"github.com/grimdork/opt"
	"github.com/grimdork/sprawl/client"
)

// DeleteUserCmd options.
type DeleteUserCmd struct {
	opt.DefaultHelp
	Username string `placeholder:"NAME" help:"Nameof the user to delete."`
}

// Run command.
func (cmd *DeleteUserCmd) Run(args []string) error {
	if cmd.Help || cmd.Username == "" {
		return opt.ErrUsage
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	err = c.DeleteUser(cmd.Username)
	if err != nil {
		return err
	}

	pr("User %s deleted.", cmd.Username)
	return nil
}
