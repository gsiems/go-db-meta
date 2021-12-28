package null

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Schemata defines the query for obtaining a list of schemata
// as filtered by the (nclude, xclude) parameters and returns the
// results of executing the query
func Schemata(db *m.DB, nclude, xclude string) ([]m.Schema, error) {

	q := ``
	return db.Schemata(q, nclude, xclude)
}
