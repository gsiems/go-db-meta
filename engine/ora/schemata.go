package ora

import (
	"database/sql"
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

// Schemata defines the query for obtaining a list of schemata
// as filtered by the (nclude, xclude) parameters and returns the
// results of executing the query
func Schemata(db *m.DB, nclude, xclude string) ([]m.Schema, error) {

	// TODO: Is there a way to query Oracle for the list of system accounts?

	q := `
WITH cs AS (
    SELECT value$ AS cs_name
      FROM props$
      WHERE name = 'NLS_CHARACTERSET'
)
SELECT sys_context ( 'userenv', 'DB_NAME' ) AS catalog_name,
        usr.username AS schema_name,
        usr.username AS owner,
        NULL AS DefaultCharacterSetCatalog,
        NULL AS DefaultCharacterSetSchema,
        cs.cs_name AS DefaultCharacterSetName,
        -- sys_context ( 'userenv', 'NLS_SORT' ) AS default_collation_name,
        NULL AS comment
    FROM dba_users usr
    CROSS JOIN cs
    WHERE usr.username NOT IN ( %s )
        AND usr.username NOT LIKE '%$%'
        AND EXISTS (
            SELECT 1
                FROM dba_objects obj
                WHERE obj.owner = usr.username
                    AND obj.object_type IN ( 'TABLE', 'VIEW', 'MATERIALIZED VIEW' ) )

`
	q2 := fmt.Sprintf(q, systemTables, "%$%")
	return db.Tables(q2, schema)

	d, err := db.Schemata(q, nclude, xclude)
	if err != nil {
		return d, err
	}

	// loop the results and populate comments
	commented, err := commentedSchemas(db)

	for i := range d {
		schema := d[i].SchemaName.String
		_, ok := commented[schema]
		if ok {
			d[i].Comment, err = schemaComment(db, schema)
		}
	}

	return d, err
}

func commentedSchemas(db *m.DB) (map[string]bool, error) {

	d := make(map[string]bool)

	q := `
    SELECT owner
        FROM sys.dba_objects
        WHERE object_type = 'VIEW'
            AND object_name = 'SCHEMA_COMMENT'
`

	rows, err := db.Query(q, tableSchema)
	if err != nil {
		return d, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for rows.Next() {
		var u sqlNullString
		err = rows.Scan(&u)
		if err != nil {
			return d, err
		} else {
			d[u.String] = true
		}
	}

	return d, err
}

func schemaComment(db *m.DB, schema string) (sql.NullString, error) {
	return db.QSingleString(fmt.Sprintf(`SELECT schema_comment FROM %q."SCHEMA_COMMENT"`, schema))
}
