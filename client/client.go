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

// Token from successful authentication.
type Token struct {
	// Data is the token string.
	Data string `json:"token"`
}

// AuthUser authenticates a user.
func (c *SprawlClient) AuthUser(name, password string) (string, error) {
	data, err := c.Post(sprawl.EPAuth, sprawl.Request{
		"username": name,
		"password": password,
	})
	if err != nil {
		return "", err
	}

	var t Token
	err = json.Unmarshal(data, &t)
	if err != nil {
		return "", err
	}

	return t.Data, nil
}

// CreateUser with a username and password.
// Use UpdateUser() to set additional fields.
func (c *SprawlClient) CreateUser(name, password string) error {
	_, err := c.Post(sprawl.EPUser, sprawl.Request{
		"name":     name,
		"password": password,
	})
	return err
}

// UpdateUser details like email and full name.
func (c *SprawlClient) UpdateUser(name, password string) error {
	return nil
}

// DeleteUser permanently.
func (c *SprawlClient) DeleteUser(name string) error {
	err := c.Delete(sprawl.EPUser, sprawl.Request{"name": name})
	return err
}

// GetUser details.
func (c *SprawlClient) GetUser(name string) (sprawl.User, error) {
	var u sprawl.User
	data, err := c.Get(sprawl.EPUsers, nil)
	if err != nil {
		return u, err
	}

	err = json.Unmarshal(data, &u)
	return u, err
}

// ListUsers returns a list of users.
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

// SetPassword for a user.
func (c *SprawlClient) SetPassword(name, password string) error {
	_, err := c.Post(sprawl.EPUser+sprawl.EPSetPassword, sprawl.Request{
		"name":     name,
		"password": password,
	})
	return err
}

// SetSiteAdmin sets admin status of a user on a site.
func (c *SprawlClient) SetSiteAdmin(site, name string, admin bool) error {
	return c.Put(sprawl.EPSite+sprawl.EPAdmin, sprawl.Request{
		"site":  site,
		"name":  name,
		"admin": fmt.Sprintf("%t", admin),
	})
}

//
// Site API
//

// CreateSite creates a new site.
func (c *SprawlClient) CreateSite(name string) error {
	_, err := c.Post(sprawl.EPSite, sprawl.Request{
		"name": name,
	})
	return err
}

// DeleteSite deletes a site.
func (c *SprawlClient) DeleteSite(name string) error {
	return c.Delete(sprawl.EPSite, sprawl.Request{"name": name})
}

// ListSites returns a list of sites.
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

// AddSiteMember to a site.
func (c *SprawlClient) AddSiteMember(site, name, data string, admin bool) error {
	_, err := c.Post(sprawl.EPSite+sprawl.EPMember, sprawl.Request{
		"site":  site,
		"name":  name,
		"admin": fmt.Sprintf("%t", admin),
		"data":  string(data),
	})
	return err
}

// RemoveSiteMember from a site.
func (c *SprawlClient) RemoveSiteMember(site, name string) error {
	err := c.Delete(sprawl.EPSite+sprawl.EPMember, sprawl.Request{
		"site": site,
		"name": name,
	})
	return err
}

// ListSiteMembers of a site.
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

// CreateGroup on a site.
func (c *SprawlClient) CreateGroup(site, name string) error {
	_, err := c.Post(sprawl.EPGroup, sprawl.Request{
		"site": site,
		"name": name,
	})
	return err
}

// DeleteGroup from a site.
func (c *SprawlClient) DeleteGroup(site, name string) error {
	return c.Delete(sprawl.EPGroup, sprawl.Request{
		"site": site,
		"name": name},
	)
}

// ListGroups on a site.
func (c *SprawlClient) ListGroups(site string) (sprawl.GroupList, error) {
	var list sprawl.GroupList
	data, err := c.Get(sprawl.EPGroups, sprawl.Request{"site": site})
	if err != nil {
		return list, err
	}

	err = json.Unmarshal(data, &list)
	return list, err
}

// ListGroupMembers lists members of a group on a site.
func (c *SprawlClient) ListGroupMembers(site, group string) (sprawl.UserList, error) {
	var list sprawl.UserList
	data, err := c.Get(sprawl.EPGroup+sprawl.EPMembers, sprawl.Request{
		"site": site,
		"name": group,
	})
	if err != nil {
		return list, err
	}

	err = json.Unmarshal(data, &list)
	return list, err
}

// AddGroupMember to a group on a site.
func (c *SprawlClient) AddGroupMember(site, group, name string) error {
	_, err := c.Post(sprawl.EPGroup+sprawl.EPMember, sprawl.Request{
		"site": site,
		"name": group,
		"user": name,
	})
	return err
}

// RemoveGroupMember from a group on a site.
func (c *SprawlClient) RemoveGroupMember(site, group, name string) error {
	return c.Delete(sprawl.EPGroup+sprawl.EPMember, sprawl.Request{
		"site": site,
		"name": group,
		"user": name,
	})
}

// AddGroupPermission to a group on a site.
func (c *SprawlClient) AddGroupPermission(site, group, name string) error {
	_, err := c.Post(sprawl.EPGroup+sprawl.EPPermission, sprawl.Request{
		"site":       site,
		"group":      group,
		"permission": name,
	})
	return err
}

// RemoveGroupPermission from a group on a site.
func (c *SprawlClient) RemoveGroupPermission(site, group, name string) error {
	return c.Delete(sprawl.EPGroup+sprawl.EPPermission, sprawl.Request{
		"site":       site,
		"group":      group,
		"permission": name,
	})
}

// ListGroupPermissions of a group on a site.
func (c *SprawlClient) ListGroupPermissions(site, group string) (sprawl.PermissionList, error) {
	var list sprawl.PermissionList
	data, err := c.Get(sprawl.EPGroup+sprawl.EPPermissions, sprawl.Request{
		"site":  site,
		"group": group,
	})
	if err != nil {
		return list, err
	}

	err = json.Unmarshal(data, &list)
	return list, err
}

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
func (c *SprawlClient) GetPermission(name string) (sprawl.Permission, error) {
	var p sprawl.Permission
	data, err := c.Get(sprawl.EPPermission, sprawl.Request{"name": name})
	if err != nil {
		return p, err
	}

	err = json.Unmarshal(data, &p)
	return p, err
}

// UpdatePermission updates a permission with a new description.
func (c *SprawlClient) UpdatePermission(name, description string) error {
	err := c.Put(sprawl.EPPermission, sprawl.Request{
		"name":        name,
		"description": description,
	})
	return err
}

// DeletePermission deletes a permission.
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
