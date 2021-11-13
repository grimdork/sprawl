//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"os"

	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
)

// AddSiteMemberCmd options.
type AddSiteMemberCmd struct {
	opt.DefaultHelp
	Site  string `placeholder:"SITE" help:"Site to add the member to."`
	Name  string `placeholder:"USERNAME" help:"User to add to the site."`
	Admin bool   `short:"a" help:"Add user as an administrator."`
	Data  string `short:"d" help:"Data to add to the user."`
}

// Run command.
func (cmd AddSiteMemberCmd) Run(args []string) error {
	if cmd.Name == "" {
		return opt.ErrUsage
	}

	var err error
	var data []byte
	if cmd.Data != "" {
		data, err = os.ReadFile(cmd.Data)
		if err != nil {
			return err
		}
	}

	cfg, err := sprawl.LoadConfig(configPath)
	if err != nil {
		return err
	}

	_, err = cfg.Post(sprawl.EPSite+sprawl.EPMember, sprawl.Request{
		"name":  cmd.Name,
		"site":  cmd.Site,
		"admin": fmt.Sprintf("%t", cmd.Admin),
		"data":  string(data),
	})
	return err
}
