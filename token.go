package sprawl

import (
	"context"
	"crypto/sha512"
	"fmt"
	"time"
)

// GenerateToken if needed.
func (db *Database) GenerateToken(name string) string {
	sql := "select hash,expires from tokens inner join users on users.id=tokens.uid where users.name=$1"
	var token string
	var expires int64
	err := db.QueryRow(context.Background(), sql, name).Scan(&token, &expires)
	ex := time.Unix(expires, 0)
	if err == nil && ex.After(time.Now()) {
		return token
	}

	// Delete all expired tokens for this user.
	_, _ = db.Exec(context.Background(), "delete from tokens where uid=(select id from users where name=$1)", name)

	// And finally generate a new one.
	s := RandString(32)
	h := sha512.New()
	h.Write([]byte(s))
	token = fmt.Sprintf("%x", h.Sum(nil))
	t := time.Now().Add(time.Hour * 8)
	expires = t.Unix()
	sql = "insert into tokens (uid,hash,expires) values ((select id from users where name=$1),$2,$3)"
	_, err = db.Exec(context.Background(), sql, name, token, expires)
	if err != nil {
		// No point in returning a token if it can't be looked up.
		return ""
	}

	return token
}

// VerifyToken against expiry.
func (db *Database) VerifyToken(username, token string) bool {
	sql := "select expires from tokens inner join users on users.id=tokens.uid where users.name=$1 and hash=$2"
	var expires int64
	err := db.QueryRow(context.Background(), sql, username, token).Scan(&expires)
	ex := time.Unix(expires, 0)
	if err != nil || ex.Before(time.Now()) {
		return false
	}

	return true
}
