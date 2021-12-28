package null

import (
	m "github.com/gsiems/go-db-meta/model"
)

// ReferentialConstraints defines the query for obtaining the
// referential constraints for the (tableSchema, tableName) parameters
// (as either the parent or child) and returns the results of executing
// the query
func ReferentialConstraints(db *m.DB, tableSchema, tableName string) ([]m.ReferentialConstraint, error) {

	q := ``
	return db.ReferentialConstraints(q, tableSchema, tableName)
}
