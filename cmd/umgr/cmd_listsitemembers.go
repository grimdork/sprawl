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

// ListSiteMembersCmd options.
type ListSiteMembersCmd struct {
	opt.DefaultHelp
	Site string `placeholder:"SITE" help:"Site to list."`
}

// Run command.
func (cmd *ListSiteMembersCmd) Run(args []string) error {
	if cmd.Site == "" {
		return opt.ErrUsage
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	list, err := c.ListSiteMembers(cmd.Site)
	if err != nil {
		return err
	}

	if len(list) == 0 {
		fmt.Printf("No members found.\n")
		return nil
	}

	w := &tabwriter.Writer{}
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	w.Write([]byte("ID\tUsername\tProfiles\tAdmin\n"))
	for _, user := range list {
		s := fmt.Sprintf("%d\t%s\t%t\t%t\n", user.ID, user.Name, user.Profile != "", user.Admin)
		w.Write([]byte(s))
	}
	w.Flush()
	return nil
}
