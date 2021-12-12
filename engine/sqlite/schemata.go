package sqlite

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Schemata doesn't do much as sqlite doesn't appear to have schemas
func Schemata(db *m.DB, nclude, xclude string) (d []m.Schema, err error) {

	var u m.Schema

	catName, err := catalogName(db)
	if err != nil {
		return d, err
	}
	u.CatalogName = catName

	charSetName, err := defaultCharacterSetName(db)
	if err != nil {
		return d, err
	}
	u.DefaultCharacterSetName = charSetName

	u.SchemaName = "default"

	d = append(d, u)

	return d, err
}
