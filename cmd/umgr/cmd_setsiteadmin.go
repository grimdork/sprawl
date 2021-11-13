//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
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

	cfg, err := sprawl.LoadConfig(configPath)
	if err != nil {
		return err
	}

	req := sprawl.Request{
		"site": cmd.Site,
		"name": cmd.Name,
	}

	if cmd.Disable {
		err = cfg.Delete(sprawl.EPSite+sprawl.EPAdmin, req)
	} else {
		err = cfg.Put(sprawl.EPSite+sprawl.EPAdmin, req)
	}
	return err
}
