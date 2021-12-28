package mssql

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Dependencies defines the query for obtaining a list of dependencies
// for the (schemaName, objectName) parameters and returns the results
// of executing the query
func Dependencies(db *sql.DB, schemaName, objectName string) ([]m.Dependency, error) {

	q := ``
	return m.Dependencies(db, q, schemaName, objectName)
}
