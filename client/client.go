package client

import (
	"encoding/json"
	"fmt"

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

// NewWithSettings creates a new SprawlClient from variables in the environment.
func NewWithSettings(host, username, password string) (*SprawlClient, error) {
	cfg := &Config{
		URL:      host,
		Username: username,
		Password: password,
	}

	err := cfg.GetLoginToken()
	return &SprawlClient{cfg}, err
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
// Site API
//

func (c *SprawlClient) CreateSite(name string) error {
	_, err := c.Post(sprawl.EPSite, sprawl.Request{
		"name": name,
	})
	return err
}

func (c *SprawlClient) DeleteSite(name string) error {
	err := c.Delete(sprawl.EPSite, sprawl.Request{"name": name})
	return err
}

func (c *SprawlClient) ListSites() ([]sprawl.Site, error) {
	var list []sprawl.Site
	data, err := c.Get(sprawl.EPSites, nil)
	if err != nil {
		return list, err
	}

	err = json.Unmarshal(data, &list)
	return list, err
}

//
// Site member API
//

func (c *SprawlClient) AddSiteMember(site, name, data string, admin bool) error {
	_, err := c.Post(sprawl.EPSite+sprawl.EPMember, sprawl.Request{
		"site":  site,
		"name":  name,
		"admin": fmt.Sprintf("%t", admin),
		"data":  string(data),
	})
	return err
}

func (c *SprawlClient) RemoveSiteMember(site, name string) error {
	err := c.Delete(sprawl.EPSite+sprawl.EPMember, sprawl.Request{
		"site": site,
		"name": name,
	})
	return err
}

func (c *SprawlClient) ListSiteMembers(site string) ([]sprawl.User, error) {
	var list []sprawl.User
	data, err := c.Get(sprawl.EPSite+sprawl.EPMembers, sprawl.Request{"site": site})
	if err != nil {
		return list, err
	}

	err = json.Unmarshal(data, &list)
	return list, err
}

//
// Group API
//

//
// Permission API
//

// CreatePermission creates a new permission with a name and description.
func (c *SprawlClient) CreatePermission(name, description string) error {
	_, err := c.Post(sprawl.EPPermission, sprawl.Request{
		"name":        name,
		"description": description,
	})
	return err
}

// GetPermission returns a permission and description.
func (c *SprawlClient) GetPermission(name string) (*sprawl.Permission, error) {
	var p sprawl.Permission
	data, err := c.Get(sprawl.EPPermission, sprawl.Request{"name": name})
	if err != nil {
		return &p, err
	}

	err = json.Unmarshal(data, &p)
	return &p, err
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
