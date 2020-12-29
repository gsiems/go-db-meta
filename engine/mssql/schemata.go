package mssql

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Schemata defines the query for obtaining a list of schemata
// as filtered by the (nclude, xclude) parameters and returns the
// results of executing the query
func Schemata(db *m.DB, nclude, xclude string) ([]m.Schema, error) {

	q := `
SELECT catalog_name,
        schema_name,
        schema_owner,
        default_character_set_catalog,
        default_character_set_schema,
        default_character_set_name,
        NULL AS comment
    FROM information_schema.schemata
    WHERE schema_name NOT IN ( 'INFORMATION_SCHEMA', 'sys' )
        AND substring ( schema_name, 1, 3 ) <> 'db_'
`
	return db.Schemata(q, nclude, xclude)
}
