package pg

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Indexes defines the query for obtaining a list of indexes
// for the (schemaName, tableName) parameters and returns the results
// of executing the query
func Indexes(db *sql.DB, schemaName, tableName string) ([]m.Index, error) {

	q := `
WITH args AS (
    SELECT current_database () AS db_name,
            coalesce ( $1, '' ) AS schema_name,
            coalesce ( $2, '' ) AS table_name,
            coalesce ( $1, $2, '' ) = '' AS ignore_schema,
            coalesce ( $2, '' ) = '' AS ignore_table
),
idx AS (
    SELECT nr.nspname::text AS index_schema,
            c2.relname::text AS index_name,
            nr.nspname::text AS table_schema,
            c.relname::text AS table_name,
            i.indisunique,
            d.description AS comments,
            regexp_split_to_array (
                split_part ( pg_catalog.pg_get_indexdef ( i.indexrelid, 0, true ), 'INDEX', 2 ),
                '[\(\)]' ) AS def
        FROM pg_catalog.pg_index i
        CROSS JOIN args
        INNER JOIN pg_catalog.pg_class c
            ON ( c.oid = i.indrelid )
        INNER JOIN pg_catalog.pg_class c2
            ON ( c2.oid = i.indexrelid )
        LEFT OUTER JOIN pg_catalog.pg_description d
            ON ( d.objoid = i.indexrelid )
        INNER JOIN pg_namespace nr
            ON ( nr.oid = c.relnamespace )
        WHERE i.indislive
            AND nr.nspname <> 'information_schema'
            AND nr.nspname !~ '^pg_'
            AND ( nr.nspname = args.schema_name OR args.ignore_schema )
            AND ( c.relname = args.table_name OR args.ignore_table )
)
SELECT args.db_name AS index_catalog,
        idx.index_schema,
        idx.index_name,
        split_part ( idx.def[1], ' ', 6 ) AS index_type,
        idx.def[2] AS index_columns,
        args.db_name AS table_catalog,
        idx.table_schema,
        idx.table_name,
        CASE
           WHEN idx.indisunique THEN 'YES'
           ELSE 'NO'
           END AS is_unique,
        comments
    FROM idx
    CROSS JOIN args
`
	return m.Indexes(db, q, schemaName, tableName)
}
