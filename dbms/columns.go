package dbms

import (
	"github.com/gsiems/go-db-meta/dbms/mariadb"
	"github.com/gsiems/go-db-meta/dbms/mssql"
	"github.com/gsiems/go-db-meta/dbms/ora"
	"github.com/gsiems/go-db-meta/dbms/pg"
	"github.com/gsiems/go-db-meta/dbms/sqlite"
	m "github.com/gsiems/go-db-meta/model"
)

// Columns returns a slice of Columns for the (schemaName, tableName) parameters
func (db *DBMS) Columns(schemaName, tableName string) ([]m.Column, error) {

	var d []m.Column

	switch db.id {
	case PostgreSQL:
		return pg.Columns(db.Connection, schemaName, tableName)
	case SQLite:
		return sqlite.Columns(db.Connection, schemaName, tableName)
	case MariaDB, MySQL:
		return mariadb.Columns(db.Connection, schemaName, tableName)
	case Oracle:
		return ora.Columns(db.Connection, schemaName, tableName)
	case MSSQL:
		return mssql.Columns(db.Connection, schemaName, tableName)
	}

	return d, unsupportedDBErr(db.id)
}
