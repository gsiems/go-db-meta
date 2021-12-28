package dbms

import (
	"github.com/gsiems/go-db-meta/dbms/mariadb"
	"github.com/gsiems/go-db-meta/dbms/mssql"
	"github.com/gsiems/go-db-meta/dbms/ora"
	"github.com/gsiems/go-db-meta/dbms/pg"
	"github.com/gsiems/go-db-meta/dbms/sqlite"
	m "github.com/gsiems/go-db-meta/model"
)

// Indexes returns a slice of Indexes for the (schemaName, tableName) parameters
func (db *DBMS) Indexes(schemaName, tableName string) ([]m.Index, error) {

	var d []m.Index

	switch db.id {
	case PostgreSQL:
		return pg.Indexes(db.Connection, schemaName, tableName)
	case SQLite:
		return sqlite.Indexes(db.Connection, schemaName, tableName)
	case MariaDB, MySQL:
		return mariadb.Indexes(db.Connection, schemaName, tableName)
	case Oracle:
		return ora.Indexes(db.Connection, schemaName, tableName)
	case MSSQL:
		return mssql.Indexes(db.Connection, schemaName, tableName)
	}

	return d, unsupportedDBErr(db.id)
}
