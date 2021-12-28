package dbms

import (
	"github.com/gsiems/go-db-meta/dbms/mariadb"
	"github.com/gsiems/go-db-meta/dbms/mssql"
	"github.com/gsiems/go-db-meta/dbms/ora"
	"github.com/gsiems/go-db-meta/dbms/pg"
	"github.com/gsiems/go-db-meta/dbms/sqlite"
	m "github.com/gsiems/go-db-meta/model"
)

// Tables returns a slice of Tables for the (schema) parameter
func (db *DBMS) Tables(schemaName string) ([]m.Table, error) {

	var d []m.Table

	switch db.id {
	case PostgreSQL:
		return pg.Tables(db.Connection, schemaName)
	case SQLite:
		return sqlite.Tables(db.Connection, schemaName)
	case MariaDB, MySQL:
		return mariadb.Tables(db.Connection, schemaName)
	case Oracle:
		return ora.Tables(db.Connection, schemaName)
	case MSSQL:
		return mssql.Tables(db.Connection, schemaName)
	}

	return d, unsupportedDBErr(db.id)
}
