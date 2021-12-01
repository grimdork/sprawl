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

// ListGroupsCmd options.
type ListGroupsCmd struct {
	opt.DefaultHelp
	Site string `placeholder:"SITE" help:"Site to list groups for."`
}

// Run command.
func (cmd *ListGroupsCmd) Run(args []string) error {
	if cmd.Help {
		return opt.ErrUsage
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	list, err := c.ListGroups(cmd.Site)
	if err != nil {
		return err
	}

	if len(list.Groups) == 0 {
		fmt.Println("No groups found.")
		return nil
	}

	w := &tabwriter.Writer{}
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	w.Write([]byte("ID\tGroup\tSite\n"))
	for _, g := range list.Groups {
		s := fmt.Sprintf("%d\t%s\t%s\n", g.ID, g.Name, g.Site)
		w.Write([]byte(s))
	}
	w.Flush()
	return nil
}
