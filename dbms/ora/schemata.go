package ora

import (
	"database/sql"
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

// Schemata defines the query for obtaining a list of schemata
// as filtered by the (nclude, xclude) parameters and returns the
// results of executing the query
func Schemata(db *sql.DB, nclude, xclude string) ([]m.Schema, error) {

	// TODO: Is there a way to query Oracle for the list of system accounts?

	q := `
WITH cs AS (
    SELECT value$ AS cs_name
      FROM sys.props$
      WHERE name = 'NLS_CHARACTERSET'
)
SELECT sys_context ( 'userenv', 'DB_NAME' ) AS catalog_name,
        usr.username AS schema_name,
        usr.username AS owner,
        NULL AS DefaultCharacterSetCatalog,
        NULL AS DefaultCharacterSetSchema,
        cs.cs_name AS DefaultCharacterSetName,
        -- sys_context ( 'userenv', 'NLS_SORT' ) AS default_collation_name,
        NULL AS comments
    FROM dba_users usr
    CROSS JOIN cs
    WHERE usr.username NOT IN ( ` + systemTables + ` )
        AND EXISTS (
            SELECT 1
                FROM dba_objects obj
                WHERE obj.owner = usr.username
                    AND obj.object_type IN ( 'TABLE', 'VIEW', 'MATERIALIZED VIEW' ) )
`

	d, err := m.Schemata(db, q, nclude, xclude)
	if err != nil {
		return d, err
	}

	// loop the results and populate comments
	commented, err := commentedSchemas(db)

	for i := range d {
		schemaName := d[i].SchemaName.String
		_, ok := commented[schemaName]
		if ok {
			d[i].Comment, err = schemaComment(db, schemaName)
		}
	}

	return d, err
}

func commentedSchemas(db *sql.DB) (map[string]bool, error) {

	d := make(map[string]bool)

	q := `
    SELECT owner
        FROM all_objects
        WHERE object_type = 'VIEW'
            AND object_name = 'SCHEMA_COMMENT'
`

	rows, err := db.Query(q)
	if err != nil {
		return d, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for rows.Next() {
		var u sql.NullString
		err = rows.Scan(&u)
		if err != nil {
			return d, err
		} else {
			d[u.String] = true
		}
	}

	return d, err
}

func schemaComment(db *sql.DB, schemaName string) (sql.NullString, error) {
	return m.QSingleString(db, fmt.Sprintf(`SELECT schema_comment FROM %q."SCHEMA_COMMENT"`, schemaName))
}
