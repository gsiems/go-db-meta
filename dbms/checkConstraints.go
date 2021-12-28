package dbms

import (
	"github.com/gsiems/go-db-meta/dbms/mariadb"
	"github.com/gsiems/go-db-meta/dbms/mssql"
	"github.com/gsiems/go-db-meta/dbms/ora"
	"github.com/gsiems/go-db-meta/dbms/pg"
	"github.com/gsiems/go-db-meta/dbms/sqlite"
	m "github.com/gsiems/go-db-meta/model"
)

// CheckConstraints returns a slice of Check Constraints for the
// (schemaName, tableName) parameters
func (db *DBMS) CheckConstraints(schemaName, tableName string) ([]m.CheckConstraint, error) {

	var d []m.CheckConstraint

	switch db.id {
	case PostgreSQL:
		return pg.CheckConstraints(db.Connection, schemaName, tableName)
	case SQLite:
		return sqlite.CheckConstraints(db.Connection, schemaName, tableName)
	case MariaDB, MySQL:
		return mariadb.CheckConstraints(db.Connection, schemaName, tableName)
	case Oracle:
		return ora.CheckConstraints(db.Connection, schemaName, tableName)
	case MSSQL:
		return mssql.CheckConstraints(db.Connection, schemaName, tableName)
	}

	return d, unsupportedDBErr(db.id)
}
