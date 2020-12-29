package mssql

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Tables defines the query for obtaining a list of tables and views
// for the (schema) parameter and returns the results of
// executing the query
func Tables(db *m.DB, schema string) ([]m.Table, error) {

	q := `
SELECT tabs.table_catalog,
        tabs.table_schema,
        tabs.table_name,
        tabs.table_schema AS table_owner,
        tabs.table_type,
        NULL AS row_count,
        NULL AS comment,
        v.view_definition
    FROM information_schema.tables tabs
    LEFT JOIN information_schema.views v
        ON ( v.table_catalog = tabs.table_catalog
            AND v.table_schema = tabs.table_schema
            AND v.table_name = tabs.table_name )
    WHERE substring ( tabs.table_name, 1, 1 ) <> '#'
        AND ( tabs.table_schema = $1
            OR $1 = '' )
`
	return db.Tables(q, schema)
}
