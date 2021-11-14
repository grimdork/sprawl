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

// AddGroupMember adds a user to a group.
func (db *Database) AddGroupMember(site, group, name string) error {
	sql := `insert into members (uid,gid)
		values(
			(select u.id from users u where u.name=$3),
			(select g.id from groups g where g.name=$2 and
			g.sid=(select s.id from sites s where s.name=$1))
		)`
	_, err := db.Pool.Exec(context.Background(), sql, site, group, name)
	return err
}

// RemoveGroupMember removes a user from a group.
func (db *Database) RemoveGroupMember(site, group, name string) error {
	sql := `delete from members where
		uid=(select u.id from users u where u.name=$3) and
		gid=(select g.id from groups g where g.name=$2 and
			g.sid=(select s.id from sites s where s.name=$1))`
	_, err := db.Pool.Exec(context.Background(), sql, site, group, name)
	return err
}

// GetGroupMembers returns a UserList.
func (db *Database) GetGroupMembers(site, group string) (UserList, error) {
	sql := `select users.id,users.name,profiles.admin from users
	inner join profiles on profiles.uid=users.id
	inner join sites on sites.id=profiles.sid;`
	rows, err := db.Pool.Query(context.Background(), sql)
	if err != nil {
		return UserList{}, err
	}
	defer rows.Close()

	var users UserList
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Name, &u.Admin)
		if err != nil {
			return UserList{}, err
		}

		users.Users = append(users.Users, u)
	}
	return users, nil
}

func (db *Database) GetGroupPermissions(site, group string) (PermissionList, error) {
	sql := `select permissions.id,permissions.name,permissions.description from permissions
	inner join roles on roles.pid=permissions.id
	inner join groups on groups.id=roles.gid
	inner join sites on sites.id=groups.sid
	where sites.name=$1 and groups.name=$2`
	rows, err := db.Pool.Query(context.Background(), sql, site, group)
	if err != nil {
		return PermissionList{}, err
	}

	defer rows.Close()
	var permissions PermissionList
	for rows.Next() {
		var p Permission
		err = rows.Scan(&p.ID, &p.Name, &p.Description)
		if err != nil {
			return PermissionList{}, err
		}
		permissions.Permissions = append(permissions.Permissions, p)
	}
	return permissions, nil
}

// AddGroupPermission adds a permission to a group.
func (db *Database) AddGroupPermission(site, group, name string) error {
	sql := `insert into roles (gid,pid)
		values(
			(select g.id from groups g where g.name=$2 and
			g.sid=(select s.id from sites s where s.name=$1)),
			(select p.id from permissions p where p.name=$3)
		)`
	_, err := db.Pool.Exec(context.Background(), sql, site, group, name)
	return err
}

func (db *Database) RemoveGroupPermission(site, group, name string) error {
	sql := `delete from roles where
		gid=(select g.id from groups g where g.name=$2 and
		g.sid=(select s.id from sites s where s.name=$1)) and
		pid=(select p.id from permissions p where p.name=$3)`
	_, err := db.Pool.Exec(context.Background(), sql, site, group, name)
	return err
}
