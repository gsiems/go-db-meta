package ora

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Domains returns an empty set of domains as Oracle does not support domains
func Domains(db *m.DB, tableSchema, tableName string) ([]m.Domain, error) {
	var d []m.Domain
	return d, nil
}
