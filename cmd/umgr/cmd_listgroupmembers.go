//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/grimdork/opt"
	"github.com/grimdork/sprawl/client"
)

// ListGroupMembersCmd options.
type ListGroupMembersCmd struct {
	opt.DefaultHelp
	Site  string `placeholder:"SITE" help:"Site of the group."`
	Group string `placeholder:"GROUP" help:"Group to list."`
}

// Run command.
func (cmd *ListGroupMembersCmd) Run(args []string) error {
	if cmd.Site == "" {
		return opt.ErrUsage
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	list, err := c.ListGroupMembers(cmd.Site, cmd.Group)
	if err != nil {
		return err
	}

	if len(list.Users) == 0 {
		fmt.Printf("No members found.\n")
		return nil
	}

	w := &tabwriter.Writer{}
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	w.Write([]byte("ID\tUsername\n"))
	for _, user := range list.Users {
		s := fmt.Sprintf("%d\t%s\n", user.ID, user.Name)
		w.Write([]byte(s))
	}
	w.Flush()
	return nil
}
