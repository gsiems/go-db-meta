package mariadb

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Schemata defines the query for obtaining a list of schemata
// as filtered by the (nclude, xclude) parameters and returns the
// results of executing the query
func Schemata(db *sql.DB, nclude, xclude string) ([]m.Schema, error) {

	q := `
SELECT catalog_name,
        schema_name,
        NULL AS schema_owner,
        NULL AS default_character_set_catalog,
        NULL AS default_character_set_schema,
        default_character_set_name,
        -- default_collation_name,
        -- sql_path
        NULL AS comment
    FROM information_schema.schemata
    WHERE schema_name NOT IN ( 'information_schema', 'mysql', 'performance_schema', 'sys' )
`

	return m.Schemata(db, q, nclude, xclude)
}
