package main

import (
	"github.com/grimdork/opt"
	"github.com/grimdork/sprawl/client"
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

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	return c.RemoveGroupMember(cmd.Site, cmd.Group, cmd.Name)
}
