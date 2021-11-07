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

// DeletePermission deletes a permission keyword.
func (db *Database) DeletePermission(name string) error {
	sql := `delete from permissions where name=$1;`
	_, err := db.Exec(context.Background(), sql, name)
	return err
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
