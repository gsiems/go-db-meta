package sqlite

import (
	m "github.com/gsiems/go-db-meta/model"
)

// CheckConstraints returns an empty set it is unclear how to extract
// check constraints from SQLite other than parsing the sqlite_master.sql
// column
func CheckConstraints(db *m.DB, tableSchema, tableName string) ([]m.CheckConstraint, error) {

	q := ``
	return db.CheckConstraints(q, tableSchema, tableName)
}
