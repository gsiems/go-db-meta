package sqlite

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// PrimaryKeys defines the query for obtaining the primary keys
// for the (schemaName, tableName) parameters and returns the results
// of executing the query
func PrimaryKeys(db *sql.DB, schemaName, tableName string) ([]m.PrimaryKey, error) {

	var r []m.PrimaryKey

	// Primary key names may show in the .schema command output but not,
	// apparently in the output of any pragma queries.

	catName, err := catalogName(db)
	if err != nil {
		return r, err
	}

	q := `
SELECT pk_col.table_catalog,
        pk_col.table_schema,
        pk_col.table_name,
        'pk_' || pk_col.table_name AS constraint_name,
        group_concat ( pk_col.column_name, ', ' ) AS constraint_columns,
        'Enabled' AS status,
        '' AS comments
    FROM (
        SELECT args.table_catalog,
                m.name as table_name,
                args.table_schema,
                col.name AS column_name,
                col.pk AS ordinal_position
            FROM sqlite_master AS m
            JOIN pragma_table_info ( m.name ) AS col
            CROSS JOIN (
                SELECT '` + catName.String + `' AS table_catalog,
                        coalesce ( $1, '' ) AS table_schema,
                        coalesce ( $2, '' ) AS table_name
                ) AS args
            WHERE m.type = 'table'
                AND substr ( m.tbl_name, 1, 7 ) <>  'sqlite_'
                AND ( args.table_name = '' OR args.table_name = m.name )
                AND col.pk > 0
            ORDER BY m.name,
                col.pk
        ) AS pk_col
    GROUP BY pk_col.table_schema,
        pk_col.table_name
`
	return m.PrimaryKeys(db, q, schemaName, tableName)
}
