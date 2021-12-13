package mssql

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Tables defines the query for obtaining a list of tables and views
// for the (schema) parameter and returns the results of
// executing the query
func Tables(db *m.DB, schema string) ([]m.Table, error) {

	q := `
WITH args AS (
    SELECT coalesce ( $1, '' ) AS schema_name
)
SELECT tab.table_catalog,
        tab.table_schema,
        tab.table_name,
        tab.table_schema AS table_owner,
        upper ( tab.table_type ) AS table_type,
        NULL AS row_count,
        convert ( varchar ( 8000 ), xp.value ) AS comments,
        v.view_definition
    FROM information_schema.tables tab
    CROSS JOIN args
    LEFT JOIN information_schema.views v
        ON ( v.table_catalog = tab.table_catalog
            AND v.table_schema = tab.table_schema
            AND v.table_name = tab.table_name )
    OUTER APPLY ::fn_listextendedproperty ( 'MS_Description', 'schema', tab.table_schema, 'table', tab.table_name ) xp
    WHERE tab.table_schema NOT IN ( 'INFORMATION_SCHEMA', 'sys' )
        AND substring ( tab.table_schema, 1, 3 ) <> 'db_'
        AND substring ( tab.table_name, 1, 1 ) <> '#'
        AND ( tab.table_schema = args.schema_name
            OR args.schema_name = '' )
`
	return db.Tables(q, schema)
}
