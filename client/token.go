// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/grimdork/sprawl"
)

// GetLoginToken gets a speawl login token or the current stored token if valid.
func (cfg *Config) GetLoginToken() error {
	url := fmt.Sprintf("%s%s", cfg.URL, sprawl.EPAuth)
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("username", cfg.Username)
	req.Header.Set("password", cfg.Password)
	c := &http.Client{Timeout: time.Second * 10}
	res, err := c.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, cfg)
}
