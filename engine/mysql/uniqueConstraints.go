package mysql

import (
	e "github.com/gsiems/go-db-meta/engine/mariadb"
	m "github.com/gsiems/go-db-meta/model"
)

// UniqueConstraints defines the query for obtaining a list of unique
// constraints for the (tableSchema, tableName) parameters and returns the
// results of executing the query
func UniqueConstraints(db *m.DB, tableSchema, tableName string) ([]m.UniqueConstraint, error) {
	return e.UniqueConstraints(tableSchema, tableName)
}
