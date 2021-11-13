package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
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

	cfg, err := sprawl.LoadConfig(configPath)
	if err != nil {
		return err
	}

	_, err = cfg.Post(sprawl.EPSite, sprawl.Request{"name": cmd.Name})
	return err
}
