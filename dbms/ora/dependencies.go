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
    SELECT DISTINCT owner AS object_schema,
            name AS object_name,
            CASE
                WHEN d.type = 'PACKAGE BODY' THEN 'PACKAGE'
                WHEN d.type = 'TYPE BODY' THEN 'TYPE'
                ELSE d.type
                END AS object_type,
            referenced_owner AS dep_object_schema,
            referenced_name AS dep_object_name,
            CASE
                WHEN d.referenced_type = 'PACKAGE BODY' THEN 'PACKAGE'
                WHEN d.referenced_type = 'TYPE BODY' THEN 'TYPE'
                ELSE d.referenced_type
                END AS dep_object_type
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
SELECT DISTINCT sys_context ( 'userenv', 'DB_NAME' ) AS object_catalog,
        d.object_schema,
        d.object_name,
        d.object_schema AS object_owner,
        coalesce ( mv.object_type, d.object_type ) AS object_type,
        sys_context ( 'userenv', 'DB_NAME' ) AS dep_object_catalog,
        d.dep_object_schema,
        d.dep_object_name,
        d.dep_object_schema AS dep_object_owner,
        coalesce ( rmv.object_type, d.dep_object_type ) AS dep_object_type
    FROM base d
    LEFT OUTER JOIN dba_objects mv
        ON ( mv.owner = d.object_schema
            AND mv.object_name = d.object_name
            AND mv.object_type = 'MATERIALIZED VIEW' )
    LEFT OUTER JOIN dba_objects rmv
        ON ( rmv.owner = d.dep_object_schema
            AND rmv.object_name = d.dep_object_name
            AND rmv.object_type = 'MATERIALIZED VIEW' )
`
	return m.Dependencies(db, q, schemaName, objectName)
}
