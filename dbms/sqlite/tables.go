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
	r, err = m.Tables(db, q2, schemaName)
	if err != nil {
		return r, err
	}

	// Wanting (approximate) row counts and not knowing any better way to do so in sqlite...
	for i, v := range r {
		if v.TableType.String == "TABLE" {
			qc := `SELECT count() FROM "` + v.TableName.String + `"`
			rows, err := db.Query(qc)
			if err != nil {
				return r, err
			}
			defer func() {
				if cerr := rows.Close(); cerr != nil && err == nil {
					err = cerr
				}
			}()

			if rows.Next() {
				err = rows.Scan(&r[i].RowCount)
				if err != nil {
					return r, err
				}
			}
		}
	}

	return r, err
}
