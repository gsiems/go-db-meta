package sqlite

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Types returns an empty set of types as SQLite does not support user defined types
func Types(db *m.DB, schema string) ([]m.Type, error) {

	q := ``
	return db.Types(q, schema)
}
