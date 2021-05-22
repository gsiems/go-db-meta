package mariadb

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Tables defines the query for obtaining a list of tables and views
// for the (schema) parameter and returns the results of
// executing the query
func Tables(db *m.DB, schema string) ([]m.Table, error) {

	// NB that mariadb doesn't appear to support CTEs
	q := `
SELECT t.table_catalog,
        t.table_schema,
        t.table_name,
        NULL AS table_owner,
        t.table_type,
        t.table_rows,
        NULL AS comment, -- I have no idea how to retrieve this
        v.view_definition
    FROM information_schema.tables t
    CROSS JOIN (
            SELECT coalesce ( $1, '' ) AS schema_name
    ) AS args
    LEFT JOIN information_schema.views v
        ON ( v.table_schema = t.table_schema
            AND v.table_name = t.table_name )
    WHERE t.table_schema NOT IN ( 'information_schema', 'mysql', 'performance_schema' )
        AND ( t.table_schema = args.schema_name
            OR args.schema_name = '' )
`
	return db.Tables(q, schema)
}
