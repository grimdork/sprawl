package main

import (
	"github.com/grimdork/opt"
	"github.com/grimdork/sprawl/client"
)

// CreateSiteCmd options.
type CreateSiteCmd struct {
	opt.DefaultHelp
	Name string `placeholder:"NAME" help:"A site name to create (use domain name rules for allowed symbols)."`
}

// Run command.
func (cmd *CreateSiteCmd) Run(args []string) error {
	if cmd.Help || cmd.Name == "" {
		return opt.ErrUsage
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	err = c.CreateSite(cmd.Name)
	return err
}
