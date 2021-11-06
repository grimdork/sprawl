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

func (db *Database) createPermission(name, description string) error {
	sql := `insert into permissions (name,description) values($1,$2);`
	_, err := db.Exec(context.Background(), sql, name, description)
	return err
}

func (db *Database) deletePermission(name string) error {
	sql := `delete from permissions where name=$1;`
	_, err := db.Exec(context.Background(), sql, name)
	return err
}

func (db *Database) listPermissions() ([]Permission, error) {
	sql := `select name,description from permissions;`
	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var permissions []Permission
	for rows.Next() {
		var name, description string
		err = rows.Scan(&name, &description)
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, Permission{Name: name, Description: description})
	}
	return permissions, nil
}
