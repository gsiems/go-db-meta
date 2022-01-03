package mariadb

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Columns defines the query for obtaining a list of columns
// for the (schemaName, tableName) parameters and returns the results
// of executing the query
func Columns(db *sql.DB, schemaName, tableName string) ([]m.Column, error) {

	q := `
SELECT col.table_catalog,
        col.table_schema,
        col.table_name,
        col.column_name,
        col.ordinal_position,
        col.column_default,
        col.is_nullable,
        CASE
            WHEN col.datetime_precision IS NOT NULL
                THEN concat( col.data_type, "(", col.datetime_precision, ")" )
            WHEN col.numeric_scale IS NOT NULL
                THEN concat( col.data_type, "(", col.numeric_precision, ",", col.numeric_scale, ")" )
            WHEN col.numeric_precision IS NOT NULL AND coalesce ( col.numeric_precision, 0 ) > 0
                THEN concat( col.data_type, "(", col.numeric_precision, ",", col.numeric_scale, ")" )
            WHEN col.numeric_precision IS NOT NULL
                THEN concat( col.data_type, "(", col.numeric_precision, ")" )
            WHEN col.data_type IN ( "char", "varchar" )
                AND coalesce ( col.character_maximum_length, 0 ) > 0
                -- TODO: bytes vs. chars
                THEN concat( col.data_type, "(", col.character_maximum_length, ")" )
            ELSE col.data_type
            END AS data_type,
        NULL AS DomainCatalog,
        NULL AS DomainSchema,
        NULL AS DomainName,
        #UdtCatalog,
        #UdtSchema,
        #UdtName,
        col.column_comment AS comments
    FROM information_schema.columns col
    CROSS JOIN (
            SELECT coalesce ( ?, '' ) AS schema_name,
                    coalesce ( ?, '' ) AS table_name
    ) AS args
    WHERE col.table_schema NOT IN ( 'information_schema', 'mysql', 'performance_schema', 'sys' )
        AND ( col.table_schema = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
        AND ( col.table_name = args.table_name OR args.table_name = '' )
`
	return m.Columns(db, q, schemaName, tableName)
}
