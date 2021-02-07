package mariadb

import (
	m "github.com/gsiems/go-db-meta/model"
)

// CurrentCatalog defines the query for obtaining information about the
// currently connected catalog (database) and returns the results of
// executing the query
func CurrentCatalog(db *m.DB) (m.Catalog, error) {

	var d m.Catalog
	d.CatalogName = "def"

	d.DefaultCharacterSetName, err = defaultCharacterSetName(db)
	if err != nil {
		return d, err
	}

	d.DBMSVersion, err = dbmsVersion(db)
	if err != nil {
		return d, err
	}

	return d, nil
}

func defaultCharacterSetName(db *m.DB) (sql.NullString, error) {
	return db.QSingleString(`
SELECT variable_value
    FROM information_schema.global_variables
    WHERE variable_name = 'CHARACTER_SET_SERVER'`)
}

func dbmsVersion(db *m.DB) (sql.NullString, error) {
	return db.QSingleString(`
SELECT variable_value
    FROM information_schema.global_variables
    WHERE variable_name = 'VERSION'`)
}
