package sqlite

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// CurrentCatalog defines the query for obtaining information about the
// currently connected catalog (database) and returns the results of
// executing the query
func CurrentCatalog(db *sql.DB) (m.Catalog, error) {

	var r m.Catalog
	var err error

	r.CatalogName, err = catalogName(db)
	if err != nil {
		return r, err
	}

	r.DefaultCharacterSetName, err = defaultCharacterSetName(db)
	if err != nil {
		return r, err
	}

	r.DBMSVersion, err = dbmsVersion(db)
	if err != nil {
		return r, err
	}

	return r, err
}

func catalogName(db *sql.DB) (sql.NullString, error) {
	return m.QSingleString(db, "SELECT file FROM pragma_database_list WHERE seq = 0")
}

func defaultCharacterSetName(db *sql.DB) (sql.NullString, error) {
	return m.QSingleString(db, "SELECT encoding FROM pragma_encoding")
}

func dbmsVersion(db *sql.DB) (sql.NullString, error) {
	return m.QSingleString(db, "SELECT 'SQLite ' || sqlite_version ()")
}
