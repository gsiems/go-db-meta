package sqlite

import (
	"database/sql"
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

// Tables defines the query for obtaining a list of tables and views
// for the (schemaName) parameter and returns the results of
// executing the query
func Tables(db *sql.DB, schemaName string) ([]m.Table, error) {

	var r []m.Table

	q := `
SELECT '%s' AS table_catalog,
        coalesce ( $1, '' ) AS table_schema,
        name AS table_name,
        NULL AS table_owner,
        upper ( type ) AS table_type,
        NULL AS row_count,
        NULL AS comment,
        CASE
            WHEN type = 'view' THEN sql
            END AS view_definition
    FROM sqlite_master
    WHERE type IN ( 'table', 'view' )
        AND tbl_name NOT LIKE '%s'
`

	catName, err := catalogName(db)
	if err != nil {
		return r, err
	}

	q2 := fmt.Sprintf(q, catName, "sqlite_%")
	return m.Tables(db, q2, schemaName)
}
