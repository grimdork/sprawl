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
	sql := "insert into sites (name) values ($1)"
	_, err := db.Pool.Exec(context.Background(), sql, name)
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
