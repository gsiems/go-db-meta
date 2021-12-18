package mysql

import (
	e "github.com/gsiems/go-db-meta/engine/mariadb"
	m "github.com/gsiems/go-db-meta/model"
)

// Dependencies defines the query for obtaining a list of dependencies
// for the (objectSchema, objectName) parameters and returns the results
// of executing the query
func Dependencies(db *m.DB, objectSchema, objectName string) ([]m.Dependency, error) {
	return e.Dependencies(objectSchema, objectName)
}
