package sqlite

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Indexes defines the query for obtaining a list of indexes
// for the (tableSchema, tableName) parameters and returns the results
// of executing the query
func Indexes(db *m.DB, tableSchema, tableName string) ([]m.Index, error) {

	q := `
SELECT '%s' AS index_catalog,
        con.table_schema AS index_schema,
        con.index_name,
        '' AS index_type,
        group_concat ( con.column_name, ', ' ) AS index_columns,
        '%s' AS table_catalog,
        con.table_schema,
        con.table_name,
        CASE
            WHEN max ( con.origin ) IN ( 'pk', 'u' ) THEN 'Y'
            ELSE 'N'
            END AS is_unique,
        -- status
        '' AS comments
    FROM (
        SELECT tab.name AS table_name,
                args.table_schema,
                idx.name AS index_name,
                idx."unique",
                idx.origin,
                idx.partial,
                col.name AS column_name,
                col.seqno AS ordinal_position
            FROM sqlite_master AS tab
            CROSS JOIN (
                SELECT coalesce ( $1, '' ) AS table_schema,
                        coalesce ( $2, '' ) AS table_name
                ) AS args
            JOIN pragma_index_list ( tab.name ) AS idx
            JOIN pragma_index_info ( idx.name ) AS col
            WHERE tab.type = 'table'
                AND m.tbl_name NOT LIKE '%s'
                AND ( args.table_name = '' OR args.table_name = m.name )
            ORDER BY tab.name,
                idx.name,
                col.seqno
        ) AS con
    GROUP BY con.table_schema,
        con.table_name,
        con.index_name
`
	catName, err := catalogName(db)
	if err != nil {
		return d, err
	}

	q2 := fmt.Sprintf(q, catName, catName, "sqlite_%")
	return db.Indexes(q2, tableSchema, tableName)
}
