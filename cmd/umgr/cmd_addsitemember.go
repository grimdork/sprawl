//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"os"

	"github.com/grimdork/opt"
	"github.com/grimdork/sprawl/client"
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

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	err = c.AddSiteMember(cmd.Site, cmd.Name, string(data), cmd.Admin)
	return err
}
