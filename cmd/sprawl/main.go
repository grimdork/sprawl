// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"os"

	"github.com/Urethramancer/daemon"
)

func main() {
	srv, err := NewServer()
	if err != nil {
		fmt.Printf("Couldn't start server: %s\n", err.Error())
		os.Exit(2)
	}

	srv.Start()
	<-daemon.BreakChannel()
	srv.Stop()
}
