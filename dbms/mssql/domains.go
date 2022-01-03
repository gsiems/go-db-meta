package mssql

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Domains defines the query for obtaining a list of domains
// for the (schemaName) parameters and returns the results
// of executing the query
func Domains(db *sql.DB, schemaName string) ([]m.Domain, error) {

	q := `
WITH args AS (
    SELECT coalesce ( $1, '' ) AS schema_name
)
SELECT dom.domain_catalog,
        dom.domain_schema,
        dom.domain_name,
        dom.domain_name AS domain_owner,
        CASE
            WHEN dom.datetime_precision IS NOT NULL THEN
                THEN dom.data_type || '(' || dom.datetime_precision || ')'
            WHEN dom.numeric_scale IS NOT NULL THEN
                THEN dom.data_type || '(' || dom.numeric_precision || ',' || dom.numeric_scale || ')'
            WHEN dom.numeric_precision IS NOT NULL AND coalesce ( dom.numeric_precision, 0 ) > 0 THEN
                THEN dom.data_type || '(' || dom.numeric_precision || ',' || dom.numeric_scale || ')'
            WHEN dom.numeric_precision IS NOT NULL THEN
                THEN dom.data_type || '(' || dom.numeric_precision || ')'
            WHEN dom.data_type IN ( 'char', 'varchar' )
                AND coalesce ( dom.character_maximum_length, 0 ) > 0
                -- TODO: bytes vs. chars
                THEN dom.data_type || '(' || dom.character_maximum_length || ')'
            ELSE dom.data_type
            END AS data_type,
        dom.domain_default,
        NULL AS check_clause,
        NULL AS comments,
    FROM information_schema.domains dom
    CROSS JOIN args
    WHERE dom.domain_schema NOT IN ( 'INFORMATION_SCHEMA', 'sys' )
        AND substring ( dom.domain_schema, 1, 3 ) <> 'db_'
        AND ( dom.table_schema = args.schema_name
            OR args.schema_name = '' )
`

	return m.Domains(db, q, schemaName)
}
