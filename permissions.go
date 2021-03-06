package sprawl

import "context"

// Permission is just a keyword with actual meaning defined by a client app.
type Permission struct {
	// ID is uthe unique identifier.
	ID int64
	// Name is the keyword of the permission.
	Name string
	// Description is optional.
	Description string
}

// PermissionList is a list of permissions.
type PermissionList struct {
	Permissions []Permission `json:"permissions"`
}

// CreatePermission adds a new keyword and optional description.
func (db *Database) CreatePermission(name, description string) error {
	sql := `insert into permissions (name,description) values($1,$2);`
	_, err := db.Exec(context.Background(), sql, name, description)
	return err
}

// UpdatePermission updates a permission description.
func (db *Database) UpdatePermission(name, description string) error {
	sql := `update permissions set description=$2 where name=$';`
	_, err := db.Exec(context.Background(), sql, name, description)
	return err
}

// DeletePermission deletes a permission keyword.
func (db *Database) DeletePermission(name string) error {
	sql := `delete from permissions where name=$1;`
	_, err := db.Exec(context.Background(), sql, name)
	return err
}

// GetPermission returns a permission keyword and its description.
func (db *Database) GetPermission(name string) (Permission, error) {
	sql := `select id,name,description from permissions where name=$1;`
	var p Permission
	err := db.QueryRow(context.Background(), sql, name).Scan(&p.ID, &p.Name, &p.Description)
	return p, err
}

// GetPermissions returns all permission keywords and their descriptions.
func (db *Database) GetPermissions() (PermissionList, error) {
	sql := `select id,name,description from permissions;`
	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		return PermissionList{}, err
	}

	defer rows.Close()
	var list PermissionList
	for rows.Next() {
		var p Permission
		err = rows.Scan(&p.ID, &p.Name, &p.Description)
		if err != nil {
			return PermissionList{}, err
		}
		list.Permissions = append(list.Permissions, p)
	}
	return list, nil
}

// Has the user got a specific permission?
func (db *Database) Has(name, permission string) bool {
	sql := `select count(p.id) from users as u
		inner join permissions as p on p.name=$2
		inner join roles as r on r.pid=p.id
		inner join groups as g on g.id=r.gid
		inner join sites as s on s.id=g.sid
		inner join members as m on m.uid=u.id
		where u.name=$1`
	var count int
	_ = db.QueryRow(context.Background(), sql, name, permission).Scan(&count)
	return count > 0
}

// IsSiteAdmin returns true if the user is a site admin for the specified site.
func (db *Database) IsSiteAdmin(name, site string) bool {
	if name == "admin" {
		// The superadmin is site admin everywhere.
		return true
	}

	sql := `select count(profiles.uid) from profiles
		inner join users on profiles.uid=users.id
		inner join sites on profiles.sid=sites.id
		where users.name=$1 and sites.name=$2;`
	var count int64
	err := db.QueryRow(context.Background(), sql, name, site).Scan(&count)
	if err != nil {
		return false
	}

	return count > 0
}
