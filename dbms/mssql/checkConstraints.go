package mssql

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// CheckConstraints defines the query for obtaining the check
// constraints for the tables specified by the (schemaName, tableName)
// parameters and returns the results of executing the query
func CheckConstraints(db *sql.DB, schemaName, tableName string) ([]m.CheckConstraint, error) {

	q := `
WITH args AS (
    SELECT coalesce ( $1, '' ) AS schema_name,
            coalesce ( $2, '' ) AS table_name
)
SELECT tab.table_catalog,
        tab.table_schema,
        tab.table_name,
        con.constraint_name,
        con.check_clause,
        'Enabled' AS status, -- can be disabled?
        NULL AS comments
FROM information_schema.check_constraints con
INNER JOIN information_schema.constraint_table_usage tab
    ON ( con.constraint_catalog = tab.constraint_catalog
        AND con.constraint_schema = tab.constraint_schema
        AND con.constraint_name = tab.constraint_name )
    WHERE tab.table_schema NOT IN ( 'INFORMATION_SCHEMA', 'sys' )
        AND substring ( tab.table_schema, 1, 3 ) <> 'db_'
        AND substring ( tab.table_name, 1, 1 ) <> '#'
        AND ( tab.table_schema = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
        AND ( tab.table_name = args.table_name OR args.table_name = '' )
`
	return m.CheckConstraints(db, q, schemaName, tableName)
}
