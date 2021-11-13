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

	cfg, err := sprawl.LoadConfig(configPath)
	if err != nil {
		return err
	}

	data, err := cfg.Get(sprawl.EPGroups, sprawl.Request{"site": cmd.Site})
	if err != nil {
		return err
	}

	var list sprawl.GroupList
	err = json.Unmarshal(data, &list)
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
