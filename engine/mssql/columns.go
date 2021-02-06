package mssql

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
SELECT col.table_catalog,
        col.table_schema,
        col.table_name,
        col.column_name,
        col.ordinal_position,
        col.column_default,
        col.is_nullable,
        CASE
            WHEN col.datetime_precision IS NOT NULL THEN
                THEN col.data_type || '(' || col.datetime_precision || ')'
            WHEN col.numeric_scale IS NOT NULL THEN
                THEN col.data_type || '(' || col.numeric_precision || ',' || col.numeric_scale || ')'
            WHEN col.numeric_precision IS NOT NULL AND coalesce ( col.numeric_precision, 0 ) > 0 THEN
                THEN col.data_type || '(' || col.numeric_precision || ',' || col.numeric_scale || ')'
            WHEN col.numeric_precision IS NOT NULL THEN
                THEN col.data_type || '(' || col.numeric_precision || ')'
            WHEN col.data_type IN ( 'char', 'varchar' )
                AND coalesce ( col.character_maximum_length, 0 ) > 0
                -- TODO: bytes vs. chars
                THEN col.data_type || '(' || col.character_maximum_length || ')'
            ELSE col.data_type
            END AS data_type,
        col.domain_catalog,
        col.domain_schema,
        col.domain_name,
        convert ( varchar ( 8000 ), xp.value ) AS comments
    FROM information_schema.columns col
    CROSS JOIN args
    OUTER APPLY ::fn_listextendedproperty ( 'MS_Description', 'schema', col.table_schema, 'table', col.table_name, 'column', col.column_name ) xp
    WHERE col.table_schema NOT IN ( 'INFORMATION_SCHEMA', 'sys' )
        AND substring ( col.table_schema, 1, 3 ) <> 'db_'
        AND substring ( col.table_name, 1, 1 ) <> '#'
        AND ( col.table_schema = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
        AND ( col.table_name = args.table_name OR args.table_name = '' )

`
	return db.Columns(q, tableSchema, tableName)
}
