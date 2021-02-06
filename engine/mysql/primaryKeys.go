package mysql

import (
	e "github.com/gsiems/go-db-meta/engine/mariadb"
	m "github.com/gsiems/go-db-meta/model"
)

// PrimaryKey currently wraps the mariadb.PrimaryKeys function
func PrimaryKeys(db *m.DB, tableSchema, tableName string) ([]m.PrimaryKey, error) {
	return e.PrimaryKeys(db, tableSchema, tableName)
}
