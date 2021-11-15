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
	ID int64 `json:"id"`
	// Name may be an e-mail address.
	Name string `json:"name"`
	// Fullname is the user's full name.
	Fullname string `json:"fullname,omitempty"`
	// Email is the user's e-mail address.
	Email string `json:"email"`
	// Password is a bcrypt hash.
	Password string `json:"password,omitempty"`
	// Profiles for different sites.
	Profile string `json:"profile,omitempty"`
	// Admin is true if the user is an admin.
	Admin bool `json:"admin,omitempty"`
}

// UserList is returned from lookup.
type UserList struct {
	// Users is a list of users, possibly limited by search parameters.
	Users []User `json:"users"`
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

// UpdateUser with new details.
func (db *Database) UpdateUser(name, newname, fullname, email string) error {
	_, err := db.Exec(context.Background(), "update users set name=$2,fullname=$3,email=$4 where name=$1", name, newname, fullname, email)
	return err
}

// GetUser by name.
func (db *Database) GetUser(name string) (User, error) {
	u := User{}
	err := db.QueryRow(context.Background(), "select id,name,fullname,email,password from users where name=$1 limit 1;", name).Scan(&u.ID, &u.Name, &u.Fullname, &u.Email, &u.Password)
	if err != nil {
		return u, err
	}

	return u, nil
}

// SetPassword sets the password.
func (db *Database) SetPassword(name, password string, c int) error {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), c)
	if err != nil {
		return err
	}

	sql := `update users set password=$1
		where id=(select id from users where name=$2 limit 1)`
	_, err = db.Exec(context.Background(), sql, pw, name)
	return err
}

// CheckPassword checks the password.
func (u *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

// GetUsers returns a range of users.
func (db *Database) GetUsers(start, max int64) ([]User, error) {
	rows, err := db.Query(context.Background(), "select id,name,fullname,email from users order by id asc limit $1 offset $2;", max, start)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []User
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.ID, &u.Name, &u.Fullname, &u.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}
	return users, nil
}
