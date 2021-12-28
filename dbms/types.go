package dbms

import (
	"github.com/gsiems/go-db-meta/dbms/ora"
	"github.com/gsiems/go-db-meta/dbms/pg"
	m "github.com/gsiems/go-db-meta/model"
)

// Types returns a slice of Types for the (schema) parameter
func (db *DBMS) Types(schemaName string) ([]m.Type, error) {

	var d []m.Type

	switch db.id {
	case PostgreSQL:
		return pg.Types(db.Connection, schemaName)
	case SQLite, MariaDB, MySQL, MSSQL:
		return d, nil
	case Oracle:
		return ora.Types(db.Connection, schemaName)
	}

	return d, unsupportedDBErr(db.id)
}
