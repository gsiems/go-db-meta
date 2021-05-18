package pg

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Types defines the query for obtaining a list of user defined types
// for the (schema) parameters and returns the results of executing the
// query
func Types(db *m.DB, schema string) ([]m.Type, error) {

	q := `
WITH args AS (
    SELECT $1 AS schema_name
)
SELECT current_database () AS user_defined_type_catalog,
        n.nspname AS user_defined_type_schema,
        t.typname AS user_defined_type_name,
        pg_catalog.pg_get_userbyid ( t.typowner ) AS user_defined_type_owner,
        pg_catalog.obj_description ( t.oid, 'pg_type' ) AS comment
    FROM pg_catalog.pg_type t
    CROSS JOIN args
    JOIN pg_catalog.pg_namespace n
        ON n.oid = t.typnamespace
    LEFT JOIN pg_catalog.pg_type bt
        ON ( bt.oid = t.typbasetype )
    WHERE ( t.typrelid = 0
            OR ( SELECT c.relkind = 'c'
                    FROM pg_catalog.pg_class c
                    WHERE c.oid = t.typrelid ) )
        AND NOT EXISTS (
                SELECT 1
                    FROM pg_catalog.pg_type el
                    WHERE el.oid = t.typelem
                        AND el.typarray = t.oid )
        AND n.nspname <> 'information_schema'
        AND n.nspname !~ '^pg_'
        AND t.typtype NOT IN ( 'p' )
        AND NOT ( t.typtype = 'c'
            AND n.nspname = 'pg_catalog' )
        AND ( n.nspname = args.schema_name
            OR args.schema_name = '' )
`
	return db.Types(q, schema)
}

/*
Columns for composite types

WITH args AS (
    SELECT $1 AS schema_name,
            $2 AS table_name
)
SELECT current_database() AS type_catalog,
        n.nspname AS type_schema,
        t.typname AS type_name,
        a.attname AS column_name,
        a.attnum AS ordinal_position,
        pg_catalog.format_type ( a.atttypid, a.atttypmod ) AS data_type,
        CASE
            WHEN a.attnotnull THEN 'NO'
            ELSE 'YES'
            END AS is_nullable--,
        pg_catalog.pg_get_expr ( ad.adbin, ad.adrelid )  AS column_default
    FROM pg_catalog.pg_type t
    CROSS JOIN args
    JOIN pg_catalog.pg_namespace n
        ON n.oid = t.typnamespace
    JOIN pg_catalog.pg_attribute a
        ON ( a.attrelid = t.typrelid )
    LEFT OUTER JOIN pg_catalog.pg_attrdef ad
        ON ( a.attrelid = ad.adrelid
            AND a.attnum = ad.adnum )
    WHERE ( t.typrelid = 0
            OR ( SELECT c.relkind = 'c'
                    FROM pg_catalog.pg_class c
                    WHERE c.oid = t.typrelid ) )
        AND NOT EXISTS (
                SELECT 1
                    FROM pg_catalog.pg_type el
                    WHERE el.oid = t.typelem
                        AND el.typarray = t.oid )
        AND n.nspname <> 'information_schema'
        AND n.nspname !~ '^pg_'
        AND t.typtype NOT IN ( 'p' )
        AND NOT ( t.typtype = 'c'
            AND n.nspname = 'pg_catalog' )
        AND a.attnum > 0
        AND NOT a.attisdropped
        AND ( n.nspname = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
        AND ( t.typname = args.table_name OR args.table_name = '' )
    ORDER BY n.nspname,
        t.typname,
        a.attnum
*/
