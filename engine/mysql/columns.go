package mysql

import (
	e "github.com/gsiems/go-db-meta/engine/mariadb"
	m "github.com/gsiems/go-db-meta/model"
)

// Columns currently wraps the mariadb.Columns function
func Columns(db *m.DB, tableSchema, tableName string) ([]m.Column, error) {
	return e.Schemata(db, tableSchema, tableName)
}
