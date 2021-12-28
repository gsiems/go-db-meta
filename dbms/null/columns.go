package null

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Columns defines the query for obtaining a list of columns
// for the (tableSchema, tableName) parameters and returns the results
// of executing the query
func Columns(db *m.DB, tableSchema, tableName string) ([]m.Column, error) {

	q := ``
	return db.Columns(q, tableSchema, tableName)
}
