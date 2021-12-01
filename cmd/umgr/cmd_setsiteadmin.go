//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"github.com/grimdork/opt"
	"github.com/grimdork/sprawl/client"
)

// SetSiteAdminCmd options.
type SetSiteAdminCmd struct {
	opt.DefaultHelp
	Site    string `placeholder:"SITE" help:"Site the user is a member of"`
	Name    string `placeholder:"USERNAME" help:"Name of the user."`
	Disable bool   `short:"d" long:"disable" help:"Disable admin status instead of enabling it."`
}

// Run command.
func (cmd SetSiteAdminCmd) Run(args []string) error {
	if cmd.Name == "" {
		return opt.ErrUsage
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	return c.SetSiteAdmin(cmd.Site, cmd.Name, cmd.Disable)
}
