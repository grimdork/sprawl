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
	"strconv"

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
	sql := `insert into groups (name,sid) select $1,s.id from sites s where s.name=$2;`
	_, err := db.Exec(context.Background(), sql, name, site)
	return err
}

// DeleteGroup deletes a group.
func (db *Database) DeleteGroup(name, site string) error {
	var sql string
	gid, _ := strconv.ParseInt(name, 10, 64)
	sid, _ := strconv.ParseInt(site, 10, 64)
	if gid == 0 {
		if sid == 0 {
			sql = "delete from groups where name=$1 and sid=(select id from sites where name=$2);"
		} else {
			sql = "delete from groups where name=$1 and sid=$2;"
		}
	} else {
		if sid == 0 {
			sql = "delete from groups where id=$1 and sid=(select id from sites where name=$2);"
		} else {
			sql = "delete from groups where id=$1 and sid=$2;"
		}
	}
	_, err := db.Exec(context.Background(), sql, name, site)
	return err
}

// DeleteGroupByID deletes a group by ID.
func (db *Database) DeleteGroupByID(id int64) error {
	sql := "delete from groups where id = $1;"
	_, err := db.Exec(context.Background(), sql, id)
	return err
}

// GetGroups returns many groups.
func (db *Database) GetGroups(start, max int64, site string) (GroupList, error) {
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
		return GroupList{}, err
	}

	defer rows.Close()
	var list GroupList
	for rows.Next() {
		g := Group{}
		err := rows.Scan(&g.ID, &g.Name, &g.Site)
		if err != nil {
			return GroupList{}, err
		}
		list.Groups = append(list.Groups, g)
	}

	return list, nil
}
