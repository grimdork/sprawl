package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
)

type CreateSiteCmd struct {
	opt.DefaultHelp
	Name string `placeholder:"NAME" help:"A site name to create (use domain name rules for allowed symbols)."`
}

// Run the site creation.
func (cmd *CreateSiteCmd) Run(args []string) error {
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

	_, err = cfg.Post(sprawl.EPCreateSite, sprawl.Request{"name": cmd.Name})
	return err
}
