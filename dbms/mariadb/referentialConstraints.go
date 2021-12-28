package mariadb

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
SELECT col.table_catalog,
        col.table_schema,
        col.table_name,
        group_concat(col.column_name
            order by col.position_in_unique_constraint
            separator ', ') AS table_columns,
        con.constraint_name,
        con.unique_constraint_catalog AS ref_table_catalog,
        col.referenced_table_schema AS ref_table_schema,
        col.referenced_table_name AS ref_table_name,
        group_concat( col.referenced_column_name
            order by col.position_in_unique_constraint
            separator ', ' ) AS ref_table_columns,
        con.unique_constraint_name AS ref_constraint_name,
        con.match_option,
        con.update_rule,
        con.delete_rule,
        'YES' AS is_enforced,
        #is_deferrable,
        #initially_deferred,
        NULL AS comments
    FROM information_schema.referential_constraints con
    JOIN information_schema.key_column_usage col
         ON ( con.constraint_schema = col.table_schema
             AND con.table_name = col.table_name
             AND con.constraint_name = col.constraint_name )
    CROSS JOIN (
        SELECT coalesce ( $1, '' ) AS schema_name,
                coalesce ( $2, '' ) AS table_name
        ) AS args
    WHERE con.constraint_schema NOT IN ( 'information_schema', 'mysql', 'performance_schema', 'sys' )
        AND con.unique_constraint_schema NOT IN ( 'information_schema', 'mysql', 'performance_schema', 'sys' )
        AND ( ( ( col.table_schema = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
                AND ( col.table_name = args.table_name OR args.table_name = '' ) )
            OR ( ( col.referenced_table_schema = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
                AND ( col.referenced_table_name = args.table_name OR args.table_name = '' ) ) )
    GROUP BY col.table_catalog,
        col.table_schema,
        col.table_name,
        con.constraint_name,
        con.unique_constraint_catalog,
        col.referenced_table_schema,
        col.referenced_table_name,
        con.unique_constraint_name,
        con.match_option,
        con.update_rule,
        con.delete_rule
`
	return m.ReferentialConstraints(db, q, schemaName, tableName)
}
