package mssql

import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

// CurrentCatalog defines the query for obtaining information about the
// currently connected catalog (database) and returns the results of
// executing the query
func CurrentCatalog(db *sql.DB) (m.Catalog, error) {

	q := `
SELECT db_name () AS catalog_name,
        NULL AS catalog_owner,
        serverproperty ( 'SqlCharSetName' ) AS character_set_name,
        concat ( 'Microsoft SQL Server ',
            convert ( varchar ( 100 ), serverproperty ( 'Edition' ) ),
            '; ',
            convert ( varchar ( 50 ), serverproperty ( 'ProductVersion' ) ) ) AS dbms_version,
        NULL AS comment
`

	return m.CurrentCatalog(db, q)
}
