//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
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

	data, err := c.Get(sprawl.EPGroup+sprawl.EPMembers, sprawl.Request{
		"site":  cmd.Site,
		"group": cmd.Group,
	})
	if err != nil {
		return err
	}

	var list sprawl.UserList
	err = json.Unmarshal(data, &list)
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
