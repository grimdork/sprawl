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

type ListGroupsCmd struct {
	opt.DefaultHelp
	Site string `placeholder:"SITE" help:"Site to list groups for."`
}

// Run group listing.
func (cmd *ListGroupsCmd) Run(args []string) error {
	if cmd.Help {
		return opt.ErrUsage
	}

	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	data, err := cfg.Get(sprawl.EPListGroups, map[string]string{"site": cmd.Site})
	if err != nil {
		return err
	}

	var list sprawl.GroupList
	err = json.Unmarshal(data, &list)
	if err != nil {
		return err
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
