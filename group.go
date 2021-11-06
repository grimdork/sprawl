//	Copyright 2021 Ronny Bangsund
//
//	This software is released under the MIT License.
//	https://opensource.org/licenses/MIT

// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package sprawl

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// Group has a name and site it belongs to.
type Group struct {
	ID   int64
	Name string
	Site string
}

// CreateGroup creates a new group.
func (db *Database) CreateGroup(name, site string) error {
	sql := `insert into groups (name, sid) select $1,s.id from sites s where s.name=$2;`
	_, err := db.Exec(context.Background(), sql, name, site)
	return err
}

// DeleteGroup deletes a group.
func (db *Database) DeleteGroup(name, site string) error {
	sql := "delete from groups where name=$1 and sid=(select id from sites where name=$2);"
	_, err := db.Exec(context.Background(), sql, name, site)
	return err
}

// DeleteGroupByID deletes a group by ID.
func (db *Database) DeleteGroupByID(id int64) error {
	sql := "delete from groups where id = $1;"
	_, err := db.Exec(context.Background(), sql, id)
	return err
}

// GetGroupByID returns a group by ID.
func (db *Database) GetGroupByID(id int64) (Group, error) {
	g := Group{}
	sql := "select id, name, site from groups where id=$1 "
	err := db.QueryRow(context.Background(), sql, id).Scan(&g.ID, &g.Name, &g.Site)
	return g, err
}

// GetGroups returns many groups.
func (db *Database) GetGroups(start, max int64, site string) ([]Group, error) {
	var err error
	var rows pgx.Rows

	if site == "" {
		sql := "select groups.id,groups.name,sites.name from groups inner join sites on groups.sid=sites.id order by sites.id,groups.id asc limit $1 offset $2;"
		rows, err = db.Query(context.Background(), sql, max, start)
	} else {
		sql := "select groups.id,groups.name,sites.name from groups inner join sites on groups.sid=sites.id where sites.name=$3 order by sites.id,groups.id asc limit $1 offset $2;"
		rows, err = db.Query(context.Background(), sql, max, start, site)
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var groups []Group
	for rows.Next() {
		g := Group{}
		err := rows.Scan(&g.ID, &g.Name, &g.Site)
		if err != nil {
			return nil, err
		}
		groups = append(groups, g)
	}

	return groups, nil
}
