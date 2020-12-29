package ora

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Tables defines the query for obtaining a list of tables and views
// for the (schema) parameter and returns the results of
// executing the query
func Tables(db *m.DB, schema string) ([]m.Table, error) {

	q := `
WITH tabs AS (
    SELECT owner,
            object_name,
            min ( object_type ) AS object_type
        FROM dba_objects obj
        WHERE ( owner = $1
                OR $1 = '' )
            AND owner NOT IN (
                    'APPQOSSYS', 'AWR_STAGE', 'CSMIG', 'CTXSYS', 'DBSNMP',
                    'DIP', 'DMSYS', 'DSSYS', 'EXFSYS', 'LBACSYS', 'MDSYS',
                    'OLAPSYS', 'ORACLE_OCM', 'ORDPLUGINS', 'ORDSYS', 'OUTLN',
                    'PERFSTAT', 'PUBLIC', 'SQLTXPLAIN', 'SYS', 'SYSMAN',
                    'SYSTEM', 'TRACESVR', 'TSMSYS', 'WMSYS', 'XDB' )
            AND owner NOT LIKE '%$%'
            AND object_type IN ( 'TABLE', 'VIEW', 'MATERIALIZED VIEW' )
        GROUP BY owner,
            object_name
)
SELECT sys_context ( 'userenv', 'DB_NAME' ) AS table_catalog,
        tabs.owner AS table_schema,
        tabs.object_name AS table_name
        tabs.owner AS table_owner,
        tabs.object_type AS table_type,
        dt.num_rows AS row_count,
        cmt.comments,
        coalesce ( mv.query, dv.text ) AS view_definition
    FROM tabs
    LEFT JOIN dba_tables dt
        ON ( dt.owner = tabs.owner
            AND dt.table_name = tabs.object_name )
    LEFT JOIN dba_snapshots mv
        ON ( mv.owner = tabs.owner
            AND mv.table_name = tabs.object_name )
    LEFT JOIN dba_views dv
        ON ( dv.owner = tabs.owner
            AND dv.view_name = tabs.object_name )
    LEFT JOIN dba_tab_comments cmt
        ON ( tabs.owner = cmt.owner
            AND tabs.object_name = cmt.table_name )
`
	return db.Tables(q, schema)
}
