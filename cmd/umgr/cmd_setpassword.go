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

type SetPasswordCmd struct {
	opt.DefaultHelp
	Name    string `placeholder:"NAME" help:"Username to set the password for."`
	AskPass bool   `short:"p" long:"password" help:"Prompt for a password instead of generating it."`
}

// Run the user creation.
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

	cfg, err := sprawl.LoadConfig(configPath)
	if err != nil {
		return err
	}

	err = cfg.GetLoginToken()
	if err != nil {
		return err
	}

	_, err = cfg.Post(sprawl.EPUser+sprawl.EPSetPassword, sprawl.Request{
		"name":     cmd.Name,
		"password": pw,
	})
	if err != nil {
		return err
	}

	if !cmd.AskPass {
		println(pw)
	}
	return nil
}
