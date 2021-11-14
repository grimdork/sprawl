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

// ListGroupPermissionsCmd options.
type ListGroupPermissionsCmd struct {
	opt.DefaultHelp
	Site  string `placeholder:"SITE" help:"Site of the group."`
	Group string `placeholder:"GROUP" help:"Group to list."`
}

// Run command.
func (cmd *ListGroupPermissionsCmd) Run(args []string) error {
	if cmd.Site == "" {
		return opt.ErrUsage
	}

	cfg, err := sprawl.LoadConfig(configPath)
	if err != nil {
		return err
	}

	data, err := cfg.Get(sprawl.EPGroup+sprawl.EPPermissions, sprawl.Request{
		"site":  cmd.Site,
		"group": cmd.Group,
	})
	if err != nil {
		return err
	}

	var list sprawl.PermissionList
	err = json.Unmarshal(data, &list)
	if err != nil {
		return err
	}

	if len(list.Permissions) == 0 {
		fmt.Printf("No permissions found.\n")
		return nil
	}

	w := &tabwriter.Writer{}
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	w.Write([]byte("ID\tPermission\tDescription\n"))
	for _, p := range list.Permissions {
		s := fmt.Sprintf("%d\t%s\t%s\n", p.ID, p.Name, p.Description)
		w.Write([]byte(s))
	}
	w.Flush()
	return nil
}
