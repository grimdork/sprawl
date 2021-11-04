//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
)

type CreateGroupCmd struct {
	opt.DefaultHelp
	Name string `placeholder:"NAME" help:"An alphanumeric group name to create."`
	Site string `placeholder:"SITE" help:"The site to create the groups for."`
}

// Run the group creation.
func (cmd *CreateGroupCmd) Run(args []string) error {
	if cmd.Help || cmd.Site == "" {
		return opt.ErrUsage
	}

	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	_, err = cfg.Post(sprawl.EPCreateGroup, sprawl.Request{
		"name": cmd.Name,
		"site": cmd.Site},
	)
	return err
}
