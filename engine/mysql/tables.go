package mysql

import (
	e "github.com/gsiems/go-db-meta/engine/mariadb"
	m "github.com/gsiems/go-db-meta/model"
)

// Tables currently wraps the mariadb.Tables function
func Tables(db *m.DB, schema string) ([]m.Table, error) {
	return e.Tables(db, schema)

}
