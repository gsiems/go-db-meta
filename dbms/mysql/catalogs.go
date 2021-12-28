package mariadb

import (
	e "github.com/gsiems/go-db-meta/engine/mariadb"
	m "github.com/gsiems/go-db-meta/model"
)

// CurrentCatalog currently wraps the mariadb.CurrentCatalog function
func CurrentCatalog(db *m.DB) (m.Catalog, error) {
	return e.CurrentCatalog(db)
}
