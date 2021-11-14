// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"os"

	"github.com/Urethramancer/ansi"
	"github.com/Urethramancer/signor/opt"
	"github.com/grimdork/sprawl"
	"github.com/grimdork/sprawl/client"
	"golang.org/x/term"
)

// CreateUserCmd options.
type CreateUserCmd struct {
	opt.DefaultHelp
	Name     string `placeholder:"NAME" help:"An alphanumeric username to create."`
	Generate bool   `short:"g" long:"generate" help:"Generate a password and show it instead of prompting for it."`
}

// Run command.
func (cmd *CreateUserCmd) Run(args []string) error {
	if cmd.Help || cmd.Name == "" {
		return opt.ErrUsage
	}

	var pw string
	if !cmd.Generate {
		fmt.Printf("Password: ")
		pass, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return err
		}

		pw = string(pass)
	} else {
		pw = sprawl.RandString(20)
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	err = c.CreateUser(cmd.Name, pw)
	if err != nil {
		return err
	}

	if !cmd.Generate {
		pr("User added.")
	} else {
		pr("User added with password %s%s%s", ansi.Green, pw, ansi.Normal)
	}

	return nil
}
