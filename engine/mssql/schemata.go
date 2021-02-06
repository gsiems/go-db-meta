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
        convert ( varchar ( 8000 ), xp.value ) AS comments
    FROM information_schema.schemata
    OUTER APPLY ::fn_listextendedproperty ( 'MS_Description', 'schema', schema_name ) xp
    WHERE schema_name NOT IN ( 'INFORMATION_SCHEMA', 'sys' )
        AND substring ( schema_name, 1, 3 ) <> 'db_'
`
	return db.Schemata(q, nclude, xclude)
}
