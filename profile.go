// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package sprawl

import "context"

// Prfile for a user+site combination.
type Profile struct {
	// Site the profile is for.
	Site string
	// Data is a site-specific JSON structure.
	Data string
	// Groups for this site.
	Groups map[string]interface{}
}

// CreateProfile both adds a user to a site and creates its site-specific profile entry.
func (db *Database) CreateProfile(username, site, data string, admin bool) error {
	sql := `insert into  profiles (uid,sid,data,admin)
	values (
		(select users.id from users where users.name=$1),
		(select sites.id from sites where sites.name=$2),
		$3,$4);`
	_, err := db.Exec(context.Background(), sql, username, site, data, admin)
	return err
}

// UpdateProfile updates the profile for a user+site combination.
func (db *Database) UpdateProfile(username, site, data string) error {
	sql := `update profiles set data=$3 where
	uid=(select users.id from users where users.name=$1) and
	sid=(select sites.id from sites where sites.name=$2);`
	_, err := db.Exec(context.Background(), sql, username, site, data)
	return err
}

// RemoveProfile removes a user from a site.
func (db *Database) RemoveProfile(username, site string) error {
	sql := `delete from profiles where
	uid=(select users.id from users where users.name=$1) and
	sid=(select sites.id from sites where sites.name=$2);`
	_, err := db.Exec(context.Background(), sql, username, site)
	return err
}
