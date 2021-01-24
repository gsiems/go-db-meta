package null

import (
	m "github.com/gsiems/go-db-meta/model"
)

// CurrentCatalog defines the query for obtaining information about the
// currently connected catalog (database) and returns the results of
// executing the query
func CurrentCatalog(db *m.DB) (m.Catalog, error) {

	q := ``
	return db.CurrentCatalog(q)
}
