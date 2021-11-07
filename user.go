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

	"golang.org/x/crypto/bcrypt"
)

// User data contains the barebones login info and site membership.
type User struct {
	// ID of the user (unique).
	ID int64
	// Name may be an e-mail address.
	Name string
	// Password is a bcrypt hash.
	Password string
	// Profiles for different sites.
	Profiles map[string]Profile
}

// CreateUser in database.
func (db *Database) CreateUser(name, password string) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	if name == "admin" {
		_, err = db.Exec(context.Background(), "insert into users (id,name,password) values (1,$1,$2);", name, pass)
	} else {
		_, err = db.Exec(context.Background(), "insert into users(name,password) values($1,$2)", name, pass)
	}
	return err
}

// DeleteUser deletes a user from the database.
func (db *Database) DeleteUser(name string) error {
	_, err := db.Exec(context.Background(), "delete from users where name=$1;", name)
	return err
}

// GetUser by name.
func (db *Database) GetUser(name string) *User {
	u := &User{}
	err := db.QueryRow(context.Background(), "select id,name,password from users where name=$1 limit 1;", name).Scan(&u.ID, &u.Name, &u.Password)
	if err != nil {
		return nil
	}

	return u
}

// SetPassword sets the password.
func (db *Database) SetPassword(u *User, password string) error {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	u.Password = string(pw)
	_, err = db.Exec(context.Background(), "update users set password=$1 where id=$2; limit 1", u.Password, u.ID)
	return err
}

// CheckPassword checks the password.
func (u *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

// GetUsers returns a range of users.
func (db *Database) GetUsers(start, max int64) ([]User, error) {
	rows, err := db.Query(context.Background(), "select id,name,password from users order by id asc limit $1 offset $2;", max, start)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []User
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.ID, &u.Name, &u.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}
	return users, nil
}
