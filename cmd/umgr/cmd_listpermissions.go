//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/grimdork/sprawl/client"
)

// ListPermsCmd options.
type ListPermsCmd struct{}

// Run command.
func (cmd *ListPermsCmd) Run(args []string) error {
	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	list, err := c.ListPermissions()
	if err != nil {
		return err
	}

	if len(list.Permissions) == 0 {
		fmt.Println("No permissions found.")
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
