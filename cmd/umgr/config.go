// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/grimdork/sprawl"
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
func LoadConfig() (*Config, error) {
	var cfg Config
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	err = cfg.GetLoginToken()
	return &cfg, err
}

// Save Config to JSON file.
func (cfg *Config) Save() error {
	data, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0600)
}

// Get an endpoint and return the JSON string or an error.
func (cfg *Config) Get(ep string, args sprawl.Request) ([]byte, error) {
	url := fmt.Sprintf("%s%s", cfg.URL, ep)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("username", cfg.Username)
	req.Header.Set("token", cfg.Token)
	for k, v := range args {
		req.Header.Set(k, v)
	}

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
func (cfg *Config) Post(ep string, args sprawl.Request) ([]byte, error) {
	url := fmt.Sprintf("%s%s", cfg.URL, ep)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("username", cfg.Username)
	req.Header.Set("token", cfg.Token)
	for k, v := range args {
		req.Header.Set(k, v)
	}

	c := &http.Client{Timeout: time.Second * 10}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("couldn't POST: %s", res.Status)
	}

	defer res.Body.Close()
	return io.ReadAll(res.Body)
}
