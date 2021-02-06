package sqlite

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

// Tables defines the query for obtaining a list of tables and views
// for the (schema) parameter and returns the results of
// executing the query
func Tables(db *m.DB, schema string) ([]m.Table, error) {

	var d []m.Table

	q := `
SELECT '%s' AS table_catalog,
        coalesce ( $1, '' ) AS table_schema,
        name AS table_name,
        NULL AS table_owner,
        type AS table_type,
        NULL AS row_count,
        NULL AS comment,
        CASE
            WHEN type = 'view' THEN sql
            END AS view_definition
    FROM sqlite_master
    WHERE type IN ( 'table', 'view' )
        AND tbl_name NOT LIKE 'sqlite_%'
`

	catName, err := catalogName(db)
	if err != nil {
		return d, err
	}

	q2 = fmt.Sprintf(q, catName)
	return db.Tables(q2, schema)
}
