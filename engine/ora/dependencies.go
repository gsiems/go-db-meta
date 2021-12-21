package ora

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Dependencies defines the query for obtaining a list of dependencies
// for the (objectSchema, objectName) parameters and returns the results
// of executing the query
func Dependencies(db *m.DB, objectSchema, objectName string) ([]m.Dependency, error) {

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

	q := ``

//	q2 := fmt.Sprintf(q, systemTables, systemTables)
	return db.Dependencies(q, objectSchema, objectName)
}


/*


WITH args AS (
    SELECT :1 AS schema_name,
            $2 AS object_name
        FROM dual
)
SELECT d.owner,
        d.name,
        d.type,
        d.referenced_owner,
        d.referenced_name,
        d.referenced_type
        coalesce ( mv.referenced_type, d.referenced_type ) AS referenced_type
    FROM dependencies d
    LEFT OUTER JOIN dependencies mv
        ON ( mv.referenced_owner = d.referenced_owner
            AND mv.referenced_name = d.referenced_name
            AND mv.referenced_type = 'MATERIALIZED VIEW' )
    WHERE d.referenced_type <> 'MATERIALIZED VIEW'
        AND d.owner NOT IN ( %s )
        AND d.referenced_owner NOT IN ( %s )
        AND ( d.owner || '.' || d.name <> d.referenced_owner || '.' || d.referenced_name )
        AND ( ( ( d.owner = args.schema_name OR ( args.schema_name IS NULL AND args.table_name IS NULL ) )
                AND ( d.name = args.object_name OR args.object_name IS NULL ) )
            OR ( ( d.referenced_owner = args.schema_name OR ( args.schema_name IS NULL AND args.table_name IS NULL ) )
                    AND ( d.referenced_name = args.object_name OR args.object_name IS NULL ) ) )









SELECT DISTINCT d.owner,
        d.name,
        d.type,
        d.referenced_owner,
        d.referenced_name,
        coalesce ( mv.referenced_type, d.referenced_type ) AS referenced_type
    FROM sys.$table d
    LEFT OUTER JOIN sys.$table mv
        ON ( mv.referenced_owner = d.referenced_owner
            AND mv.referenced_name = d.referenced_name
            AND mv.referenced_type = 'MATERIALIZED VIEW' )
    WHERE d.owner = ? $table_filter
        AND d.referenced_owner NOT IN ( $not_in )
        AND d.referenced_type <> 'MATERIALIZED VIEW'
        AND ( d.owner || '.' || d.name <> d.referenced_owner || '.' || d.referenced_name )


SELECT DISTINCT d.referenced_owner,
        d.referenced_name,
        coalesce ( mv.referenced_type, d.referenced_type ) AS referenced_type,
        d.owner,
        d.name,
        CASE
            WHEN d.type = 'PACKAGE BODY'
            THEN 'PACKAGE'
            WHEN d.type = 'TYPE BODY'
            THEN 'TYPE'
            ELSE d.type
        END AS type
    FROM sys.$dep_table d
    LEFT OUTER JOIN sys.$trig_table t
        ON ( d.owner = t.table_owner
            AND d.name = t.table_name )
    LEFT OUTER JOIN sys.$dep_table mv
        ON ( mv.referenced_owner = d.referenced_owner
            AND mv.referenced_name = d.referenced_name
            AND mv.referenced_type = 'MATERIALIZED VIEW' )
    WHERE d.referenced_owner = ? $table_filter
        AND d.referenced_type <> 'MATERIALIZED VIEW'
        AND ( ( d.type <> 'TRIGGER'
                AND d.owner || '.' || d.name <> d.referenced_owner || '.' || d.referenced_name )
            OR ( d.type = 'TRIGGER'
                AND ( t.table_name <> d.name
                    OR d.referenced_owner <> d.owner ) ) )

*/
