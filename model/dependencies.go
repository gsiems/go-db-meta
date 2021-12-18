package model

import (
	"database/sql"
)

// Dependency contains details for Dependencies
type Dependency struct {
	ObjectCatalog    sql.NullString `json:"objectCatalog"`
	ObjectSchema     sql.NullString `json:"objectSchema"`
	ObjectName       sql.NullString `json:"objectName"`
	ObjectOwner      sql.NullString `json:"objectOwner"`
	ObjectType       sql.NullString `json:"objectType"`
	DepObjectCatalog sql.NullString `json:"depObjectCatalog"`
	DepObjectSchema  sql.NullString `json:"depObjectSchema"`
	DepObjectName    sql.NullString `json:"depObjectName"`
	DepObjectOwner   sql.NullString `json:"depObjectOwner"`
	DepObjectType    sql.NullString `json:"depObjectType"`
}

// Dependencies returns a slice of Dependecies for the
// (objectSchema, objectName) parameters
func (db *DB) Dependencies(q, objectSchema, objectName string) ([]Dependency, error) {

	var d []Dependency

	if q == "" {
		return d, nil
	}

	rows, err := db.Query(q, objectSchema, objectName)
	if err != nil {
		return d, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for rows.Next() {
		var u Dependency
		err = rows.Scan(&u.ObjectCatalog,
			&u.ObjectSchema,
			&u.ObjectName,
			&u.ObjectOwner,
			&u.ObjectType,
			&u.DepObjectCatalog,
			&u.DepObjectSchema,
			&u.DepObjectName,
			&u.DepObjectOwner,
			&u.DepObjectType,
		)
		if err != nil {
			return d, err
		} else {
			d = append(d, u)
		}
	}

	return d, err
}
