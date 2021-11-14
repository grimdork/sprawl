//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl/client"
)

// DeletePermCmd options.
type DeletePermCmd struct {
	opt.DefaultHelp
	Name string `placeholder:"NAME" help:"Name of permission to delete."`
}

// Run command.
func (cmd *DeletePermCmd) Run(args []string) error {
	if cmd.Help || cmd.Name == "" {
		return opt.ErrUsage
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	err = c.DeletePermission(cmd.Name)
	return err
}
