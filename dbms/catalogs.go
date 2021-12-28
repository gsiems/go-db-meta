package dbms

import (
	"github.com/gsiems/go-db-meta/dbms/mariadb"
	"github.com/gsiems/go-db-meta/dbms/mssql"
	"github.com/gsiems/go-db-meta/dbms/ora"
	"github.com/gsiems/go-db-meta/dbms/pg"
	"github.com/gsiems/go-db-meta/dbms/sqlite"
	m "github.com/gsiems/go-db-meta/model"
)

// CurrentCatalog returns the current catalog
func (db *DBMS) CurrentCatalog() (m.Catalog, error) {

	var d m.Catalog

	switch db.id {
	case PostgreSQL:
		return pg.CurrentCatalog(db.Connection)
	case SQLite:
		return sqlite.CurrentCatalog(db.Connection)
	case MariaDB, MySQL:
		return mariadb.CurrentCatalog(db.Connection)
	case Oracle:
		return ora.CurrentCatalog(db.Connection)
	case MSSQL:
		return mssql.CurrentCatalog(db.Connection)
	}

	return d, unsupportedDBErr(db.id)
}
