package pg

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Dependencies defines the query for obtaining a list of dependencies
// for the (schemaName, objectName) parameters and returns the results
// of executing the query
func Dependencies(db *sql.DB, schemaName, objectName string) ([]m.Dependency, error) {

	q := `
WITH args AS (
    SELECT current_database () AS db_name,
            coalesce ( $1, '' ) AS schema_name,
            coalesce ( $2, '' ) AS object_name,
            coalesce ( $1, $2, '' ) = '' AS ignore_schema,
            coalesce ( $2, '' ) = '' AS ignore_object
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
),
dependents AS (
    SELECT ev_class,
            split_part ( regexp_split_to_table ( ev_action, ':relid ' ), ' ', 1 ) AS dependent_oid
        FROM pg_rewrite
),
dep_map AS (
    SELECT ev_class AS parent_oid,
            dependent_oid::oid AS child_oid
        FROM dependents
        WHERE dependent_oid NOT LIKE '%QUERY'
            AND ev_class <> dependent_oid::oid
)
SELECT DISTINCT args.db_name AS object_catalog,
        n.nspname AS object_schema,
        c.relname AS object_name,
        pg_catalog.pg_get_userbyid ( c.relowner ) AS object_owner,
        coalesce ( tt.label, 'other (' || c.relkind || ')' ) AS object_type,
        args.db_name AS dep_object_catalog,
        n2.nspname AS dep_object_schema,
        c2.relname AS dep_object_name,
        pg_catalog.pg_get_userbyid ( c2.relowner ) AS dep_object_owner,
        coalesce ( rt.label, 'other (' || c2.relkind || ')' ) AS dep_object_type
    FROM pg_catalog.pg_class c
    CROSS JOIN args
    INNER JOIN dep_map d
        ON ( c.oid = d.parent_oid )
    INNER JOIN pg_catalog.pg_class c2
        ON ( c2.oid = d.child_oid )
    INNER JOIN types AS tt
        ON ( tt.relkind = c.relkind )
    INNER JOIN types AS rt
        ON ( rt.relkind = c2.relkind )
    INNER JOIN pg_catalog.pg_namespace n
        ON ( n.oid = c.relnamespace )
    INNER JOIN pg_catalog.pg_namespace n2
        ON ( n2.oid = c2.relnamespace )
    WHERE n.nspname <> 'information_schema'
        AND n.nspname !~ '^pg_'
        AND n2.nspname <> 'information_schema'
        AND n2.nspname !~ '^pg_'
        AND ( ( ( n.nspname = args.schema_name OR ( args.ignore_schema ) )
                    AND ( c.relname = args.object_name OR args.ignore_object ) )
                OR ( ( n2.nspname = args.schema_name OR args.ignore_schema )
                    AND ( c2.relname = args.object_name OR args.ignore_object ) ) )
`
	return m.Dependencies(db, q, schemaName, objectName)
}
