package mssql

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Types returns an empty set of types as MS-SQL does not appear to support user defined types
func Types(db *m.DB, schema string) ([]m.Type, error) {
	var d []m.Type
	return d, nil
}
