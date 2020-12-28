package MODEL

import (
	"database/sql"
	"errors"
)

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
		return d, errors.New("No query provided to Schemata")
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
