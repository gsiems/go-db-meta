package ora

import (
	"database/sql"
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

// CurrentCatalog defines the query for obtaining information about the
// currently connected catalog (database) and returns the results of
// executing the query
func CurrentCatalog(db *m.DB) (m.Catalog, error) {

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

func catalogName(db *m.DB) (sql.NullString, error) {
	return db.QSingleString(`
SELECT sys_context ( 'userenv', 'DB_NAME' ) AS catalog_name
    FROM dual`)
}

func defaultCharacterSetName(db *m.DB) (sql.NullString, error) {

	var d sql.NullString
	var nlsLanguage sql.NullString
	var nlsTerritory sql.NullString
	var nlsCharacterset sql.NullString

	q := `
SELECT d.value
    FROM sys.nls_database_parameters d
    WHERE d.parameter = 'NLS_CHARACTERSET'
`

	nlsCharacterset, err = db.QSingleString(q)
	if err != nil {
		return d, err
	}

	q = `
SELECT coalesce ( s.value, d.value ) AS value
    FROM sys.nls_database_parameters d
    LEFT OUTER JOIN sys.nls_session_parameters s
        ON ( d.parameter = s.parameter )
    WHERE d.parameter = 'NLS_TERRITORY'
`

	nlsTerritory, err = db.QSingleString(q)
	if err != nil {
		return d, err
	}

	q = `
SELECT coalesce ( s.value, d.value ) AS value
    FROM sys.nls_database_parameters d
    LEFT OUTER JOIN sys.nls_session_parameters s
        ON ( d.parameter = s.parameter )
    WHERE d.parameter = 'NLS_LANGUAGE'
`

	nlsLanguage, err = db.QSingleString(q)
	if err != nil {
		return d, err
	}

	d.String = fmt.Sprintf("%s_%s.%s", nlsLanguage, nlsTerritory, nlsCharacterset)

	return d, nil

}

func dbmsVersion(db *m.DB) (sql.NullString, error) {
	return db.QSingleString(`
SELECT banner
    FROM v$version
    WHERE lower ( banner ) LIKE '%database%'`)
}
