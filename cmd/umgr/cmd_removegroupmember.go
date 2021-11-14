package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
)

// RemoveGroupMemberCmd foptions.
type RemoveGroupMemberCmd struct {
	opt.DefaultHelp
	Site  string `placeholder:"SITE" help:"Site of the group."`
	Group string `placeholder:"GROUP" help:"Group to remove the member from."`
	Name  string `placeholder:"USERNAME" help:"User to remove."`
}

// Run the command.
func (cmd *RemoveGroupMemberCmd) Run(args []string) error {
	if cmd.Help || cmd.Name == "" {
		return opt.ErrUsage
	}

	cfg, err := sprawl.LoadConfig(configPath)
	if err != nil {
		return err
	}

	err = cfg.Delete(sprawl.EPGroup+sprawl.EPMember, sprawl.Request{
		"site":  cmd.Site,
		"group": cmd.Group,
		"name":  cmd.Name,
	})
	return err
}
