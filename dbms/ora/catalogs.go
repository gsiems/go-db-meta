package ora

import (
	"database/sql"
	"fmt"

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
	return m.QSingleString(db, `
SELECT sys_context ( 'userenv', 'DB_NAME' ) AS catalog_name
    FROM dual`)
}

func defaultCharacterSetName(db *sql.DB) (sql.NullString, error) {

	var err error
	var d sql.NullString
	var nlsLanguage sql.NullString
	var nlsTerritory sql.NullString
	var nlsCharacterset sql.NullString

	q := `
SELECT d.value
    FROM nls_database_parameters d
    WHERE d.parameter = 'NLS_CHARACTERSET'
`

	nlsCharacterset, err = m.QSingleString(db, q)
	if err != nil {
		return d, err
	}

	q = `
SELECT coalesce ( s.value, d.value ) AS value
    FROM nls_database_parameters d
    LEFT OUTER JOIN sys.nls_session_parameters s
        ON ( d.parameter = s.parameter )
    WHERE d.parameter = 'NLS_TERRITORY'
`

	nlsTerritory, err = m.QSingleString(db, q)
	if err != nil {
		return d, err
	}

	q = `
SELECT coalesce ( s.value, d.value ) AS value
    FROM nls_database_parameters d
    LEFT OUTER JOIN sys.nls_session_parameters s
        ON ( d.parameter = s.parameter )
    WHERE d.parameter = 'NLS_LANGUAGE'
`

	nlsLanguage, err = m.QSingleString(db, q)
	if err != nil {
		return d, err
	}

	d.String = fmt.Sprintf("%s_%s.%s", nlsLanguage.String, nlsTerritory.String, nlsCharacterset.String)

	return d, nil

}

func dbmsVersion(db *sql.DB) (sql.NullString, error) {
	return m.QSingleString(db, `
SELECT banner
    FROM v$version
    WHERE lower ( banner ) LIKE '%database%'`)
}
