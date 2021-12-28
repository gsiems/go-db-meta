package pg

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Columns defines the query for obtaining a list of columns
// for the (schemaName, tableName) parameters and returns the results
// of executing the query
func Columns(db *sql.DB, schemaName, tableName string) ([]m.Column, error) {

	q := `
WITH args AS (
    SELECT coalesce ( $1, '' ) AS schema_name,
            coalesce ( $2, '' ) AS table_name
)
SELECT current_database () AS table_catalog,
        n.nspname AS table_schema,
        c.relname AS table_name,
        a.attname AS column_name,
        a.attnum AS ordinal_position,
        pg_catalog.format_type ( a.atttypid, a.atttypmod ) AS data_type,
        CASE
            WHEN a.attnotnull THEN 'NO'
            ELSE 'YES'
            END AS is_nullable,
        pg_catalog.pg_get_expr ( ad.adbin, ad.adrelid )  AS column_default,
        CASE
            WHEN t.typtype = 'd' THEN current_database()
        END AS domain_catalog,
        CASE
            WHEN t.typtype = 'd' THEN nt.nspname
        END AS domain_schema,
         CASE
            WHEN t.typtype = 'd' THEN t.typname
        END AS domain_name,
        -- udt_catalog
        -- udt_schema
        -- udt_name
        pg_catalog.col_description ( a.attrelid, a.attnum ) AS comments
    FROM pg_catalog.pg_class c
    CROSS JOIN args
    LEFT OUTER JOIN pg_catalog.pg_namespace n
        ON ( n.oid = c.relnamespace )
    LEFT OUTER JOIN pg_catalog.pg_attribute a
        ON ( c.oid = a.attrelid
            AND a.attnum > 0
            AND NOT a.attisdropped )
    LEFT OUTER JOIN pg_catalog.pg_attrdef ad
        ON ( a.attrelid = ad.adrelid
            AND a.attnum = ad.adnum )
    JOIN pg_type t
        ON ( a.atttypid = t.oid )
    JOIN pg_namespace nt
        ON ( t.typnamespace = nt.oid )
    WHERE c.relkind IN ( 'v', 'r', 'f', 'p', 'm' )
        AND a.attnum > 0
        AND NOT a.attisdropped
        AND n.nspname <> 'information_schema'
        AND n.nspname !~ '^pg_'
        AND ( n.nspname = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
        AND ( c.relname = args.table_name OR args.table_name = '' )
    ORDER BY n.nspname,
        c.relname,
        a.attnum
`

	return m.Columns(db, q, schemaName, tableName)
}
