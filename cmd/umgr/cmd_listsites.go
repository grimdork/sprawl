package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/grimdork/sprawl"
)

type ListSitesCmd struct{}

// Run the site listing.
func (cmd *ListSitesCmd) Run(args []string) error {
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}

	err = cfg.GetLoginToken()
	if err != nil {
		return err
	}

	data, err := cfg.Get(sprawl.EPListSites, nil)
	if err != nil {
		return err
	}

	var list []sprawl.Site
	err = json.Unmarshal(data, &list)
	if err != nil {
		return err
	}

	w := &tabwriter.Writer{}
	w.Init(os.Stdout, 0, 8, 1, '\t', 0)
	w.Write([]byte("UID\tDomain\n"))
	for _, site := range list {
		s := fmt.Sprintf("%d\t%s\n", site.ID, site.Name)
		w.Write([]byte(s))
	}
	w.Flush()
	return nil
}
