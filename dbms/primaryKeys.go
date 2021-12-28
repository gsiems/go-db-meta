package dbms

import (
	"github.com/gsiems/go-db-meta/dbms/mariadb"
	"github.com/gsiems/go-db-meta/dbms/mssql"
	"github.com/gsiems/go-db-meta/dbms/ora"
	"github.com/gsiems/go-db-meta/dbms/pg"
	"github.com/gsiems/go-db-meta/dbms/sqlite"
	m "github.com/gsiems/go-db-meta/model"
)

// PrimaryKeys returns a slice of primary keys for the (schemaName, tableName) parameters
func (db *DBMS) PrimaryKeys(schemaName, tableName string) ([]m.PrimaryKey, error) {

	var d []m.PrimaryKey

	switch db.id {
	case PostgreSQL:
		return pg.PrimaryKeys(db.Connection, schemaName, tableName)
	case SQLite:
		return sqlite.PrimaryKeys(db.Connection, schemaName, tableName)
	case MariaDB, MySQL:
		return mariadb.PrimaryKeys(db.Connection, schemaName, tableName)
	case Oracle:
		return ora.PrimaryKeys(db.Connection, schemaName, tableName)
	case MSSQL:
		return mssql.PrimaryKeys(db.Connection, schemaName, tableName)
	}

	return d, unsupportedDBErr(db.id)
}
