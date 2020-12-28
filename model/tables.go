package model

import (
	"database/sql"
	"errors"
)

// Table contains details for tables and views
type Table struct {
	TableCatalog sql.NullString `json:"tableCatalog"`
	TableSchema  sql.NullString `json:"tableSchema"`
	TableName    sql.NullString `json:"tableName"`
	TableOwner   sql.NullString `json:"tableOwner"`
	TableType    sql.NullString `json:"tableType"`
	//ColumnCount    sql.NullInt64  `json:"columnCount"` // ??
	RowCount       sql.NullInt64  `json:"rowCount"`
	Comment        sql.NullString `json:"comment"`
	ViewDefinition sql.NullString `json:"viewDefinition"`
}

// Tables returns a slice of Tables for the (schema) parameter
func (db *m.DB) Tables(q, tableSchema string) ([]Table, error) {

	var d []Table

	if q == "" {
		return d, errors.New("No query provided to Tables")
	}

	rows, err := db.Query(q, tableSchema)
	if err != nil {
		return d, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for rows.Next() {
		var u Table
		err = rows.Scan(&u.TableCatalog,
			&u.TableSchema,
			&u.TableName,
			&u.TableOwner,
			&u.TableType,
			&u.RowCount,
			&u.Comment,
			&u.ViewDefinition,
		)
		if err != nil {
			return d, err
		} else {
			d = append(d, u)
		}
	}

	return d, err
}
