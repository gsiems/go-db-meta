package pg

import (
	m "github.com/gsiems/go-db-meta/model"
)

// CheckConstraints defines the query for obtaining the check
// constraints for the tables specified by the (tableSchema, tableName)
// parameters and returns the results of executing the query
func CheckConstraints(db *m.DB, tableSchema, tableName string) ([]m.CheckConstraint, error) {

	q := `
WITH args AS (
    SELECT coalesce ( $1, '' ) AS schema_name,
            coalesce ( $2, '' ) AS table_name
)
SELECT current_database() AS table_catalog,
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
        AND ( n.nspname = args.schema_name OR ( args.schema_name = '' AND args.table_name = '' ) )
        AND ( r.relname = args.table_name OR args.table_name = '' )
`
	return db.CheckConstraints(q, tableSchema, tableName)
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
