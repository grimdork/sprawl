// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"os"

	"github.com/grimdork/opt"
	"github.com/grimdork/sprawl/client"
	"golang.org/x/term"
)

// InitCmd options.
type InitCmd struct {
	opt.DefaultHelp
	Force bool `short:"f" long:"force" help:"Force (re-)initialization."`
}

// Run command.
func (cmd *InitCmd) Run(args []string) error {
	if cmd.Help {
		return opt.ErrUsage
	}

	_, err := client.New(configPath)
	if err == nil && !cmd.Force {
		pr("Already initialized. Use -f to force re-initialization.")
		return nil
	}

	cfg := &client.Config{}
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
