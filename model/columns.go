package model

import (
	"database/sql"
	"errors"
)

// Column contains details for Columns
type Column struct {
	TableCatalog    sql.NullString `json:"tableCatalog"`
	TableSchema     sql.NullString `json:"tableSchema"`
	TableName       sql.NullString `json:"tableName"`
	ColumnName      sql.NullString `json:"columnName"`
	OrdinalPosition sql.NullInt32  `json:"ordinalPosition"`
	ColumnDefault   sql.NullString `json:"columnDefault"`
	IsNullable      sql.NullString `json:"isNullable"`
	DataType        sql.NullString `json:"dataType"`
	DomainCatalog   sql.NullString `json:"domainCatalog"`
	DomainSchema    sql.NullString `json:"domainSchema"`
	DomainName      sql.NullString `json:"domainName"`
	//UdtCatalog      sql.NullString `json:"udtCatalog"`
	//UdtSchema       sql.NullString `json:"udtSchema"`
	//UdtName         sql.NullString `json:"udtName"`
	//IsGenerated          sql.NullString `json:"isGenerated"`          // ??
	//GenerationExpression sql.NullString `json:"generationExpression"` // ??
	Comment sql.NullString `json:"comment"`
}

// Columns returns a slice of Columns for the (tableSchema, tableName) parameters
func (db *m.DB) Columns(q, tableSchema, tableName string) ([]Column, error) {

	var d []Column

	if q == "" {
		return d, errors.New("No query provided to Columns")
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
		var u Column
		err = rows.Scan(&u.TableCatalog,
			&u.TableSchema,
			&u.TableName,
			&u.ColumnName,
			&u.OrdinalPosition,
			&u.ColumnDefault,
			&u.IsNullable,
			&u.DataType,
			&u.DomainCatalog,
			&u.DomainSchema,
			&u.DomainName,
			//&u.UdtCatalog,
			//&u.UdtSchema,
			//&u.UdtName,
			&u.Comment,
		)
		if err != nil {
			return d, err
		} else {
			d = append(d, u)
		}
	}

	return d, err
}
