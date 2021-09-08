// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/grimdork/xos"
)

// Config for sprawl connection.
type Config struct {
	// URL for the sprawl server.
	URL string
	// Username for an administrator.
	Username string `json:"username"`
	// Password for the same administrator.
	Password string `json:"password"`
	// Token for session.
	Token string
}

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

// LoadConfig from JSON file.
func LoadConfig() (Config, error) {
	var config Config
	data, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	return config, err
}

// Save Config to JSON file.
func SaveConfig(config Config) error {
	data, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0600)
}

// Get an endpoint and return the JSON string or an error.
func (cfg *Config) Get(ep string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", cfg.URL, ep)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("username", cfg.Username)
	req.Header.Set("token", cfg.Token)
	c := &http.Client{Timeout: time.Second * 10}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 400 {
		return nil, fmt.Errorf("couldn't GET: %s", res.Status)
	}

	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

// Post to an endpoint and return the JSON string or an error.
func (cfg *Config) Post(ep string, args interface{}) ([]byte, error) {
	data, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s%s", cfg.URL, ep)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("username", cfg.Username)
	req.Header.Set("token", cfg.Token)
	c := &http.Client{Timeout: time.Second * 10}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 400 {
		return nil, fmt.Errorf("couldn't POST: %s", res.Status)
	}

	defer res.Body.Close()
	return io.ReadAll(res.Body)
}
