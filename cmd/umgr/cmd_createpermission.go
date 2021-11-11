// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
)

type CreatePermCmd struct {
	opt.DefaultHelp
	Name        string `placeholder:"KEYWORD" help:"Permission keyword to create."`
	Description string `placeholder:"DESCRIPTION" help:"Optional description of the permission."`
}

// Run the permission creation.
func (cmd *CreatePermCmd) Run(args []string) error {
	if cmd.Help || cmd.Name == "" {
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

	_, err = cfg.Post(sprawl.EPPermission, sprawl.Request{
		"name":        cmd.Name,
		"description": cmd.Description,
	})
	return err
}
