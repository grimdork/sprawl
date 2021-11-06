//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
)

type DeleteGroupCmd struct {
	opt.DefaultHelp
	Name string `placeholder:"NAME" help:"Name of group to delete."`
	Site string `placeholder:"site" help:"Site of the group."`
}

// Run the group deletion.
func (cmd *DeleteGroupCmd) Run(args []string) error {
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

	_, err = cfg.Post(sprawl.EPDeleteGroup, sprawl.Request{
		"name": cmd.Name,
		"site": cmd.Site,
	})
	return err
}
