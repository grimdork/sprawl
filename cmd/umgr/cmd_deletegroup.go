//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
)

// DeleteGroupCmd options.
type DeleteGroupCmd struct {
	opt.DefaultHelp
	Name string `placeholder:"NAME" help:"Name of group to delete."`
	Site string `placeholder:"SITE" help:"Site of the group."`
}

// Run command.
func (cmd *DeleteGroupCmd) Run(args []string) error {
	if cmd.Help || cmd.Name == "" {
		return opt.ErrUsage
	}

	cfg, err := sprawl.LoadConfig(configPath)
	if err != nil {
		return err
	}

	err = cfg.Delete(sprawl.EPGroup, sprawl.Request{
		"name": cmd.Name,
		"site": cmd.Site,
	})
	return err
}
