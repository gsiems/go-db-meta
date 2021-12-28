package mariadb

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// CurrentCatalog defines the query for obtaining information about the
// currently connected catalog (database) and returns the results of
// executing the query
func CurrentCatalog(db *sql.DB) (m.Catalog, error) {

	var r m.Catalog
	r.CatalogName.String = "def"
	r.CatalogName.Valid = true
	var err error

	r.DefaultCharacterSetName, err = defaultCharacterSetName(db)
	if err != nil {
		return r, err
	}

	r.DBMSVersion, err = dbmsVersion(db)
	if err != nil {
		return r, err
	}

	return r, nil
}

func defaultCharacterSetName(db *sql.DB) (sql.NullString, error) {
	return m.QSingleString(db, `
SELECT variable_value
    FROM information_schema.global_variables
    WHERE variable_name = 'CHARACTER_SET_SERVER'`)
}

func dbmsVersion(db *sql.DB) (sql.NullString, error) {
	return m.QSingleString(db, `
SELECT variable_value
    FROM information_schema.global_variables
    WHERE variable_name = 'VERSION'`)
}
