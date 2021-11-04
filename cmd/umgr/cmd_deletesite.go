package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
)

type DeleteSiteCmd struct {
	opt.DefaultHelp
	Name string `placeholder:"NAME" help:"Name of site to delete."`
}

// Run the site deletion.
func (cmd *DeleteSiteCmd) Run(args []string) error {
	if cmd.Help || cmd.Name == "" {
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

	_, err = cfg.Post(sprawl.EPDeleteSite, sprawl.Request{"name": cmd.Name})
	return err
}
