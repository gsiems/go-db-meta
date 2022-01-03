package ora

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Tables defines the query for obtaining a list of tables and views
// for the (schemaName) parameter and returns the results of
// executing the query
func Tables(db *sql.DB, schemaName string) ([]m.Table, error) {

	q := `
WITH args AS (
    SELECT :1 AS schema_name
        FROM dual
),
tab AS (
    SELECT obj.owner,
            obj.object_name,
            min ( obj.object_type ) AS object_type
        FROM dba_objects obj
        CROSS JOIN args
        WHERE ( obj.owner = args.schema_name
                OR args.schema_name IS NULL )
            AND obj.owner NOT IN ( ` + systemTables + ` )
            AND obj.object_type IN ( 'TABLE', 'VIEW', 'MATERIALIZED VIEW' )
        GROUP BY obj.owner,
            obj.object_name
)
SELECT sys_context ( 'userenv', 'DB_NAME' ) AS table_catalog,
        tab.owner AS table_schema,
        tab.object_name AS table_name,
        tab.owner AS table_owner,
        tab.object_type AS table_type,
        dt.num_rows AS row_count,
        cmt.comments,
        coalesce ( mv.query, dv.text ) AS view_definition
    FROM tab
    LEFT JOIN dba_tables dt
        ON ( dt.owner = tab.owner
            AND dt.table_name = tab.object_name )
    LEFT JOIN dba_snapshots mv
        ON ( mv.owner = tab.owner
            AND mv.table_name = tab.object_name )
    LEFT JOIN dba_views dv
        ON ( dv.owner = tab.owner
            AND dv.view_name = tab.object_name )
    LEFT JOIN dba_tab_comments cmt
        ON ( tab.owner = cmt.owner
            AND tab.object_name = cmt.table_name )
`
	return m.Tables(db, q, schemaName)
}
