package null

import (
	m "github.com/gsiems/go-db-meta/model"
)

// PrimaryKeys defines the query for obtaining the primary keys
// for the (tableSchema, tableName) parameters and returns the results
// of executing the query
func PrimaryKeys(db *m.DB, tableSchema, tableName string) ([]m.PrimaryKey, error) {

	q := `
WITH args AS (
    SELECT coalesce ( $1, '' ) AS schema_name,
            coalesce ( $2, '' ) AS table_name
)
SELECT con.table_catalog,
        con.table_schema,
        con.table_name,
        con.constraint_name,
        string_agg ( col.column_name, ', ' ) WITHIN GROUP ( ORDER BY col.ordinal_position ) AS constraint_columns,
        'Enabled' AS constraint_status,
        NULL AS comments
    FROM information_schema.table_constraints con
    CROSS JOIN args
    JOIN information_schema.key_column_usage col
        ON ( col.constraint_catalog = con.constraint_catalog
            AND col.constraint_schema = con.constraint_schema
            AND col.constraint_name = con.constraint_name )
    WHERE con.table_schema NOT IN ( 'INFORMATION_SCHEMA', 'sys' )
        AND substring ( con.table_name, 1, 1 ) <> '#'
        AND con.constraint_type = 'PRIMARY KEY'
        AND ( con.table_schema = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
        AND ( con.table_name = args.table_name OR args.table_name = '' )
`
	return db.PrimaryKeys(q, tableSchema, tableName)
}
