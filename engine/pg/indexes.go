package pg

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Indexes defines the query for obtaining a list of indexes
// for the (tableSchema, tableName) parameters and returns the results
// of executing the query
func Indexes(db *m.DB, tableSchema, tableName string) ([]m.Index, error) {

	q := `
WITH args AS (
    SELECT coalesce ( $1, '' ) AS schema_name,
            coalesce ( $2, '' ) AS table_name
)
SELECT current_database () AS index_catalog,
        nr.nspname AS index_schema,
        c2.relname AS index_name,
        '' AS index_type,
        split_part ( split_part ( pg_catalog.pg_get_indexdef ( i.indexrelid, 0, true ), '(', 2 ), ')', 1 ) AS index_columns,
        current_database () AS table_catalog,
        nr.nspname AS table_schema,
        c.relname AS table_name,
        CASE
            WHEN i.indisunique THEN 'YES'
            ELSE 'NO'
            END AS is_unique,
        d.description AS comments
    FROM pg_catalog.pg_class c
    CROSS JOIN args
    INNER JOIN pg_catalog.pg_index i
        ON ( c.oid = i.indrelid )
    INNER JOIN pg_catalog.pg_class c2
        ON ( i.indexrelid = c2.oid )
    LEFT OUTER JOIN pg_catalog.pg_description d
        ON ( d.objoid = i.indexrelid )
    INNER JOIN pg_namespace nr
        ON ( nr.oid = c.relnamespace )
    WHERE nr.nspname <> 'information_schema'
        AND nr.nspname !~ '^pg_'
        AND ( nr.nspname = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
        AND ( c.relname = args.table_name OR args.table_name = '' )
`
	return db.Indexes(q, tableSchema, tableName)
}
