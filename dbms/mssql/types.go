package mssql

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Types returns an empty set of types as MS-SQL does not appear to support user defined types
func Types(db *sql.DB, schemaName string) ([]m.Type, error) {
	// MS-SQL doesn't appear to support user defined types...
	q := ``
	return m.Types(db, q, schemaName)
}
