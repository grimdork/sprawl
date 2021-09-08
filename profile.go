// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package sprawl

// Prfile for a user+site combination.
type Profile struct {
	// Data is a site-specific JSON structure.
	Data string
	// Groups and permissions for this site.
	Groups map[string]int64
}

// AddMembership adds a user to a site.
func (db *Database) AddMembership(username, site string) error {
	return nil
}

// RemoveMembership removes a user from a site.
func (db *Database) RemoveMembership(username, site string) error {
	return nil
}
