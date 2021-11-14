package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/grimdork/sprawl"
	"github.com/grimdork/sprawl/client"
)

// ListSitesCmd options.
type ListSitesCmd struct{}

// Run command.
func (cmd *ListSitesCmd) Run(args []string) error {
	c, err := client.New(configPath)
	if err != nil {
		return err
	}

	data, err := c.Get(sprawl.EPSites, nil)
	if err != nil {
		return err
	}

	var list []sprawl.Site
	err = json.Unmarshal(data, &list)
	if err != nil {
		return err
	}

	if len(list) == 0 {
		fmt.Println("No sites found.")
		return nil
	}

	w := &tabwriter.Writer{}
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	w.Write([]byte("ID\tDomain\n"))
	for _, site := range list {
		s := fmt.Sprintf("%d\t%s\n", site.ID, site.Name)
		w.Write([]byte(s))
	}
	w.Flush()
	return nil
}
