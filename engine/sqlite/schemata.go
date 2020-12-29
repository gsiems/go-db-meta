package sqlite

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Schemata doesn't do much as sqlite doesn't appear to have schemas
func Schemata(db *m.DB, nclude, xclude string) ([]m.Schema, error) {
	var d []m.Schema
	return d, nil
}
