package pg

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Tables defines the query for obtaining a list of tables and views
// for the (schemaName) parameter and returns the results of
// executing the query
func Tables(db *sql.DB, schemaName string) ([]m.Table, error) {

	q := `
WITH args AS (
    SELECT current_database () AS db_name,
            coalesce ( $1, '' ) AS schema_name,
            coalesce ( $1, '' ) = '' AS ignore_schema
)
SELECT args.db_name AS catalog_name,
        n.nspname AS table_schema,
        c.relname AS table_name,
        pg_catalog.pg_get_userbyid ( c.relowner ) AS table_owner,
        CASE
            WHEN i.inhrelid IS NOT NULL THEN 'TABLE PARTITION'
            WHEN c.relkind = 'f' THEN 'FOREIGN TABLE'
            WHEN c.relkind = 'm' THEN 'MATERIALIZED VIEW'
            WHEN c.relkind = 'p' THEN 'PARTITIONED TABLE'
            WHEN c.relkind = 'r' THEN 'TABLE'
            WHEN c.relkind = 'v' THEN 'VIEW'
            END AS table_type,
        c.reltuples::bigint AS row_count,
        pg_catalog.obj_description ( c.oid, 'pg_class' ) AS comment,
        CASE c.relkind
            WHEN 'v' THEN pg_catalog.pg_get_viewdef ( c.oid, true )
            WHEN 'm' THEN pg_catalog.pg_get_viewdef ( c.oid, true )
            END AS query
    FROM pg_catalog.pg_class c
    CROSS JOIN args
    LEFT OUTER JOIN pg_catalog.pg_namespace n
        ON ( n.oid = c.relnamespace )
    LEFT OUTER JOIN pg_catalog.pg_inherits i
        ON ( c.oid = i.inhrelid )
    WHERE c.relkind IN ( 'f', 'm', 'p', 'r', 'v' )
        -- AND c.relpersistence IN ( 'p' ) -- 'u' ??
        AND n.nspname <> 'information_schema'
        AND n.nspname !~ '^pg_'
        AND ( n.nspname = args.schema_name
            OR args.ignore_schema )
`
	return m.Tables(db, q, schemaName)
}
