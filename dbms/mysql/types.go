package mysql

import (
	e "github.com/gsiems/go-db-meta/engine/mariadb"
	m "github.com/gsiems/go-db-meta/model"
)

// Types currently wraps the mariadb.Types function
func Types(db *m.DB, schema string) ([]m.Type, error) {
	return e.Types(db, schema)
}
