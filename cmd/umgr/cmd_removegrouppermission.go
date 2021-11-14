package main

import (
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
	"github.com/grimdork/sprawl/client"
)

// RemoveGroupPermissionCmd foptions.
type RemoveGroupPermissionCmd struct {
	opt.DefaultHelp
	Site    string `placeholder:"SITE" help:"Site of the group."`
	Group   string `placeholder:"GROUP" help:"Group to remove from."`
	Keyword string `placeholder:"KEYWORD" help:"Permission to remove from the group."`
}

// Run the command.
func (cmd *RemoveGroupPermissionCmd) Run(args []string) error {
	if cmd.Help || cmd.Keyword == "" {
		return opt.ErrUsage
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	err = c.Delete(sprawl.EPGroup+sprawl.EPPermission, sprawl.Request{
		"site":       cmd.Site,
		"group":      cmd.Group,
		"permission": cmd.Keyword,
	})
	return err
}
