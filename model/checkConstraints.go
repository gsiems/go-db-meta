package model

import (
	"database/sql"
)

// CheckConstraint contains details for referential constraints
type CheckConstraint struct {
	TableCatalog   sql.NullString `json:"tableCatalog"`
	TableSchema    sql.NullString `json:"tableSchema"`
	TableName      sql.NullString `json:"tableName"`
	ConstraintName sql.NullString `json:"constraintName"`
	CheckClause    sql.NullString `json:"checkClause"`
	Status         sql.NullString `json:"status"`
	Comment        sql.NullString `json:"comment"`
}

// CheckConstraints returns a slice of Check Constraints for the
// (tableSchema, tableName) parameters
func CheckConstraints(db *sql.DB, q, tableSchema, tableName string) ([]CheckConstraint, error) {

	var d []CheckConstraint

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
		var u CheckConstraint
		err = rows.Scan(&u.TableCatalog,
			&u.TableSchema,
			&u.TableName,
			&u.ConstraintName,
			&u.CheckClause,
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
