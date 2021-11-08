// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"os"
	"path/filepath"

	"github.com/grimdork/xos"
)

const program = "sprawlmgr"

var configPath string

func init() {
	cfg, err := xos.NewConfig(program)
	if err != nil {
		pr("Error: %s", err.Error())
		os.Exit(2)
	}

	err = os.MkdirAll(cfg.Path(), 0700)
	if err != nil {
		pr("Error: %s", err.Error())
		os.Exit(2)
	}

	configPath = filepath.Join(cfg.Path(), "config.json")
}
