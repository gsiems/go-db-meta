package ora

import (
	m "github.com/gsiems/go-db-meta/model"
)

// PrimaryKeys defines the query for obtaining the primary keys
// for the (tableSchema, tableName) parameters and returns the results
// of executing the query
func PrimaryKeys(db *m.DB, tableSchema, tableName string) ([]m.PrimaryKey, error) {

	q := `
WITH args AS (
    SELECT $1 AS schema_name,
            $2 AS table_name
)
SELECT sys_context ( 'userenv', 'DB_NAME' ) AS table_catalog,
        con.owner AS table_schema,
        con.table_name,
        con.constraint_name,
        listagg ( col.column_name, ', ' ) WITHIN GROUP (
            ORDER BY position ) AS constraint_columns,
        CASE con.status
            WHEN 'ENABLED' THEN 'Enabled'
            WHEN 'DISABLED' THEN 'Disabled'
            ELSE status
            END AS status,
        NULL AS comments
    FROM all_constraints con
    CROSS JOIN args
    JOIN all_cons_columns col
        ON ( col.owner = con.owner
            AND col.table_name = con.table_name
            AND col.constraint_name = con.constraint_name )
    WHERE con.constraint_type = 'P'
        AND con.owner NOT IN (%s )
        AND con.owner NOT LIKE '%s'
        AND ( col.owner = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
        AND ( col.table_name = args.table_name OR args.table_name = '' )
`
	q2 := fmt.Sprintf(q, systemTables, "%$%")
	return db.PrimaryKeys(q2, tableSchema, tableName)
}
