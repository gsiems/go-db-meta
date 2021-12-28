package mysql

import (
	e "github.com/gsiems/go-db-meta/engine/mariadb"
	m "github.com/gsiems/go-db-meta/model"
)

// Domains currently wraps the mariadb.Domains function
func Domains(db *m.DB, schema string) ([]m.Domain, error) {
	return e.Domains(db, schema)
}
