// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
)

type DeleteUserCmd struct {
	opt.DefaultHelp
	Username string `placeholder:"NAME" help:"Nameof the user to delete."`
}

// Run the user deletion.
func (cmd *DeleteUserCmd) Run(args []string) error {
	if cmd.Help || cmd.Username == "" {
		return opt.ErrUsage
	}

	cfg, err := sprawl.LoadConfig(configPath)
	if err != nil {
		return err
	}

	err = cfg.GetLoginToken()
	if err != nil {
		return err
	}

	_, err = cfg.Post(sprawl.EPDeleteUser, sprawl.Request{"name": cmd.Username})
	if err != nil {
		return err
	}

	pr("User deleted.")
	return nil
}
