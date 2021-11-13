package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
)

// DeleteSiteCmd options.
type DeleteSiteCmd struct {
	opt.DefaultHelp
	Name string `placeholder:"NAME" help:"Name of site to delete."`
}

// Run command.
func (cmd *DeleteSiteCmd) Run(args []string) error {
	if cmd.Help || cmd.Name == "" {
		return opt.ErrUsage
	}

	cfg, err := sprawl.LoadConfig(configPath)
	if err != nil {
		return err
	}

	err = cfg.GetLoginToken()
	if err != nil {
		return err
	}

	err = cfg.Delete(sprawl.EPSite, sprawl.Request{"name": cmd.Name})
	return err
}
