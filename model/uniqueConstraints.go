package model

import (
	"database/sql"
)

// UniqueConstraint contains details for Unique Constraints
type UniqueConstraint struct {
	TableCatalog      sql.NullString `json:"tableCatalog"`
	TableSchema       sql.NullString `json:"tableSchema"`
	TableName         sql.NullString `json:"tableName"`
	ConstraintName    sql.NullString `json:"constraintName"`
	ConstraintColumns sql.NullString `json:"constraintColumns"`
	Status            sql.NullString `json:"status"`
	Comment           sql.NullString `json:"comment"`
}

// UniqueConstraints returns a slice of Unique Constraints for the
// (schemaName, tableName) parameters
func UniqueConstraints(db *sql.DB, q, schemaName, tableName string) ([]UniqueConstraint, error) {

	var d []UniqueConstraint

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
		var u UniqueConstraint
		err = rows.Scan(&u.TableCatalog,
			&u.TableSchema,
			&u.TableName,
			&u.ConstraintName,
			&u.ConstraintColumns,
			&u.Status,
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
