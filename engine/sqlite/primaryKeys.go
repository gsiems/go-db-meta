package sqlite

import (
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

// PrimaryKeys defines the query for obtaining the primary keys
// for the (tableSchema, tableName) parameters and returns the results
// of executing the query
func PrimaryKeys(db *m.DB, tableSchema, tableName string) ([]m.PrimaryKey, error) {

	// Primary key names may show in the .schema command output but not,
	// apparently in the output of any pragma queries.

	q := `
SELECT '%s' AS table_catalog,
        x.table_schema,
        x.table_name,
        '' AS constraint_name,
        group_concat ( x.column_name, ', ' ) AS constraint_columns,
        'Enabled' AS status,
        '' AS comments
    FROM (
        SELECT m.name as table_name,
                args.table_schema,
                col.name AS column_name,
                col.pk AS ordinal_position
            FROM sqlite_master AS m
            JOIN pragma_table_info ( m.name ) AS col
            CROSS JOIN (
                SELECT coalesce ( $1, '' ) AS table_schema,
                        coalesce ( $2, '' ) AS table_name
                ) AS args
            WHERE m.type = 'table'
                AND m.tbl_name NOT LIKE 'sqlite_%'
                AND ( args.table_name = '' OR args.table_name = m.name )
                AND col.pk > 0
            ORDER BY m.name,
                col.pk
        ) AS x
    GROUP BY x.table_name
`

	catName, err := catalogName(db)
	if err != nil {
		return d, err
	}

	q2 := fmt.Sprintf(q, catName)
	return db.PrimaryKeys(q2, tableSchema, tableName)
}
