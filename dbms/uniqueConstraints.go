package dbms

import (
	"github.com/gsiems/go-db-meta/dbms/mariadb"
	"github.com/gsiems/go-db-meta/dbms/mssql"
	"github.com/gsiems/go-db-meta/dbms/ora"
	"github.com/gsiems/go-db-meta/dbms/pg"
	"github.com/gsiems/go-db-meta/dbms/sqlite"
	m "github.com/gsiems/go-db-meta/model"
)

// UniqueConstraints returns a slice of Unique Constraints for the
// (schemaName, tableName) parameters
func (db *DBMS) UniqueConstraints(schemaName, tableName string) ([]m.UniqueConstraint, error) {

	var d []m.UniqueConstraint

	switch db.id {
	case PostgreSQL:
		return pg.UniqueConstraints(db.Connection, schemaName, tableName)
	case SQLite:
		return sqlite.UniqueConstraints(db.Connection, schemaName, tableName)
	case MariaDB, MySQL:
		return mariadb.UniqueConstraints(db.Connection, schemaName, tableName)
	case Oracle:
		return ora.UniqueConstraints(db.Connection, schemaName, tableName)
	case MSSQL:
		return mssql.UniqueConstraints(db.Connection, schemaName, tableName)
	}

	return d, unsupportedDBErr(db.id)
}
