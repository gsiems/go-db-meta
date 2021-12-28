package null

import (
	m "github.com/gsiems/go-db-meta/model"
)

// PrimaryKeys defines the query for obtaining the primary keys
// for the (tableSchema, tableName) parameters and returns the results
// of executing the query
func PrimaryKeys(db *m.DB, tableSchema, tableName string) ([]m.PrimaryKey, error) {

	q := ``
	return db.PrimaryKeys(q, tableSchema, tableName)
}
