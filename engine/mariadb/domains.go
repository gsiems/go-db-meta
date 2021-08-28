package mariadb

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Domains returns an empty set of domains as MariaDB does not appear to support domains
func Domains(db *m.DB, schema string) ([]m.Domain, error) {

	q := ``
	return db.Domains(q, schema)
}
