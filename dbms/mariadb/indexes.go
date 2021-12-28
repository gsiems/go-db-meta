package mariadb

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Indexes defines the query for obtaining a list of indexes
// for the (schemaName, tableName) parameters and returns the results
// of executing the query
func Indexes(db *sql.DB, schemaName, tableName string) ([]m.Index, error) {

	q := `
SELECT stat.table_catalog AS index_catalog,
        stat.index_schema,
        stat.index_name,
        stat.index_type,
        group_concat( stat.column_name
            ORDER BY stat.seq_in_index
            SEPARATOR ', ' ) AS index_columns,
        stat.table_catalog,
        stat.table_schema,
        stat.table_name,
        CASE stat.non_unique
            WHEN 1 THEN 'NO'
            ELSE 'YES'
            END AS is_unique,
        index_comment AS comment
    FROM information_schema.statistics stat
    CROSS JOIN (
        SELECT coalesce ( $1, '' ) AS schema_name,
                coalesce ( $2, '' ) AS table_name
        ) AS args
    WHERE stat.table_schema NOT IN ( 'information_schema', 'mysql', 'performance_schema', 'sys' )
        AND ( stat.table_schema = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
        AND ( stat.table_name = args.table_name OR args.table_name = '' )
    GROUP BY stat.index_schema,
        stat.index_name,
        stat.index_type,
        stat.non_unique,
        stat.table_name
`
	return m.Indexes(db, q, schemaName, tableName)
}
