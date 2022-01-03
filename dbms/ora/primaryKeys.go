package ora

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// PrimaryKeys defines the query for obtaining the primary keys
// for the (schemaName, tableName) parameters and returns the results
// of executing the query
func PrimaryKeys(db *sql.DB, schemaName, tableName string) ([]m.PrimaryKey, error) {

	q := `
WITH args AS (
    SELECT :1 AS schema_name,
            :2 AS table_name
        FROM dual
)
SELECT sys_context ( 'userenv', 'DB_NAME' ) AS table_catalog,
        con.owner AS table_schema,
        con.table_name,
        con.constraint_name,
        listagg ( col.column_name, ', ' ) WITHIN GROUP (
            ORDER BY col.position ) AS constraint_columns,
        initcap ( con.status) AS status,
        NULL AS comments
    FROM dba_constraints con
    CROSS JOIN args
    JOIN dba_cons_columns col
        ON ( col.owner = con.owner
            AND col.table_name = con.table_name
            AND col.constraint_name = con.constraint_name )
    WHERE con.constraint_type = 'P'
        AND con.owner NOT IN ( ` + systemTables + ` )
        AND ( col.owner = args.schema_name OR ( args.schema_name IS NULL AND args.table_name IS NULL ) )
        AND ( col.table_name = args.table_name OR args.table_name IS NULL )
    GROUP BY sys_context ( 'userenv', 'DB_NAME' ),
        con.owner,
        con.table_name,
        con.constraint_name,
        initcap ( con.status )
`
	return m.PrimaryKeys(db, q, schemaName, tableName)
}
