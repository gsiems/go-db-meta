package pg

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// ReferentialConstraints defines the query for obtaining the
// referential constraints for the (schemaName, tableName) parameters
// (as either the parent or child) and returns the results of executing
// the query
func ReferentialConstraints(db *sql.DB, schemaName, tableName string) ([]m.ReferentialConstraint, error) {

	q := `
WITH args AS (
    SELECT current_database () AS db_name,
            coalesce ( $1, '' ) AS schema_name,
            coalesce ( $2, '' ) AS table_name,
            coalesce ( $1, $2, '' ) = '' AS ignore_schema,
            coalesce ( $2, '' ) = '' AS ignore_table
),
conftypes as (
    SELECT *
        FROM (
            VALUES
                ( 'a', 'NO ACTION' ),
                ( 'c', 'CASCADE' ),
                ( 'd', 'SET DEFAULT' ),
                ( 'f', 'FULL' ),
                ( 'n', 'SET NULL' ),
                ( 'p', 'PARTIAL' ),
                ( 'r', 'RESTRICT' ),
                ( 's', 'NONE' )
            ) AS t ( conftype, label )
),
referential_constraints AS (
    SELECT con.oid,
            tsc.nspname::text AS table_schema,
            tbl.relname::text AS table_name,
            con.conname::text AS constraint_name,
            rcon.relname::text AS ref_constraint_name,
            mr.label AS match_option,
            ur.label AS update_rule,
            dr.label AS delete_rule,
            regexp_split_to_array ( pg_catalog.pg_get_constraintdef ( con.oid ), '[\(\)]' ) AS def
        FROM pg_catalog.pg_constraint con
        JOIN pg_catalog.pg_class tbl
            ON ( tbl.oid = con.conrelid )
        JOIN pg_catalog.pg_namespace tsc
            ON ( tsc.oid = tbl.relnamespace )
        JOIN pg_catalog.pg_class rcon
            ON ( rcon.oid = con.conindid )
        LEFT JOIN conftypes mr
            ON ( mr.conftype = con.confmatchtype::text )
        LEFT JOIN conftypes ur
            ON ( ur.conftype = con.confupdtype::text )
        LEFT JOIN conftypes dr
            ON ( dr.conftype = con.confdeltype::text )
        WHERE con.contype = 'f'
            AND tsc.nspname <> 'information_schema'
            AND tsc.nspname !~ '^pg_'
),
base AS (
    SELECT con.oid,
            con.table_schema,
            con.table_name,
            con.def[2] AS column_names,
            con.constraint_name,
            split_part ( split_part ( con.def[3], ' ', 3 ), '.', 1 ) AS ref_table_schema,
            split_part ( split_part ( con.def[3], ' ', 3 ), '.', 2 ) AS ref_table_name,
            con.def[4] AS ref_column_names,
            con.ref_constraint_name,
            con.match_option,
            con.update_rule,
            con.delete_rule
        FROM referential_constraints con
)
SELECT args.db_name AS table_catalog,
        base.table_schema,
        base.table_name,
        base.column_names,
        base.constraint_name,
        args.db_name AS ref_table_catalog,
        base.ref_table_schema,
        base.ref_table_name,
        base.ref_column_names,
        base.ref_constraint_name,
        base.match_option,
        base.update_rule,
        base.delete_rule,
        'YES' AS is_enforced,
        d.description AS comments
    FROM base
    LEFT OUTER JOIN pg_catalog.pg_description d
        ON ( d.objoid = base.oid )
    CROSS JOIN args
    WHERE ( ( base.table_schema = args.schema_name OR args.ignore_schema )
            AND ( base.table_name = args.table_name OR args.ignore_table ) )
        OR ( ( base.ref_table_schema = args.schema_name OR args.ignore_schema )
            AND ( base.ref_table_name = args.table_name OR args.ignore_table ) )
`
	return m.ReferentialConstraints(db, q, schemaName, tableName)
}
