package mssql

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Indexes defines the query for obtaining a list of indexes
// for the (schemaName, tableName) parameters and returns the results
// of executing the query
func Indexes(db *sql.DB, schemaName, tableName string) ([]m.Index, error) {

	q := ``
	return m.Indexes(db, q, schemaName, tableName)
}
