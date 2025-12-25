package pg

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// CheckConstraints defines the query for obtaining the check
// constraints for the tables specified by the (schemaName, tableName)
// parameters and returns the results of executing the query
func CheckConstraints(db *sql.DB, schemaName, tableName string) ([]m.CheckConstraint, error) {

	q := `
WITH args AS (
    SELECT current_database () AS db_name,
            coalesce ( $1, '' ) AS schema_name,
            coalesce ( $2, '' ) AS table_name,
            coalesce ( $1, $2, '' ) = '' AS ignore_schema,
            coalesce ( $2, '' ) = '' AS ignore_table
)
SELECT args.db_name AS table_catalog,
        n.nspname AS table_schema,
        r.relname AS table_name,
        con.conname AS constraint_name,
        substring ( pg_get_constraintdef ( con.oid ), 7 ) AS check_clause,
        'Enabled' AS status,
        d.description AS comments
    FROM pg_class r
    CROSS JOIN args
    INNER JOIN pg_namespace n
        ON ( n.oid = r.relnamespace )
    INNER JOIN pg_constraint con
        ON ( con.conrelid = r.oid )
    LEFT OUTER JOIN pg_description d
        ON ( d.objoid = con.oid )
    WHERE r.relkind = 'r'
        AND con.contype = 'c'
        AND n.nspname <> 'information_schema'
        AND n.nspname !~ '^pg_'
        AND ( n.nspname = args.schema_name OR args.ignore_schema )
        AND ( r.relname = args.table_name OR args.ignore_table )
`
	return m.CheckConstraints(db, q, schemaName, tableName)
}

/*
SELECT conname -- pg_get_constraintdef(oid), *
FROM   pg_constraint c
JOIN   pg_attribute  a ON a.attrelid = c.conrelid     -- !
                      AND a.attnum   = ANY(c.conkey)  -- !
WHERE  c.conrelid = 'hypothetical_table'::regclass
AND    c.contype = 'c'  -- c = check constraint
AND    a.attname = 'some_col';
*/
