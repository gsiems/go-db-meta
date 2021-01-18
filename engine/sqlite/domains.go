package sqlite

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Domains returns an empty set of domains as SQLite does not support domains
func Domains(db *m.DB, schema string) ([]m.Domain, error) {
	var d []m.Domain
	return d, nil
}
