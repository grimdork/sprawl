package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/grimdork/sprawl"
)

// Config for a Sprawl connection.
type Config struct {
	// URL for the sprawl server.
	URL string `json:"url"`
	// Username for an administrator.
	Username string `json:"username"`
	// Password for the same administrator.
	Password string `json:"password"`
	// Token for session.
	Token string
}

// loadConfig from JSON file.
func loadConfig(fn string) (*Config, error) {
	var cfg Config
	data, err := os.ReadFile(fn)
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
func (cfg *Config) Save(fn string) error {
	data, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(fn, data, 0600)
}

func (cfg *Config) request(method, ep string, args sprawl.Request) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", cfg.URL, ep)
	req, err := http.NewRequest(method, url, nil)
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
	return res, err
}

// Get is for retrieval.
func (cfg *Config) Get(ep string, args sprawl.Request) ([]byte, error) {
	res, err := cfg.request(http.MethodGet, ep, args)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("couldn't GET: %s", res.Status)
	}

	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

// Post is for creation.
func (cfg *Config) Post(ep string, args sprawl.Request) ([]byte, error) {
	res, err := cfg.request(http.MethodPost, ep, args)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("couldn't POST: %s", res.Status)
	}

	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

// Delete is for removal.
func (cfg *Config) Delete(ep string, args sprawl.Request) error {
	res, err := cfg.request(http.MethodDelete, ep, args)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return fmt.Errorf("couldn't DELETE: %s", res.Status)
	}

	defer res.Body.Close()
	return nil
}

// Put is for updates.
func (cfg *Config) Put(ep string, args sprawl.Request) error {
	res, err := cfg.request(http.MethodPut, ep, args)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return fmt.Errorf("couldn't PUT: %s", res.Status)
	}

	defer res.Body.Close()
	return nil
}
