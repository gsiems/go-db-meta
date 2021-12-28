package model

import (
	"database/sql"
)

// PrimaryKey contains primary key columns for tables
type PrimaryKey struct {
	TableCatalog      sql.NullString `json:"tableCatalog"`
	TableSchema       sql.NullString `json:"tableSchema"`
	TableName         sql.NullString `json:"tableName"`
	ConstraintName    sql.NullString `json:"constraintName"`
	ConstraintColumns sql.NullString `json:"constraintColumns"`
	ConstraintStatus  sql.NullString `json:"constraintStatus"`
	Comment           sql.NullString `json:"comment"`
}

// PrimaryKeys returns a slice of primary keys for the (schemaName, tableName) parameters
func PrimaryKeys(db *sql.DB, q, schemaName, tableName string) ([]PrimaryKey, error) {

	var d []PrimaryKey

	if q == "" {
		return d, nil
	}

	rows, err := db.Query(q, schemaName, tableName)
	if err != nil {
		return d, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for rows.Next() {
		var u PrimaryKey
		err = rows.Scan(&u.TableCatalog,
			&u.TableSchema,
			&u.TableName,
			&u.ConstraintName,
			&u.ConstraintColumns,
			&u.ConstraintStatus,
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
