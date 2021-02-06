package ora

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Columns defines the query for obtaining a list of columns
// for the (tableSchema, tableName) parameters and returns the results
// of executing the query
func Columns(db *m.DB, tableSchema, tableName string) ([]m.Column, error) {

	q := `
WITH args AS (
    SELECT $1 AS schema_name,
            $2 AS table_name
)
SELECT sys_context ( 'userenv', 'DB_NAME' ) AS table_catalog,
        col.owner AS table_schema,
        col.table_name,
        col.column_name,
        col.column_id AS ordinal_position,
        CASE
            WHEN col.data_type = 'NUMBER'
                AND coalesce ( col.data_precision, 0 ) > 0
                AND coalesce ( col.data_scale, 0 ) >  0
                THEN col.data_type || '(' || col.data_precision || ',' || col.data_scale || ')'
            WHEN col.data_type = 'NUMBER'
                AND coalesce ( col.data_precision, 0 ) > 0
                AND coalesce ( col.data_scale, 0 ) =  0
                THEN col.data_type || '(' || col.data_precision || ')'
            WHEN col.data_type = 'NUMBER'
                AND coalesce ( col.data_length, 0 ) > 0
                THEN col.data_type || '(' || col.data_length || ')'
            WHEN col.data_type IN ( 'CHAR', 'NCHAR', 'VARCHAR2', 'NVARCHAR2' )
                AND coalesce ( col.char_length, 0 ) > 0
                -- TODO: bytes vs. chars
                THEN col.data_type || '(' || col.char_length || ')'
            WHEN col.data_type = 'FLOAT' THEN
                AND coalesce ( col.data_precision, 0 ) > 0
                THEN col.data_type || '(' || col.data_precision || ')'
            WHEN col.data_type IN ( 'RAW', 'UROWID' ) THEN
                AND coalesce ( col.data_length, 0 ) > 0
                THEN col.data_type || '(' || col.data_length || ')'
            ELSE col.data_type
            END AS data_type,
        col.nullable AS is_nullable,
        col.data_default,
        case ( NULL AS varchar2 ( 1 ) ) AS DomainCatalog,
        case ( NULL AS varchar2 ( 1 ) ) AS DomainSchema,
        case ( NULL AS varchar2 ( 1 ) ) AS DomainName,
        -- UdtCatalog,
        -- UdtSchema,
        -- UdtName,
        cmt.comments
    FROM all_tab_columns col
    CROSS JOIN args
    LEFT OUTER JOIN all_col_comments cmt
         ON ( col.owner = cmt.owner
            AND col.table_name = cmt.table_name
            AND col.column_name = cmt.column_name )
    WHERE col.owner NOT IN ( %s )
        AND col.owner NOT LIKE '%s'
        AND ( col.owner = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
        AND ( col.table_name = args.table_name OR args.table_name = '' )
    ORDER BY col.owner,
        col.table_name
        col.column_id
`
	q2 := fmt.Sprintf(q, systemTables, "%$%")
	return db.Columns(q2, tableSchema, tableName)
}
