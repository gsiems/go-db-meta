package model

import (
	"database/sql"
)

type ConnectInfo struct {
	Host     string
	Port     string
	DbName   string
	Username string
	Password string
	File     string
}

// DB contains an database/sql connection
type DB struct {
	*sql.DB
}

// BindConnection binds a database/sql connection to the model. This
// should allow the model to be database driver agnostic
func BindConnection(db *sql.DB) (DB, error) {
	return DB{db}, db.Ping()
}

// CloseDB closes a DB reference
func (db *DB) CloseDB() error {
	return db.DB.Close()
}
