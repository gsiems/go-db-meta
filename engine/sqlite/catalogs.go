package sqlite

import (
	m "github.com/gsiems/go-db-meta/model"
)

// CurrentCatalog defines the query for obtaining information about the
// currently connected catalog (database) and returns the results of
// executing the query
func CurrentCatalog(db *m.DB) (d m.Catalog, err error) {

	d.CatalogName, err = catalogName(db)
	if err != nil {
		return d, err
	}

	d.DefaultCharacterSetName, err = defaultCharacterSetName(db)
	if err != nil {
		return d, err
	}

	d.DBMSVersion, err = dbmsVersion(db)
	if err != nil {
		return d, err
	}

	return d, err
}

func catalogName (db *m.DB) (sql.NullString, error) {
	return db.QSingleString("SELECT file FROM pragma_database_list WHERE seq = 0")
}

func defaultCharacterSetName (db *m.DB) (sql.NullString, error) {
	return db.QSingleString("SELECT encoding FROM pragma_encoding")
}

func dbmsVersion (db *m.DB) (sql.NullString, error) {
	return db.QSingleString("SELECT 'SQLite ' || sqlite_version ()")
}
