package dbms

import (
	"github.com/gsiems/go-db-meta/dbms/mariadb"
	"github.com/gsiems/go-db-meta/dbms/mssql"
	"github.com/gsiems/go-db-meta/dbms/ora"
	"github.com/gsiems/go-db-meta/dbms/pg"
	"github.com/gsiems/go-db-meta/dbms/sqlite"
	m "github.com/gsiems/go-db-meta/model"
)

// Schemata returns a slice of Schemas, optionally filtered on the (nclude, xclude) parameters
func (db *DBMS) Schemata(nclude, xclude string) ([]m.Schema, error) {

	var d []m.Schema

	switch db.id {
	case PostgreSQL:
		return pg.Schemata(db.Connection, nclude, xclude)
	case SQLite:
		return sqlite.Schemata(db.Connection, nclude, xclude)
	case MariaDB, MySQL:
		return mariadb.Schemata(db.Connection, nclude, xclude)
	case Oracle:
		return ora.Schemata(db.Connection, nclude, xclude)
	case MSSQL:
		return mssql.Schemata(db.Connection, nclude, xclude)
	}

	return d, unsupportedDBErr(db.id)
}
