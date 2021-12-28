package mssql

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// UniqueConstraints defines the query for obtaining a list of unique
// constraints for the (schemaName, tableName) parameters and returns the
// results of executing the query
func UniqueConstraints(db *sql.DB, schemaName, tableName string) ([]m.UniqueConstraint, error) {

	q := ``
	return m.UniqueConstraints(db, q, schemaName, tableName)
}
