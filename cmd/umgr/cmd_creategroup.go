//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import "github.com/Urethramancer/signor/opt"

type CreateGroupCmd struct {
	opt.DefaultHelp
	Site string   `placeholder:"SITE" help:"The site to create the groups for."`
	Name []string `placeholder:"NAME" help:"An alphanumeric group name to create."`
}

// Run the group creation.
func (cmd *CreateGroupCmd) Run(args []string) error {
	if cmd.Help || cmd.Name == nil {
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

	return nil
}
