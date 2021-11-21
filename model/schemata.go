package model

import (
	"database/sql"
)

/*

| Table Name                            | Column Name                       | Position | Matches                                 | Qty |
| ------------------------------------- | --------------------------------- | -------- | --------------------------------------- | --- |
| SCHEMATA                              | CATALOG_NAME                      | 1        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| SCHEMATA                              | SCHEMA_NAME                       | 2        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| SCHEMATA                              | SCHEMA_OWNER                      | 3        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| SCHEMATA                              | DEFAULT_CHARACTER_SET_CATALOG     | 4        | sql2003, pg, mssql, hsqldb              | 4   |
| SCHEMATA                              | DEFAULT_CHARACTER_SET_SCHEMA      | 5        | sql2003, pg, mssql, hsqldb              | 4   |
| SCHEMATA                              | DEFAULT_CHARACTER_SET_NAME        | 6        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| SCHEMATA                              | SQL_PATH                          | 7        | sql2003, pg, mariadb, hsqldb            | 4   |
| SCHEMATA                              | DEFAULT_COLLATION_NAME            |          | mariadb, h2                             | 2   |

*/

// Schema contains details for Schemata
type Schema struct {
	CatalogName                sql.NullString `json:"catalogName"`
	SchemaName                 sql.NullString `json:"schemaName"`
	SchemaOwner                sql.NullString `json:"schemaOwner"`
	DefaultCharacterSetCatalog sql.NullString `json:"defaultCharacterSetCatalog"`
	DefaultCharacterSetSchema  sql.NullString `json:"defaultCharacterSetSchema"`
	DefaultCharacterSetName    sql.NullString `json:"defaultCharacterSetName"`
	Comment                    sql.NullString `json:"comment"`
}

// Schemata returns a slice of Schemas, optionally filtered on the (nclude, xclude) parameters
func (db *m.DB) Schemata(q, nclude, xclude string) ([]Schema, error) {

	var d []Schema

	if q == "" {
		return d, nil
	}

	included := csvToMap(nclude)
	excluded := csvToMap(xclude)

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
		var u Schema
		err = rows.Scan(&u.CatalogName,
			&u.SchemaName,
			&u.SchemaOwner,
			&u.DefaultCharacterSetCatalog,
			&u.DefaultCharacterSetSchema,
			&u.DefaultCharacterSetName,
			&u.Comment,
		)
		if err != nil {
			return d, err
		} else {
			if u.SchemaName.Valid {
				switch {
				case nclude != "":
					_, ok := included[u.SchemaName.String]
					if ok {
						d = append(d, u)
					}
				case xclude != "":
					_, ok := excluded[u.SchemaName.String]
					if !ok {
						d = append(d, u)
					}
				default:
					d = append(d, u)
				}
			}
		}
	}

	return d, err
}
