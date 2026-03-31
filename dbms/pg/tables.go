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
    SELECT pg_catalog.current_database () AS db_name,
            coalesce ( $1, '' ) AS schema_name,
            coalesce ( $1, '' ) = '' AS ignore_schema
),
types AS (
    SELECT *
        FROM (
            VALUES
                ( 'f', 'FOREIGN TABLE' ),
                ( 'm', 'MATERIALIZED VIEW' ),
                ( 'p', 'PARTITIONED TABLE' ),
                ( 'r', 'TABLE' ),
                ( 't', 'TABLE' ),
                ( 'v', 'VIEW' )
            ) AS t ( relkind, label )
)
SELECT args.db_name::text AS catalog_name,
        n.nspname::text AS table_schema,
        c.relname::text AS table_name,
        pg_catalog.pg_get_userbyid ( c.relowner )::text AS table_owner,
        CASE WHEN i.inhrelid IS NOT NULL THEN 'TABLE PARTITION' ELSE types.label END AS table_type,
        c.reltuples::bigint AS row_count,
        pg_catalog.obj_description ( c.oid, 'pg_class' ) AS comment,
        CASE WHEN c.relkind IN ( 'm', 'v' ) THEN pg_catalog.pg_get_viewdef ( c.oid, true ) END AS query
    FROM pg_catalog.pg_class c
    CROSS JOIN args
    INNER JOIN types
        ON ( types.relkind = c.relkind::text )
    LEFT OUTER JOIN pg_catalog.pg_namespace n
        ON ( n.oid = c.relnamespace )
    LEFT OUTER JOIN pg_catalog.pg_inherits i
        ON ( c.oid = i.inhrelid )
    WHERE n.nspname <> 'information_schema'
        AND n.nspname !~ '^pg_'
        AND ( n.nspname = args.schema_name
            OR args.ignore_schema )
        -- AND c.relpersistence IN ( 'p' ) -- 'u' ??
`
	return m.Tables(db, q, schemaName)
}
