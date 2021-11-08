// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"os"

	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
	"golang.org/x/term"
)

// InitCmd flags.
type InitCmd struct {
	opt.DefaultHelp
	Force bool `short:"f" long:"force" help:"Force (re-)initialization."`
}

// Run connection setup.
func (cmd *InitCmd) Run(args []string) error {
	if cmd.Help {
		return opt.ErrUsage
	}

	cfg, err := sprawl.LoadConfig(configPath)
	if err == nil && !cmd.Force {
		pr("Already initialized. Use -f to force re-initialization.")
		return nil
	}

	fmt.Printf("URL (<protocol>://<domain.tld>): ")
	fmt.Scanln(&cfg.URL)
	fmt.Printf("Username: ")
	fmt.Scanln(&cfg.Username)
	fmt.Printf("Password: ")
	pw, err := term.ReadPassword(int(os.Stdin.Fd()))
	println()
	if err != nil {
		return err
	}

	cfg.Password = string(pw)
	return cfg.Save(configPath)
}
