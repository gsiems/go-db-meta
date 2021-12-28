package dbms

import (
	"database/sql"

	"github.com/gsiems/go-db-meta/dbms/mariadb"
	"github.com/gsiems/go-db-meta/dbms/mssql"
	"github.com/gsiems/go-db-meta/dbms/mysql"
	"github.com/gsiems/go-db-meta/dbms/ora"
	"github.com/gsiems/go-db-meta/dbms/pg"
	"github.com/gsiems/go-db-meta/dbms/sqlite"
)

const (
	NullDB     = iota // zero, so no "default" database allowed
	PostgreSQL = iota
	SQLite     = iota
	MariaDB    = iota
	MySQL      = iota
	Oracle     = iota
	MSSQL      = iota
)

type DBMS struct {
	Connection *sql.DB
	id         int
}

func Init(db *sql.DB, id int) (DBMS, error) {

	var d DBMS

	switch id {
	case PostgreSQL, SQLite, MariaDB, MySQL, Oracle, MSSQL:
		d.id = id
		d.Connection = db
	default:
		return d, unsupportedDBErr(id)
	}

	return d, db.Ping()
}

// CloseDB closes a DB reference
func (db *DBMS) CloseDB() error {
	return db.Connection.Close()
}

func (db *DBMS) Name() string {
	switch db.id {
	case PostgreSQL:
		return pg.Name()
	case SQLite:
		return sqlite.Name()
	case MariaDB:
		return mariadb.Name()
	case MySQL:
		return mysql.Name()
	case Oracle:
		return ora.Name()
	case MSSQL:
		return mssql.Name()
	}
	return ""
}

func (db *DBMS) ID() int {
	switch db.id {
	case PostgreSQL, SQLite, MariaDB, MySQL, Oracle, MSSQL:
		return db.id
	}
	return NullDB
}
