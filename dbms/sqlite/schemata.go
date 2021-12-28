package sqlite

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Schemata doesn't do much as sqlite doesn't appear to have schemas
func Schemata(db *sql.DB, nclude, xclude string) ([]m.Schema, error) {

	var u m.Schema
	var r []m.Schema

	catName, err := catalogName(db)
	if err != nil {
		return r, err
	}
	u.CatalogName = catName

	charSetName, err := defaultCharacterSetName(db)
	if err != nil {
		return r, err
	}
	u.DefaultCharacterSetName = charSetName

	r = append(r, u)

	return r, err
}
