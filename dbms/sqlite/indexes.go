package sqlite

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Indexes defines the query for obtaining a list of indexes
// for the (schemaName, tableName) parameters and returns the results
// of executing the query
func Indexes(db *sql.DB, schemaName, tableName string) ([]m.Index, error) {

	var r []m.Index

	catName, err := catalogName(db)
	if err != nil {
		return r, err
	}

	q := `
SELECT con.table_catalog AS index_catalog,
        con.table_schema AS index_schema,
        con.index_name,
        '' AS index_type,
        group_concat ( con.column_name, ', ' ) AS index_columns,
        con.table_catalog,
        con.table_schema,
        con.table_name,
        CASE
            WHEN max ( con.origin ) IN ( 'pk', 'u' ) THEN 'YES'
            ELSE 'NO'
            END AS is_unique,
        -- status
        '' AS comments
    FROM (
        SELECT args.table_catalog,
                tab.name AS table_name,
                args.table_schema,
                idx.name AS index_name,
                idx."unique",
                idx.origin,
                idx.partial,
                col.name AS column_name,
                col.seqno AS ordinal_position
            FROM sqlite_master AS tab
            CROSS JOIN (
                SELECT '` + catName.String + `' AS table_catalog,
                        coalesce ( $1, '' ) AS table_schema,
                        coalesce ( $2, '' ) AS table_name
                ) AS args
            JOIN pragma_index_list ( tab.name ) AS idx
            JOIN pragma_index_info ( idx.name ) AS col
            WHERE tab.type = 'table'
                AND substr ( tab.name, 1, 7 ) <>  'sqlite_'
                AND ( args.table_name = '' OR args.table_name = tab.name )
            ORDER BY tab.name,
                idx.name,
                col.seqno
        ) AS con
    GROUP BY con.table_schema,
        con.table_name,
        con.index_name
`
	return m.Indexes(db, q, schemaName, tableName)
}
