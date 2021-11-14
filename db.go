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

-- The user table holds only the minimum required for logins, plus a freeform column for JSON data or similar.
create table if not exists users (
	id serial primary key,
	name text unique not null,
	fullname text not null default ''::text,
	email text not null default ''::text,
	password text not null,
	data text
);

-- Sites are very simple - just somewhere to reference the name for other tables.
create table if not exists sites (
	id serial primary key,
	name text unique not null
);

-- Profiles is where the site-specific userdata is stored. JSON is recommended.
create table if not exists profiles (
	uid bigint not null,
	sid bigint not null,
	data text not null default ''::text,
	admin boolean not null default false,
	constraint profile_uid foreign key(uid) references users(id) on delete cascade,
	constraint profile_sid foreign key(Sid) references sites(id) on delete cascade
);

-- Users may only have one profile per site.
create unique index if not exists idx_userprofile on profiles(uid,sid);

--cGroups represent a role for users of a site.
create table if not exists groups (
	id serial primary key,
	name text not null,
	sid bigint not null,
	constraint group_sid foreign key(sid) references sites(id) on delete cascade
);

create unique index if not exists idx_groupsite on groups(id,sid);

-- Memberships are the many-to-many relationship between users and groups.
create table if not exists members (
	uid bigint not null,
	gid bigint not null,
	constraint members_uid foreign key(uid) references users(id) on delete cascade,
	constraint members_gid foreign key(gid) references groups(id) on delete cascade
);

create unique index if not exists idx_gmembersgroup on members(uid,gid);

-- Tokens are created when authenticating.
create table if not exists tokens (
	hash text primary key,
	uid bigint,
	expires bigint,
	constraint fk_tokens_uid foreign key(uid) references users(id) on delete cascade
);

-- Permissions are attached to groups, and are configurable as you like. Sprawl will create some defaults.
create table if not exists permissions (
	id serial primary key,
	name text unique not null,
	description text not null default ''::text
);

-- Roles match permissions to groups.
create table if not exists roles (
	gid bigint not null,
	pid bigint not null,
	constraint role_gid foreign key(gid) references groups(id) on delete cascade,
	constraint role_pid foreign key(pid) references permissions(id) on delete cascade
);

-- Ensure a GID/PID pair only appears once.
create unique index if not exists idx_roles on roles(gid,pid);

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
