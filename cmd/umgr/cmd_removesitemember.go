package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
)

// RemoveSiteMemberCmd foptions.
type RemoveSiteMemberCmd struct {
	opt.DefaultHelp
	Site string `placeholder:"SITE" help:"Site to remove from."`
	Name string `placeholder:"USER" help:"User to remove."`
}

// Run the command.
func (cmd *RemoveSiteMemberCmd) Run(args []string) error {
	if cmd.Help || cmd.Name == "" {
		return opt.ErrUsage
	}

	cfg, err := sprawl.LoadConfig(configPath)
	if err != nil {
		return err
	}

	err = cfg.Delete(sprawl.EPSite+sprawl.EPMember, sprawl.Request{
		"site": cmd.Site,
		"name": cmd.Name,
	})
	return err
}
