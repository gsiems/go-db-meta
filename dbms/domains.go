package dbms

import (
	"github.com/gsiems/go-db-meta/dbms/mssql"
	"github.com/gsiems/go-db-meta/dbms/pg"
	m "github.com/gsiems/go-db-meta/model"
)

// Domains returns a slice of Domains for the (schemaName) parameter
func (db *DBMS) Domains(schemaName string) ([]m.Domain, error) {

	var d []m.Domain

	switch db.id {
	case PostgreSQL:
		return pg.Domains(db.Connection, schemaName)
	case SQLite, MariaDB, MySQL, Oracle: // don't support
		return d, nil
	case MSSQL:
		return mssql.Domains(db.Connection, schemaName)
	}

	return d, unsupportedDBErr(db.id)
}
