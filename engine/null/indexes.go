package null

import (
	m "github.com/gsiems/go-db-meta/model"
)

// Indexes defines the query for obtaining a list of indexes
// for the (tableSchema, tableName) parameters and returns the results
// of executing the query
func Indexes(db *m.DB, tableSchema, tableName string) ([]m.Index, error) {

	q := ``
	return db.Indexes(q, tableSchema, tableName)
}

/*
type Index struct {
	IndexCatalog sql.NullString `json:"indexCatalog"`
	IndexSchema  sql.NullString `json:"indexSchema"`
	IndexName    sql.NullString `json:"indexName"`
	IndexType    sql.NullString `json:"indexType"`
	IndexColumns sql.NullString `json:"indexColumns"`
	TableCatalog sql.NullString `json:"tableCatalog"`
	TableSchema  sql.NullString `json:"tableSchema"`
	TableName    sql.NullString `json:"tableName"`
	IsUnique     sql.NullString `json:"isUnique"`
	Comment      sql.NullString `json:"comment"`
}
*/
