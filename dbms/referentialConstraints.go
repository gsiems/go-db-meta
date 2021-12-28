package dbms

import (
	"github.com/gsiems/go-db-meta/dbms/mariadb"
	"github.com/gsiems/go-db-meta/dbms/mssql"
	"github.com/gsiems/go-db-meta/dbms/ora"
	"github.com/gsiems/go-db-meta/dbms/pg"
	"github.com/gsiems/go-db-meta/dbms/sqlite"
	m "github.com/gsiems/go-db-meta/model"
)

// ReferentialConstraints returns a slice of Referential Constraints
// for the (schemaName, tableName) parameters
func (db *DBMS) ReferentialConstraints(schemaName, tableName string) ([]m.ReferentialConstraint, error) {

	var d []m.ReferentialConstraint

	switch db.id {
	case PostgreSQL:
		return pg.ReferentialConstraints(db.Connection, schemaName, tableName)
	case SQLite:
		return sqlite.ReferentialConstraints(db.Connection, schemaName, tableName)
	case MariaDB, MySQL:
		return mariadb.ReferentialConstraints(db.Connection, schemaName, tableName)
	case Oracle:
		return ora.ReferentialConstraints(db.Connection, schemaName, tableName)
	case MSSQL:
		return mssql.ReferentialConstraints(db.Connection, schemaName, tableName)
	}

	return d, unsupportedDBErr(db.id)
}
