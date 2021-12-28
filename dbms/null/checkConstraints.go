package null

import (
	m "github.com/gsiems/go-db-meta/model"
)

// CheckConstraints defines the query for obtaining the check
// constraints for the tables specified by the (tableSchema, tableName)
// parameters and returns the results of executing the query
func CheckConstraints(db *m.DB, tableSchema, tableName string) ([]m.CheckConstraint, error) {

	q := ``
	return db.CheckConstraints(q, tableSchema, tableName)
}
