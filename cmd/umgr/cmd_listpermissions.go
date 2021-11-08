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

	"github.com/grimdork/sprawl"
)

type ListPermsCmd struct{}

// Run group listing.
func (cmd *ListPermsCmd) Run(args []string) error {
	cfg, err := sprawl.LoadConfig(configPath)
	if err != nil {
		return err
	}

	data, err := cfg.Get(sprawl.EPListPermissions, nil)
	if err != nil {
		return err
	}

	var list sprawl.PermissionList
	err = json.Unmarshal(data, &list)
	if err != nil {
		return err
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
