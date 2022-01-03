package mariadb

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// UniqueConstraints defines the query for obtaining a list of unique
// constraints for the (schemaName, tableName) parameters and returns the
// results of executing the query
func UniqueConstraints(db *sql.DB, schemaName, tableName string) ([]m.UniqueConstraint, error) {

	q := `
SELECT tab.table_catalog,
        tab.table_schema,
        tab.table_name,
        con.constraint_name,
        group_concat( col.column_name
            ORDER BY col.ordinal_position
            SEPARATOR ', ' ) AS constraint_columns,
        NULL AS status, -- 'Enabled'??
        NULL AS comment
    FROM information_schema.tables tab
    CROSS JOIN (
        SELECT coalesce ( ?, '' ) AS schema_name,
                coalesce ( ?, '' ) AS table_name
        ) AS args
    JOIN information_schema.table_constraints con
        ON ( con.table_schema = tab.table_schema
            AND  con.table_name = tab.table_name )
    JOIN information_schema.key_column_usage col
        ON ( con.constraint_schema = col.constraint_schema
            AND con.constraint_name = col.constraint_name
            AND con.table_name = col.table_name )
    WHERE con.constraint_type = 'UNIQUE'
        AND tab.table_schema NOT IN ( 'information_schema', 'mysql', 'performance_schema', 'sys' )
        AND ( tab.table_schema = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
        AND ( tab.table_name = args.table_name OR args.table_name = '' )
    GROUP BY tab.table_catalog,
        tab.table_schema,
        tab.table_name,
        con.constraint_name
`
	return m.UniqueConstraints(db, q, schemaName, tableName)
}
