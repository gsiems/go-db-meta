package model

import (
	"database/sql"
)

type Catalog struct {
	CatalogName             sql.NullString `json:"catalogName"`
	CatalogOwner            sql.NullString `json:"catalogOwner"`
	DefaultCharacterSetName sql.NullString `json:"defaultCharacterSetName"`
	DBMSVersion             sql.NullString `json:"dbmsVersion"`
	Comment                 sql.NullString `json:"comment"`
}

// CurrentCatalog returns the current catalog
func CurrentCatalog(db *sql.DB, q string) (Catalog, error) {

	var d Catalog

	if q == "" {
		return d, nil
	}

	rows, err := db.Query(q)
	if err != nil {
		return d, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	if rows.Next() {
		err = rows.Scan(&d.CatalogName,
			&d.CatalogOwner,
			&d.DefaultCharacterSetName,
			&d.DBMSVersion,
			&d.Comment,
		)
	}

	return d, err
}
