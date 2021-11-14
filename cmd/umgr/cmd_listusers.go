// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/grimdork/sprawl"
)

// ListUsersCmd options.
type ListUsersCmd struct{}

// Run command.
func (cmd *ListUsersCmd) Run(args []string) error {
	cfg, err := sprawl.LoadConfig(configPath)
	if err != nil {
		return err
	}

	data, err := cfg.Get(sprawl.EPUsers, nil)
	if err != nil {
		return err
	}

	var list []sprawl.User
	err = json.Unmarshal(data, &list)
	if err != nil {
		return err
	}

	if len(list) == 0 {
		fmt.Println("No users found.")
		return nil
	}

	w := &tabwriter.Writer{}
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	w.Write([]byte("ID\tUsername\tFullname\tE-mail\n"))
	for _, user := range list {
		s := fmt.Sprintf("%d\t%s\t%s\t%s\n", user.ID, user.Name, user.Fullname, user.Email)
		w.Write([]byte(s))
	}
	w.Flush()
	return nil
}
