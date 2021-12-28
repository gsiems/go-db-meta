package mssql

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
    SELECT coalesce ( '', '' ) AS schema_name,
            coalesce ( '', '' ) AS table_name
)
tab AS (
    SELECT col.constraint_catalog,
            col.constraint_schema,
            col.constraint_name,
            col.table_catalog,
            col.table_schema,
            col.table_name,
            string_agg ( col.column_name, ', ' ) WITHIN GROUP ( ORDER BY col.ordinal_position ) AS table_columns
        FROM information_schema.key_column_usage col
        WHERE col.table_schema NOT IN ( 'INFORMATION_SCHEMA', 'sys' )
            AND substring ( col.table_schema, 1, 3 ) <> 'db_'
            AND substring ( col.table_name, 1, 1 ) <> '#'
        GROUP BY col.constraint_catalog,
            col.constraint_schema,
            col.constraint_name,
            col.table_catalog,
            col.table_schema,
            col.table_name
)
SELECT col.table_catalog,
        col.table_schema,
        col.table_name,
        col.table_columns,
        con.constraint_name,
        ref_col.table_catalog AS ref_table_catalog,
        ref_col.table_schema AS ref_table_schema,
        ref_col.table_name AS ref_table_name,
        ref_col.table_columns AS ref_table_columns,
        con.unique_constraint_name AS ref_constraint_name,
        con.match_option,
        con.update_rule,
        con.delete_rule,
        'YES' AS is_enforced,
        --is_deferrable,
        --initially_deferred,
        NULL AS comments
    FROM information_schema.referential_constraints con
    CROSS JOIN args
    INNER JOIN tab col
        ON ( con.constraint_catalog = col.constraint_catalog
            AND con.constraint_schema = col.constraint_schema
            AND con.constraint_name = col.constraint_name )
    INNER JOIN tab ref_col
        ON ( con.unique_constraint_catalog = ref_col.constraint_catalog
            AND con.unique_constraint_schema = ref_col.constraint_schema
            AND con.unique_constraint_name = ref_col.constraint_name )
    WHERE ( ( ( col.table_schema = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
                AND ( col.table_name = args.table_name OR args.table_name = '' ) )
            OR ( ( ref_col.table_schema = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
                AND ( ref_col.table_name = args.table_name OR args.table_name = '' ) ) )
`
	return m.ReferentialConstraints(db, q, schemaName, tableName)
}
