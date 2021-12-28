package null

import (
	m "github.com/gsiems/go-db-meta/model"
)

// UniqueConstraints defines the query for obtaining a list of unique
// constraints for the (tableSchema, tableName) parameters and returns the
// results of executing the query
func UniqueConstraints(db *m.DB, tableSchema, tableName string) ([]m.UniqueConstraint, error) {

	q := ``
	return db.UniqueConstraints(q, tableSchema, tableName)
}
