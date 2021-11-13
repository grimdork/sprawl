// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package sprawl

import (
	"context"
	"strconv"
)

// Site is the struct for a domain.
type Site struct {
	// ID is auto-generated on insert.
	ID int64
	// Name is the domain name.
	Name string
}

// Createsite in the database.
func (db *Database) CreateSite(name string) error {
	var err error
	if name == "system" {
		_, err = db.Pool.Exec(context.Background(), "insert into sites (id,name) values (1,$1)", name)
	} else {
		_, err = db.Pool.Exec(context.Background(), "insert into sites (name) values ($1)", name)
	}
	return err
}

// DeleteSite from the database.
func (db *Database) DeleteSite(name string) error {
	_, err := strconv.ParseInt(name, 10, 64)
	if err == nil {
		sql := "delete from sites where id = $1"
		_, err = db.Pool.Exec(context.Background(), sql, name)
		return err
	}

	sql := "delete from sites where name = $1"
	_, err = db.Pool.Exec(context.Background(), sql, name)
	return err
}

// GetSiteID from a name.
func (db *Database) GetSiteID(name string) (int64, error) {
	sql := "select id from sites where name = $1"
	var id int64
	err := db.Pool.QueryRow(context.Background(), sql, name).Scan(&id)
	return id, err
}

// GetSites in a string slice.
func (db *Database) GetSites(start, max int64) ([]Site, error) {
	sql := "select id,name from sites order by name limit $1 offset $2"
	rows, err := db.Pool.Query(context.Background(), sql, max, start)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sites []Site
	for rows.Next() {
		var s Site
		err = rows.Scan(&s.ID, &s.Name)
		if err != nil {
			return nil, err
		}

		sites = append(sites, s)
	}
	return sites, nil
}

// GetSiteMembers returns a slice of users for a site.
func (db *Database) GetSiteMembers(site string) ([]User, error) {
	sql := `select users.id,users.name,profiles.admin from users
	inner join profiles on profiles.uid=users.id
	inner join sites on sites.id=profiles.sid;`
	rows, err := db.Pool.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Name, &u.Admin)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}
	return users, nil
}

// AddSiteMember adds a profile for a user.
func (db *Database) AddSiteMember(site, name, data, admin string) error {
	sql := `insert into profiles (sid,uid,data,admin)
		values(
			(select s.id from sites s where s.name=$1),
			(select u.id from users u where u.name=$2),
			$3,$4
		)`
	_, err := db.Pool.Exec(context.Background(), sql, site, name, data, admin)
	return err
}

// RemoveSiteMember deletes a profile for a user.
func (db *Database) RemoveSiteMember(site, name string) error {
	sql := `delete from profiles where
		sid=(select s.id from sites s where s.name=$1) and
		uid=(select u.id from users u where u.name=$2)`
	_, err := db.Pool.Exec(context.Background(), sql, site, name)
	return err
}

// ToggleSiteAdmin status of a user.
func (db *Database) ToggleSiteAdmin(site, name string, on bool) error {
	sql := `update profiles set admin=$3 where
		sid=(select s.id from sites s where s.name=$1) and
		uid=(select u.id from users u where u.name=$2)`
	_, err := db.Pool.Exec(context.Background(), sql, site, name, on)
	return err
}
