package dbms

import (
	"github.com/gsiems/go-db-meta/dbms/mssql"
	"github.com/gsiems/go-db-meta/dbms/ora"
	"github.com/gsiems/go-db-meta/dbms/pg"
	m "github.com/gsiems/go-db-meta/model"
)

// Dependencies returns a slice of Dependecies for the
// (schemaName, objectName) parameters
func (db *DBMS) Dependencies(schemaName, objectName string) ([]m.Dependency, error) {

	var d []m.Dependency

	switch db.id {
	case PostgreSQL:
		return pg.Dependencies(db.Connection, schemaName, objectName)
	case SQLite, MariaDB, MySQL:
		return d, nil
	case Oracle:
		return ora.Dependencies(db.Connection, schemaName, objectName)
	case MSSQL:
		return mssql.Dependencies(db.Connection, schemaName, objectName)
	}

	return d, unsupportedDBErr(db.id)
}
