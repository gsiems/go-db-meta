package model

import (
	"database/sql"
)

// Index contains details for Indexes
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

// Indexes returns a slice of Indexes for the (tableSchema, tableName) parameters
func (db *m.DB) Indexes(q, tableSchema, tableName string) ([]Index, error) {

	var d []Index

	if q == "" {
		return d, nil
	}

	rows, err := db.Query(q, tableSchema, tableName)
	if err != nil {
		return d, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for rows.Next() {
		var u Index
		var cols string
		err = rows.Scan(&u.IndexCatalog,
			&u.IndexSchema,
			&u.IndexName,
			&u.IndexType,
			&u.IndexColumns,
			&u.TableCatalog,
			&u.TableSchema,
			&u.TableName,
			&u.IsUnique,
			&u.Comment,
		)
		if err != nil {
			return d, err
		}
	}

	return d, err
}
