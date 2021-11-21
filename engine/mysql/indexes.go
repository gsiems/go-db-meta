package mysql

import (
	e "github.com/gsiems/go-db-meta/engine/mariadb"
	m "github.com/gsiems/go-db-meta/model"
)

// Indexes defines the query for obtaining a list of indexes
// for the (tableSchema, tableName) parameters and returns the results
// of executing the query
func Indexes(db *m.DB, tableSchema, tableName string) ([]m.Index, error) {
	return e.Indexes(db, tableSchema, tableName)
}
