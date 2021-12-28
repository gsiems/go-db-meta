package pg

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Domains defines the query for obtaining a list of domains
// for the (schemaName) parameter and returns the results
// of executing the query
func Domains(db *sql.DB, schemaName string) ([]m.Domain, error) {

	q := `
WITH args AS (
    SELECT coalesce ( $1, '' ) AS schema_name
)
SELECT current_database() AS domain_catalog,
        n.nspname AS domain_schema,
        t.typname AS domain_name,
        pg_catalog.pg_get_userbyid ( t.typowner ) AS domain_owner,
        pg_catalog.format_type ( t.typbasetype, t.typtypmod ) AS domain_type,
        t.typdefault AS domain_default,
        pg_catalog.obj_description ( t.oid, 'pg_type' ) AS comments
    FROM pg_catalog.pg_type t
    CROSS JOIN args
    LEFT JOIN pg_catalog.pg_namespace n
        ON n.oid = t.typnamespace
    WHERE t.typtype = 'd'
        AND pg_catalog.pg_type_is_visible ( t.oid )
        AND n.nspname <> 'information_schema'
        AND n.nspname !~ '^pg_'
        AND ( n.nspname = args.schema_name
            OR args.schema_name = '' )
`

	return m.Domains(db, q, schemaName)
}
