package ora

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// Dependencies defines the query for obtaining a list of dependencies
// for the (schemaName, objectName) parameters and returns the results
// of executing the query
func Dependencies(db *sql.DB, schemaName, objectName string) ([]m.Dependency, error) {

	/*

		NOTE: Listing dependencies of triggers on their attached table is useless
		information. Triggers that are dependent on other tables however ARE of
		interest.

		Oddly enough, all_dependencies won't return dependencies for materialized
		views whereas dba_dependencies does. Therefore, we want to use
		dba_dependencies if possible.

		ASSERTION: Names for tables, views, materialized views, functions,
		packages, and procedures are unique.

		PROBLEM: The underlying table for materialized views does share the same
		name as the materialized view-- thereby causing apparent duplicates.
		PACKAGE/PACKAGE BODY and TYPE/TYPE BODY behave similarly.

		TODO: Problem Views/Materialized views that are dependent on another
		view/materialized view will show both the other view/materialized view AND
		the dependencies for the other view/materialized view. (Essentially showing
		two levels of dependencies)

	*/

	q := `
WITH args AS (
    SELECT :1 AS schema_name,
            :2 AS object_name
        FROM dual
),
base AS (
    SELECT DISTINCT owner,
            name,
            referenced_owner,
            referenced_name,
            CASE
                WHEN d.type = 'PACKAGE BODY' THEN 'PACKAGE'
                WHEN d.type = 'TYPE BODY' THEN 'TYPE'
                ELSE d.type
                END AS type,
            CASE
                WHEN d.referenced_type = 'PACKAGE BODY' THEN 'PACKAGE'
                WHEN d.referenced_type = 'TYPE BODY' THEN 'TYPE'
                ELSE d.referenced_type
                END AS referenced_type
        FROM dba_dependencies d
        CROSS JOIN args
        WHERE d.referenced_type <> 'MATERIALIZED VIEW'
            AND d.owner NOT IN ( ` + systemTables + ` )
            AND d.referenced_owner NOT IN ( ` + systemTables + ` )
            AND ( d.owner || '.' || d.name <> d.referenced_owner || '.' || d.referenced_name )
            AND ( ( ( d.owner = args.schema_name OR ( args.schema_name IS NULL AND args.object_name IS NULL ) )
                    AND ( d.name = args.object_name OR args.object_name IS NULL ) )
                OR ( ( d.referenced_owner = args.schema_name OR ( args.schema_name IS NULL AND args.object_name IS NULL ) )
                    AND ( d.referenced_name = args.object_name OR args.object_name IS NULL ) ) )
)
SELECT DISTINCT d.owner,
        d.name,
        coalesce ( mv.object_type, d.type ) AS type,
        d.referenced_owner,
        d.referenced_name,
        coalesce ( rmv.object_type, d.referenced_type ) AS referenced_type
    FROM base d
    LEFT OUTER JOIN dba_objects mv
        ON ( mv.owner = d.owner
            AND mv.object_name = d.name
            AND mv.object_type = 'MATERIALIZED VIEW' )
    LEFT OUTER JOIN dba_objects rmv
        ON ( rmv.owner = d.referenced_owner
            AND rmv.object_name = d.referenced_name
            AND rmv.object_type = 'MATERIALIZED VIEW' )
`
	return m.Dependencies(db, q, schemaName, objectName)
}
