package model

import (
	"database/sql"
	"errors"
)

type Catalog struct {
	CatalogName sql.NullString `json:"catalogName"`
	// Type {PostgreSQL, Oracle, MariaDB, etc.}?
	// Version?
	// CharacterSetName?
	Comment sql.NullString `json:"comment"`
}

// CurrentCatalog returns the current catalog
func (db *m.DB) CurrentCatalog(q) (Catalog, error) {

	var d Catalog

	if q == "" {
		return d, errors.New("No query provided to CurrentCatalog")
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
			&d.Comment,
		)
	}

	return d, err
}
