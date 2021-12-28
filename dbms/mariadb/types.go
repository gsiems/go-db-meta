package mariadb

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Types returns an empty set of types as MariaDB does not appear to support user defined types
func Types(db *m.DB, schema string) ([]m.Type, error) {

	q := ``
	return db.Types(q, schema)
}
