package main

import "github.com/Urethramancer/signor/opt"

// ImportCmd options.
type ImportCmd struct {
	opt.DefaultHelp
	File string `placeholder:"FILE" help:"JSON file with user database to create."`
}

// Run command.
func (cmd *ImportCmd) Run(args []string) error {
	if cmd.Help || cmd.File == "" {
		return opt.ErrUsage
	}

	return nil
}
