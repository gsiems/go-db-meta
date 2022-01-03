package sqlite

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Columns obtains the list of columns
// for the (schemaName, tableName) parameters and returns the results
func Columns(db *sql.DB, schemaName, tableName string) ([]m.Column, error) {

	var r []m.Column

	catName, err := catalogName(db)
	if err != nil {
		return r, err
	}

	q := `
SELECT args.table_catalog,
        args.table_schema,
        m.name AS table_name,
        cols.name AS column_name,
        cols.cid AS ordinal_position,
        cols."type" AS data_type,
        CASE
            WHEN cols."notnull" = 1 THEN 'NO'
            ELSE 'YES'
            END AS is_nullable,
        NULL AS column_default,
        NULL AS domain_catalog,
        NULL AS domain_schema,
        NULL AS domain_name,
        NULL AS comment
    FROM sqlite_master AS m
    JOIN pragma_table_info ( m.name ) AS cols
    CROSS JOIN (
        SELECT '` + catName.String + `' AS table_catalog,
                coalesce ( $1, '' ) AS table_schema,
                coalesce ( $2, '' ) AS table_name
        ) AS args
    WHERE m.type IN ( 'table', 'view' )
        AND substr ( m.tbl_name, 1, 7 ) <>  'sqlite_'
        AND ( args.table_name = '' OR args.table_name = m.name )
`
	return m.Columns(db, q, schemaName, tableName)
}
