package mysql

import (
	e "github.com/gsiems/go-db-meta/engine/mariadb"
	m "github.com/gsiems/go-db-meta/model"
)

// ReferentialConstraints currently wraps the mariadb.ReferentialConstraints function
func ReferentialConstraints(db *m.DB, tableSchema, tableName string) ([]m.ReferentialConstraint, error) {
	return e.ReferentialConstraints(db, tableSchema, tableName)
}
