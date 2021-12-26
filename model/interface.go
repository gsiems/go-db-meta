package model

import "database/sql"

type metaData interface {
	BindConnection(db *sql.DB) (DB, error)
	//CloseDB()
	CurrentCatalog(q string) (Catalog, error)
	CheckConstraints(q, tableSchema, tableName string) ([]CheckConstraint, error)
	Columns(q, tableSchema, tableName string) ([]Column, error)
	DatabaseList(q string) (s []string, err error)
	DatabaseInfo(q string) (d Database, err error)
	Dependencies(q, objectSchema, objectName string) ([]Dependency, error)
	Domains(q, schema string) ([]Domain, error)
	Indexes(q, tableSchema, tableName string) ([]Index, error)
	PrimaryKeys(q, tableSchema, tableName string) ([]PrimaryKey, error)
	ReferentialConstraints(q, tableSchema, tableName string) ([]ReferentialConstraint, error)
	Schemata(q, nclude, xclude string) ([]Schema, error)
	Tables(q, tableSchema string) ([]Table, error)
	Types(q, tableSchema string) ([]Type, error)
	UniqueConstraints(q, tableSchema, tableName string) ([]UniqueConstraint, error)
}
