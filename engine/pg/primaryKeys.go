package pg

import (
	m "github.com/gsiems/go-db-meta/model"
)

// PrimaryKeys defines the query for obtaining the primary keys
// for the (tableSchema, tableName) parameters and returns the results
// of executing the query
func PrimaryKeys(db *m.DB, tableSchema, tableName string) ([]m.PrimaryKey, error) {

	q := `
WITH args AS (
    SELECT $1 AS schema_name,
            $2 AS table_name
)
SELECT current_database () AS table_catalog,
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
    INNER JOIN pg_namespace nc
        ON ( nc.oid = c.connamespace )
    LEFT OUTER JOIN pg_description d
        ON ( d.objoid = c.oid )
    WHERE r.relkind = 'r'
        AND c.contype = 'p'
        AND c.contype <> 'f'
        AND n.nspname <> 'information_schema'
        AND n.nspname !~ '^pg_'
        AND ( nr.nspname = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
        AND ( c.relname = args.table_name OR args.table_name = '' )
    ORDER BY nr.nspname,
        r.relname
`
	return db.PrimaryKeys(q, tableSchema, tableName)
}
