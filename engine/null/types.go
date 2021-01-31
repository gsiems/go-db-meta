package null

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Types defines the query for obtaining a list of user defined types
// for the (schema) parameters and returns the results of executing the
// query
func Types(db *m.DB, schema string) ([]m.Type, error) {

	q := ``
	return db.Types(q, schema)
}
