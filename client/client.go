package client

import (
	"encoding/json"

	"github.com/grimdork/sprawl"
)

// SprawlClient is used to access the Sprawl API.
type SprawlClient struct {
	*Config
}

// New creates a new SprawlClient.
func New(fn string) (*SprawlClient, error) {
	c := &SprawlClient{}
	cfg, err := loadConfig(fn)
	if err != nil {
		return nil, err
	}

	c.Config = cfg
	return c, nil
}

//
// User API
//

// CreateUser creates a new user with a username and password.
// Use UpdateUser() to set additional fields.
func (c *SprawlClient) CreateUser(name, password string) error {
	_, err := c.Post(sprawl.EPUser, sprawl.Request{
		"name":     name,
		"password": password,
	})
	return err
}

func (c *SprawlClient) UpdateUser(name, password string) error {
	return nil
}

func (c *SprawlClient) DeleteUser(name string) error {
	err := c.Delete(sprawl.EPUser, sprawl.Request{"name": name})
	return err
}

func (c *SprawlClient) GetUser(name string) (sprawl.User, error) {
	var u sprawl.User
	data, err := c.Get(sprawl.EPUsers, nil)
	if err != nil {
		return u, err
	}

	err = json.Unmarshal(data, &u)
	return u, err
}

func (c *SprawlClient) ListUsers() ([]sprawl.User, error) {
	data, err := c.Get(sprawl.EPUsers, nil)
	if err != nil {
		return nil, err
	}

	var list []sprawl.User
	err = json.Unmarshal(data, &list)
	if err != nil {
		return nil, err
	}

	return list, err
}

//
// Permission API
//

func (c *SprawlClient) CreatePermission(name, description string) error {
	_, err := c.Post(sprawl.EPPermission, sprawl.Request{
		"name":        name,
		"description": description,
	})
	return err
}

func (c *SprawlClient) UpdatePermission(name, description string) error {
	err := c.Put(sprawl.EPPermission, sprawl.Request{
		"name":        name,
		"description": description,
	})
	return err
}

func (c *SprawlClient) DeletePermission(name string) error {
	err := c.Delete(sprawl.EPPermission, sprawl.Request{"name": name})
	return err
}

func (c *SprawlClient) ListPermissions() (sprawl.PermissionList, error) {
	data, err := c.Get(sprawl.EPPermissions, nil)
	if err != nil {
		return sprawl.PermissionList{}, err
	}

	var list sprawl.PermissionList
	err = json.Unmarshal(data, &list)
	return list, err
}
