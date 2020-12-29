package mysql

import (
	e "github.com/gsiems/go-db-meta/engine/mariadb"
	m "github.com/gsiems/go-db-meta/model"
)

// Schemata currently wraps the mariadb.Schemata function
func Schemata(db *m.DB, nclude, xclude string) ([]m.Schema, error) {
	return e.Schemata(db, nclude, xclude)
}
