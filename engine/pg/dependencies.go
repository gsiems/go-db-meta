package pg

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Dependencies defines the query for obtaining a list of dependencies
// for the (objectSchema, objectName) parameters and returns the results
// of executing the query
func Dependencies(db *m.DB, objectSchema, objectName string) ([]m.Dependency, error) {

	q := `
WITH args AS (
    SELECT coalesce ( $1, '' ) AS schema_name,
            coalesce ( $2, '' ) AS object_name
),
types AS (
    SELECT 'f' AS obj_kind,
            'FOREIGN TABLE' AS obj_type
    UNION
    SELECT 'm' AS obj_kind,
            'MATERIALIZED VIEW' AS obj_type
    UNION
    SELECT 'r' AS obj_kind,
            'TABLE' AS obj_type
    UNION
    SELECT 't' AS obj_kind,
            'TABLE' AS obj_type
    UNION
    SELECT 'v' AS obj_kind,
            'VIEW' AS obj_type
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
SELECT DISTINCT pg_catalog.current_database () AS object_catalog,
        n.nspname AS object_schema,
        c.relname AS object_name,
        pg_catalog.pg_get_userbyid ( c.relowner ) AS object_owner,
        coalesce ( tt.obj_type, 'other (' || c.relkind || ')' ) AS object_type,
        pg_catalog.current_database () AS dep_object_catalog,
        n2.nspname AS dep_object_schema,
        c2.relname AS dep_object_name,
        pg_catalog.pg_get_userbyid ( c2.relowner ) AS dep_object_owner,
        coalesce ( rt.obj_type, 'other (' || c2.relkind || ')' ) AS dep_object_type
    FROM pg_catalog.pg_class c
    CROSS JOIN args
    INNER JOIN dep_map d
        ON ( c.oid = d.parent_oid )
    INNER JOIN pg_catalog.pg_class c2
        ON ( c2.oid = d.child_oid )
    INNER JOIN types AS tt
        ON ( tt.obj_kind = c.relkind )
    INNER JOIN types AS rt
        ON ( rt.obj_kind = c2.relkind )
    INNER JOIN pg_catalog.pg_namespace n
        ON ( n.oid = c.relnamespace )
    INNER JOIN pg_catalog.pg_namespace n2
        ON ( n2.oid = c2.relnamespace )
    WHERE  n.nspname <> 'information_schema'
        AND n.nspname !~ '^pg_'
        AND n2.nspname <> 'information_schema'
        AND n2.nspname !~ '^pg_'
        AND ( ( ( n.nspname = args.schema_name OR ( args.schema_name = '' AND args.object_name = '' ) )
                    AND ( c.relname = args.object_name OR args.object_name = '' ) )
                OR ( ( n2.nspname = args.schema_name OR ( args.schema_name = '' AND args.object_name = '' ) )
                    AND ( c2.relname = args.object_name OR args.object_name = '' ) ) )
`
	return db.Dependencies(q, objectSchema, objectName)
}
