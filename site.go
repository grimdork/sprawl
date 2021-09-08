// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package sprawl

import "context"

// Addsite to the database.
func (db *Database) AddSite(name string) error {
	sql := "insert into sites (name) values ($1)"
	_, err := db.Pool.Exec(context.Background(), sql, name)
	return err
}

// RemoveSite removes a site from the database.
func (db *Database) RemoveSite(name string) error {
	sql := "delete from sites where name = $1"
	_, err := db.Pool.Exec(context.Background(), sql, name)
	return err
}

// GetSiteID from a name.
func (db *Database) GetSiteID(name string) (int64, error) {
	sql := "select id from sites where name = $1"
	var id int64
	err := db.Pool.QueryRow(context.Background(), sql, name).Scan(&id)
	return id, err
}

// GetAllSites in a string slice.
func (db *Database) GetAllSites() []string {
	sql := "select name from sites"
	rows, err := db.Pool.Query(context.Background(), sql)
	if err != nil {
		return nil
	}
	defer rows.Close()

	sites := []string{}
	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			return nil
		}

		sites = append(sites, s)
	}
	return sites
}
