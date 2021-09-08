// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package sprawl

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

var setupSQL string = `begin;

create table if not exists users (
	id serial primary key,
	name text unique not null,
	password text not null
);

create table if not exists sites (
	id serial primary key,
	name text unique not null
);

create table if not exists profiles (
	uid bigint not null,
	sid bigint not null,
	data text,
	constraint profile_uid foreign key(uid) references users(id) on delete cascade,
	constraint profile_sid foreign key(Sid) references sites(id) on delete cascade
);

create index if not exists idx_userprofile on profiles(uid,sid);

create table if not exists groups (
	id serial primary key,
	name text not null,
	sid bigint not null,
	constraint group_sid foreign key(sid) references sites(id) on delete cascade
);

create index if not exists idx_groupsite on groups(id,sid);

create table if not exists members (
	uid bigint not null,
	gid bigint not null,
	constraint members_uid foreign key(uid) references users(id) on delete cascade,
	constraint members_gid foreign key(gid) references groups(id) on delete cascade
);

create index if not exists idx_gmembersgroup on members(uid,gid);

create table if not exists tokens (
	hash text primary key,
	uid bigint,
	expires bigint,
	constraint fk_tokens_uid foreign key(uid) references users(id) on delete cascade
);

commit;
`

// Database pools for Sprawl.
type Database struct {
	*pgxpool.Pool
}

// NewDatabase connects to the database and creates a connection pool.
func NewDatabase(conn string) (*Database, error) {
	db := &Database{}
	var err error
	db.Pool, err = pgxpool.Connect(context.Background(), conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CreateTables sets up the database.
func (db *Database) CreateTables() error {
	_, err := db.Exec(context.Background(), setupSQL)
	return err
}
