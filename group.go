//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package sprawl

import "context"

// Group has a name and site it belongs to.
type Group struct {
	ID   int64
	Name string
	Site string
}

// CreateGroup creates a new group.
func (db *Database) CreateGroup(name string, site string) error {
	sql := "insert into groups (name, site) values ($1, $2)"
	_, err := db.Exec(context.Background(), sql, name, site)
	return err
}

// DeleteGroup deletes a group.
func (db *Database) DeleteGroup(name string) error {
	sql := "delete from groups where name = $1 limit 1"
	_, err := db.Exec(context.Background(), sql, name)
	return err
}

// DeleteGroupByID deletes a group by ID.
func (db *Database) DeleteGroupByID(id int64) error {
	sql := "delete from groups where id = $1 limit 1"
	_, err := db.Exec(context.Background(), sql, id)
	return err
}

// GetGroup returns a group by name.
func (db *Database) GetGroup(name string) (Group, error) {
	g := Group{}
	sql := "select id, name, site from groups where name = $1 limit 1"
	err := db.QueryRow(context.Background(), sql, name).Scan(&g.ID, &g.Name, &g.Site)
	return g, err
}

// GetGroupByID returns a group by ID.
func (db *Database) GetGroupByID(id int64) (Group, error) {
	g := Group{}
	sql := "select id, name, site from groups where id = $1 limit 1"
	err := db.QueryRow(context.Background(), sql, id).Scan(&g.ID, &g.Name, &g.Site)
	return g, err
}

// GetGroups returns all groups.
func (db *Database) GetGroups() []Group {
	sql := "select id, name, site from groups"
	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		return nil
	}

	groups := []Group{}
	for rows.Next() {
		g := Group{}
		err := rows.Scan(&g.ID, &g.Name, &g.Site)
		if err != nil {
			continue
		}

		groups = append(groups, g)
	}

	return groups
}

// GetGroupsBySite returns all groups by site.
func (db *Database) GetGroupsBySite(site string) []Group {
	sql := "select id, name, site from groups where site = $1"
	rows, err := db.Query(context.Background(), sql, site)
	if err != nil {
		return nil
	}

	groups := []Group{}
	for rows.Next() {
		g := Group{}
		err := rows.Scan(&g.ID, &g.Name, &g.Site)
		if err != nil {
			continue
		}

		groups = append(groups, g)
	}

	return groups
}
