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

// ListSiteMembersCmd options.
type ListSiteMembersCmd struct {
	opt.DefaultHelp
	Site string `placeholder:"Site" help:"Site to list."`
}

// Run command.
func (cmd *ListSiteMembersCmd) Run(args []string) error {
	if cmd.Site == "" {
		return opt.ErrUsage
	}

	cfg, err := sprawl.LoadConfig(configPath)
	if err != nil {
		return err
	}

	data, err := cfg.Get(sprawl.EPSite+sprawl.EPMembers, sprawl.Request{"site": cmd.Site})
	if err != nil {
		return err
	}

	var list []sprawl.User
	err = json.Unmarshal(data, &list)
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
