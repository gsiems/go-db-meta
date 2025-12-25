package pg

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// PrimaryKeys defines the query for obtaining the primary keys
// for the (schemaName, tableName) parameters and returns the results
// of executing the query
func PrimaryKeys(db *sql.DB, schemaName, tableName string) ([]m.PrimaryKey, error) {

	q := `
WITH args AS (
    SELECT current_database () AS db_name,
            coalesce ( $1, '' ) AS schema_name,
            coalesce ( $2, '' ) AS table_name,
            coalesce ( $1, $2, '' ) = '' AS ignore_schema,
            coalesce ( $2, '' ) = '' AS ignore_table
)
SELECT args.db_name AS table_catalog,
        nr.nspname AS table_schema,
        r.relname AS table_name,
        c.conname AS constraint_name,
        split_part ( split_part ( pg_get_constraintdef ( c.oid ), '(', 2 ), ')', 1 ) AS constraint_columns,
        'Enabled' AS status,
        d.description AS comments
    FROM pg_class r
    CROSS JOIN args
    INNER JOIN pg_namespace nr
        ON ( nr.oid = r.relnamespace )
    INNER JOIN pg_constraint c
        ON ( c.conrelid = r.oid )
--    INNER JOIN pg_namespace nc
--        ON ( nc.oid = c.connamespace )
    LEFT OUTER JOIN pg_description d
        ON ( d.objoid = c.oid )
    WHERE r.relkind = 'r'
        AND c.contype = 'p'
        AND c.contype <> 'f'
        AND nr.nspname <> 'information_schema'
        AND nr.nspname !~ '^pg_'
        AND ( nr.nspname = args.schema_name OR args.ignore_schema )
        AND ( r.relname = args.table_name OR args.ignore_table )
    ORDER BY nr.nspname,
        r.relname
`
	return m.PrimaryKeys(db, q, schemaName, tableName)
}
