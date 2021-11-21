package ora

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Indexes defines the query for obtaining a list of indexes
// for the (tableSchema, tableName) parameters and returns the results
// of executing the query
func Indexes(db *m.DB, tableSchema, tableName string) ([]m.Index, error) {

	q := `
WITH args AS (
    SELECT $1 AS schema_name,
            $2 AS table_name
        FROM dual
)
SELECT sys_context ( 'userenv', 'DB_NAME' ) AS index_catalog,
        idx.owner AS index_schema,
        idx.index_name,
        idx.index_type,
        listagg ( col.column_name, ', ' ) WITHIN GROUP (
            ORDER BY col.column_position ) AS index_columns,
        sys_context ( 'userenv', 'DB_NAME' ) AS table_catalog,
        idx.table_owner AS table_schema,
        idx.table_name,
        CASE idx.uniqueness
            WHEN 'UNIQUE' THEN 'Y'
            ELSE 'N'
            END AS is_unique,
        --col.descend,
        --ie.column_expression,
        --idx.status,
        NULL AS comments
    FROM dba_indexes idx
    CROSS JOIN args
    INNER JOIN dba_ind_columns col
        ON ( col.index_owner = idx.owner
            AND idx.index_name = col.index_name )
    LEFT OUTER JOIN dba_constraints con
        ON ( con.owner = idx.owner
            AND con.table_name = idx.table_name
            AND con.constraint_name = idx.index_name )
    --LEFT OUTER JOIN dba_ind_expressions ie
    --    ON ( ie.index_owner = col.index_owner
    --        AND ie.index_name = col.index_name
    --        AND ie.table_owner = col.table_owner
    --        AND ie.table_name = col.table_name
    --        AND ie.column_position = col.column_position )
    WHERE idx.table_owner NOT IN ( %s )
        AND idx.table_owner NOT LIKE '%s'
        AND ( idx.table_owner = args.schema_name OR ( args.schema_name IS NULL AND args.table_name IS NULL ) )
        AND ( idx.table_name = args.table_name OR args.table_name IS NULL )
`
	q2 := fmt.Sprintf(q, systemTables, "%$%")
	return db.Indexes(q2, tableSchema, tableName)
}
