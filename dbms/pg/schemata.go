package pg

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Schemata defines the query for obtaining a list of schemata
// as filtered by the (nclude, xclude) parameters and returns the
// results of executing the query
func Schemata(db *sql.DB, nclude, xclude string) ([]m.Schema, error) {

	q := `
SELECT pg_catalog.current_database ()::text AS catalog_name,
        n.nspname::text AS schema_name,
        pg_catalog.pg_get_userbyid ( n.nspowner )::text AS schema_owner,
        NULL::text AS character_set_catalog,
        NULL::text AS character_set_schema,
        NULL::text AS character_set_name,
        pg_catalog.obj_description ( n.oid, 'pg_namespace' ) AS comments
    FROM pg_catalog.pg_namespace n
    WHERE n.nspname <> 'information_schema'
        AND n.nspname !~ '^pg_'
    ORDER BY n.nspname
`
	return m.Schemata(db, q, nclude, xclude)
}
