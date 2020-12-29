package pg

import (
	m "github.com/gsiems/go-db-meta/model"
)

// CurrentCatalog defines the query for obtaining information about the
// currently connected catalog (database) and returns the results of
// executing the query
func CurrentCatalog(db *m.DB) (m.Catalog, error) {

	q := `
SELECT d.datname AS catalog_name,
        pg_catalog.pg_get_userbyid ( d.datdba ) AS owner,
        pg_catalog.pg_encoding_to_char ( d.encoding ) AS encoding,
        pg_catalog.version () AS version,
        pg_catalog.shobj_description ( d.oid, 'pg_database' ) AS comments
    FROM pg_catalog.pg_database d
    WHERE d.datname = pg_catalog.current_database ()
`

	return db.CurrentCatalog(q)
}
