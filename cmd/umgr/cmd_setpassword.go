// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"os"

	"github.com/grimdork/opt"
	"github.com/grimdork/sprawl"
	"github.com/grimdork/sprawl/client"
	"golang.org/x/term"
)

// SetPasswordCmd options.
type SetPasswordCmd struct {
	opt.DefaultHelp
	Name    string `placeholder:"NAME" help:"Username to set the password for."`
	AskPass bool   `short:"p" long:"password" help:"Prompt for a password instead of generating it."`
}

// Run command.
func (cmd *SetPasswordCmd) Run(args []string) error {
	if cmd.Help || cmd.Name == "" {
		return opt.ErrUsage
	}

	var pw string
	if cmd.AskPass {
		fmt.Printf("Password: ")
		pass, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return err
		}

		println("")
		pw = string(pass)
	} else {
		pw = sprawl.RandString(20)
	}

	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	err = c.SetPassword(cmd.Name, pw)
	if err != nil {
		return err
	}

	if !cmd.AskPass {
		println(pw)
	}
	return nil
}
