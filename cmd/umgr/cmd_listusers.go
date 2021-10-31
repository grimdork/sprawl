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

// ListUsersCmd flags.
type ListUsersCmd struct{}

// Run listusers.
func (cmd *ListUsersCmd) Run(args []string) error {
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	err = cfg.GetLoginToken()
	if err != nil {
		return err
	}

	data, err := cfg.Get(sprawl.EPListUsers)
	if err != nil {
		return err
	}

	var list []sprawl.User
	err = json.Unmarshal(data, &list)
	if err != nil {
		return err
	}

	w := &tabwriter.Writer{}
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	w.Write([]byte("UID\tUsername\n"))
	for _, user := range list {
		s := fmt.Sprintf("%d\t%s\n", user.ID, user.Name)
		w.Write([]byte(s))
	}
	w.Flush()
	return nil
}
