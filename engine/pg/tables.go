package pg

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Tables defines the query for obtaining a list of tables and views
// for the (schema) parameter and returns the results of
// executing the query
func Tables(db *m.DB, schema string) ([]m.Table, error) {

	q := `
WITH args AS (
    SELECT $1 AS schema_name
)
SELECT pg_catalog.current_database () AS catalog_name,
        n.nspname AS table_schema,
        c.relname AS table_name,
        pg_catalog.pg_get_userbyid ( c.relowner ) AS table_owner,
        CASE c.relkind
            WHEN 'f' THEN 'FOREIGN TABLE'
            WHEN 'm' THEN 'MATERIALIZED VIEW'
            WHEN 'r' THEN 'TABLE'
            WHEN 'v' THEN 'VIEW'
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
    WHERE c.relkind IN ( 'v', 'r', 'f', 'm' )
        -- AND c.relpersistence IN ( 'p' ) -- 'u' ??
        AND n.nspname <> 'information_schema'
        AND n.nspname !~ '^pg_'
        AND ( n.nspname = args.schema_name
            OR args.schema_name = '' )
`
	return db.Tables(q, schema)
}
