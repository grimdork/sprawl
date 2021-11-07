//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
)

type DeletePermCmd struct {
	opt.DefaultHelp
	Name string `placeholder:"NAME" help:"Name of permission to delete."`
}

// Run the permission deletion.
func (cmd *DeletePermCmd) Run(args []string) error {
	if cmd.Help || cmd.Name == "" {
		return opt.ErrUsage
	}

	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	err = cfg.GetLoginToken()
	if err != nil {
		return err
	}

	_, err = cfg.Post(sprawl.EPDeletePermission, sprawl.Request{
		"name": cmd.Name,
	})
	return err
}
