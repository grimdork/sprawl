package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
	"github.com/grimdork/sprawl/client"
)

// DeleteSiteCmd options.
type DeleteSiteCmd struct {
	opt.DefaultHelp
	Name string `placeholder:"SITE" help:"Name of site to delete."`
}

// Run command.
func (cmd *DeleteSiteCmd) Run(args []string) error {
	if cmd.Help || cmd.Name == "" {
		return opt.ErrUsage
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	err = c.Delete(sprawl.EPSite, sprawl.Request{"name": cmd.Name})
	return err
}
