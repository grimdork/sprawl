package main

import (
	"github.com/grimdork/opt"
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

	return c.DeleteSite(cmd.Name)
}
