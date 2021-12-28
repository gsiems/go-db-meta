package pg

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// UniqueConstraints defines the query for obtaining a list of unique
// constraints for the (schemaName, tableName) parameters and returns the
// results of executing the query
func UniqueConstraints(db *sql.DB, schemaName, tableName string) ([]m.UniqueConstraint, error) {

	q := `
WITH args AS (
    SELECT coalesce ( $1, '' ) AS schema_name,
            coalesce ( $2, '' ) AS table_name
)
SELECT current_database () AS table_catalog,
        nr.nspname AS table_schema,
        r.relname AS table_name,
        c.conname AS constraint_name,
        split_part ( split_part ( pg_get_constraintdef ( c.oid ), '(', 2 ), ')', 1 ) AS column_names,
        'Enabled' AS status,
        d.description AS comments
    FROM pg_class r
    CROSS JOIN args
    INNER JOIN pg_namespace nr
        ON ( nr.oid = r.relnamespace )
    INNER JOIN pg_constraint c
        ON ( c.conrelid = r.oid )
    INNER JOIN pg_namespace nc
        ON ( nc.oid = c.connamespace )
    LEFT OUTER JOIN pg_description d
        ON ( d.objoid = c.oid )
    WHERE r.relkind = 'r'
        AND c.contype = 'u'
        AND nr.nspname <> 'information_schema'
        AND nr.nspname !~ '^pg_'
        AND ( nr.nspname = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
        AND ( r.relname = args.table_name OR args.table_name = '' )
    ORDER BY r.relname,
        c.conname
`
	return m.UniqueConstraints(db, q, schemaName, tableName)
}
