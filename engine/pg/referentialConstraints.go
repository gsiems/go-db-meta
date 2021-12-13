package pg

import (
	m "github.com/gsiems/go-db-meta/model"
)

// ReferentialConstraints defines the query for obtaining the
// referential constraints for the (tableSchema, tableName) parameters
// (as either the parent or child) and returns the results of executing
// the query
func ReferentialConstraints(db *m.DB, tableSchema, tableName string) ([]m.ReferentialConstraint, error) {

	q := `
WITH args AS (
    SELECT coalesce ( $1, '' ) AS schema_name,
            coalesce ( $2, '' ) AS table_name
),
con_rules AS (
    SELECT 'f' AS con_type,
            'FULL' AS con_text
    UNION
    SELECT 'p' AS con_type,
            'PARTIAL' AS con_text
    UNION
    SELECT 's' AS con_type,
            'NONE' AS con_text
    UNION
    SELECT 'c' AS con_type,
            'CASCADE' AS con_text
    UNION
    SELECT 'n' AS con_type,
            'SET NULL' AS con_text
    UNION
    SELECT 'd' AS con_type,
            'SET DEFAULT' AS con_text
    UNION
    SELECT 'r' AS con_type,
            'RESTRICT' AS con_text
    UNION
    SELECT 'a' AS con_type,
            'NO ACTION' AS con_text
),
referential_constraints AS (
    SELECT current_database () AS constraint_catalog,
            ncon.nspname AS constraint_schema,
            con.conname AS constraint_name,
            con.oid,
            CASE
                WHEN npkc.nspname IS NOT NULL THEN current_database ()
                END AS unique_constraint_catalog,
            npkc.nspname AS unique_constraint_schema,
            pkc.conname AS unique_constraint_name,
            mr.con_text AS match_option,
            ur.con_text AS update_rule,
            dr.con_text AS delete_rule
        FROM pg_constraint con
        JOIN pg_namespace ncon
            ON ( ncon.oid = con.connamespace )
        JOIN pg_class c
            ON ( con.conrelid = c.oid
                AND con.contype = 'f' )
        LEFT JOIN con_rules mr
            ON ( mr.con_type = con.confmatchtype )
        LEFT JOIN con_rules ur
            ON ( ur.con_type = con.confupdtype )
        LEFT JOIN con_rules dr
            ON ( dr.con_type = con.confdeltype )
        LEFT JOIN pg_depend d1
            ON ( d1.objid = con.oid
                AND d1.classid = ( 'pg_constraint'::regclass )::oid
                AND d1.refclassid = ( 'pg_class'::regclass )::oid
                AND d1.refobjsubid = 0 )
        LEFT JOIN pg_depend d2
            ON ( d2.refclassid = ( 'pg_constraint'::regclass )::oid
                AND d2.classid = ( 'pg_class'::regclass )::oid
                AND d2.objid = d1.refobjid
                AND d2.objsubid = 0
                AND d2.deptype = 'i' )
        LEFT JOIN pg_constraint pkc
            ON ( pkc.oid = d2.refobjid
                AND pkc.contype IN ( 'p', 'u' )
                AND pkc.conrelid = con.confrelid )
        LEFT JOIN pg_namespace npkc
            ON ( pkc.connamespace = npkc.oid
                AND npkc.nspname <> 'information_schema'
                AND npkc.nspname !~ '^pg_' )
        WHERE ncon.nspname <> 'information_schema'
            AND ncon.nspname !~ '^pg_'
),
table_constraints AS (
    SELECT current_database () AS constraint_catalog,
        nc.nspname AS constraint_schema,
        c.conname AS constraint_name,
        current_database () AS table_catalog,
        nr.nspname AS table_schema,
        r.relname AS table_name,
        CASE c.contype
            WHEN 'f' THEN 'FOREIGN KEY'
            WHEN 'p' THEN 'PRIMARY KEY'
            WHEN 'u' THEN 'UNIQUE'
            END AS constraint_type,
        CASE
            WHEN c.condeferrable THEN 'YES'
            ELSE 'NO'
            END AS is_deferrable,
        CASE
            WHEN c.condeferred THEN 'YES'
            ELSE 'NO'
            END AS initially_deferred
    FROM pg_constraint c
    JOIN pg_namespace nc
        ON ( nc.oid = c.connamespace )
    JOIN pg_class r
        ON ( c.conrelid = r.oid )
    JOIN pg_namespace nr
        ON ( nr.oid = r.relnamespace )
    WHERE c.contype IN ( 'f', 'p', 'u' )
        AND r.relkind IN ( 'r', 'p' )
        AND NOT pg_is_other_temp_schema ( nr.oid )
        AND nc.nspname <> 'information_schema'
        AND nc.nspname !~ '^pg_'
        AND nr.nspname <> 'information_schema'
        AND nr.nspname !~ '^pg_'
)
SELECT --con.oid,
        --con.constraint_catalog,
        --con.constraint_schema,
        --con.constraint_name,
        tab.table_catalog,
        tab.table_schema,
        tab.table_name,
        split_part ( split_part ( pg_catalog.pg_get_constraintdef ( con.oid, true ), '(', 3 ), ')', 1 ) AS column_names,
        --tab.constraint_type,
        tab.constraint_name,
        --'YES' AS enforced,
        --con.unique_constraint_catalog,
        --con.unique_constraint_schema,
        --con.unique_constraint_name,
        rtab.table_catalog AS ref_table_catalog,
        rtab.table_schema AS ref_table_schema,
        rtab.table_name AS ref_table,
        split_part ( split_part ( pg_catalog.pg_get_constraintdef ( con.oid, true ), '(', 2 ), ')', 1 ) AS ref_column_names,
        --rtab.constraint_type AS ref_constraint_type,
        rtab.constraint_name AS ref_constraint_name,
        --con.unique_constraint_name,
        --con.is_deferrable,
        --con.initially_deferred,
        con.match_option,
        con.update_rule,
        con.delete_rule,
        'YES' AS is_enforced,
        d.description AS comments
    FROM referential_constraints con
    CROSS JOIN args
    INNER JOIN table_constraints tab
        ON ( con.constraint_catalog = tab.constraint_catalog
            AND con.constraint_schema = tab.constraint_schema
            AND con.constraint_name = tab.constraint_name )
    INNER JOIN table_constraints rtab
        ON ( rtab.constraint_catalog = con.unique_constraint_catalog
            AND rtab.constraint_schema = con.unique_constraint_schema
            AND rtab.constraint_name = con.unique_constraint_name )
    LEFT OUTER JOIN pg_catalog.pg_description d
        ON ( d.objoid = con.oid )
    WHERE ( ( tab.table_schema = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
            AND ( tab.table_name = args.table_name OR args.table_name = '' ) )
        OR ( ( rtab.table_schema = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
            AND ( rtab.table_name = args.table_name OR args.table_name = '' ) )
`
	return db.ReferentialConstraints(q, tableSchema, tableName)
}
