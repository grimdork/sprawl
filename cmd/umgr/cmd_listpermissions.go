//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

package main

import (
	"encoding/json"

	"github.com/grimdork/sprawl"
)

type ListPermissionsCmd struct{}

// Run group listing.
func (cmd *ListPermissionsCmd) Run(args []string) error {
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	data, err := cfg.Get(sprawl.EPListPermissions, nil)
	if err != nil {
		return err
	}

	var list []string
	err = json.Unmarshal(data, &list)
	if err != nil {
		return err
	}

	for _, g := range list {
		println(g)
	}
	return nil
}
