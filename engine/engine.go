package engine
/*
import (
	"database/sql"

	m "github.com/gsiems/go-db-meta/model"
)

type Engine interface {
	Name () (string)
	OpenDB(c *ConnectInfo) (*DB, error)
	BindConnection(db *sql.DB) (*DB, error)
	//CloseDB()
	CurrentCatalog(q string) (m.Catalog, error)
	CheckConstraints(q, tableSchema, tableName string) ([]m.CheckConstraint, error)
	Columns(q, tableSchema, tableName string) ([]m.Column, error)
	//DatabaseList(q string) (s []string, err error)
	//DatabaseInfo(q string) (d Database, err error)
	Dependencies(q, objectSchema, objectName string) ([]m.Dependency, error)
	Domains(q, schema string) ([]m.Domain, error)
	Indexes(q, tableSchema, tableName string) ([]m.Index, error)
	PrimaryKeys(q, tableSchema, tableName string) ([]m.PrimaryKey, error)
	ReferentialConstraints(q, tableSchema, tableName string) ([]m.ReferentialConstraint, error)
	Schemata(q, nclude, xclude string) ([]m.Schema, error)
	Tables(q, tableSchema string) ([]m.Table, error)
	Types(q, tableSchema string) ([]m.Type, error)
	UniqueConstraints(q, tableSchema, tableName string) ([]m.UniqueConstraint, error)
}

type ConnectInfo struct {
	Host     string
	Port     string
	DbName   string
	Username string
	Password string
	File     string
}

type DB struct {
	Connection *m.DB
}
*/
