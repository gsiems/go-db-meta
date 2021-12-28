package sqlite

import (
	"database/sql"
	"strings"

	m "github.com/gsiems/go-db-meta/model"
)

// UniqueConstraints defines the query for obtaining a list of unique
// constraints for the (schemaName, tableName) parameters and returns the
// results of executing the query
func UniqueConstraints(db *sql.DB, schemaName, tableName string) ([]m.UniqueConstraint, error) {

	var r []m.UniqueConstraint

	catName, err := catalogName(db)
	if err != nil {
		return r, err
	}

	q := `
SELECT m.name AS table_name,
        m.sql AS query
    FROM sqlite_master m
    CROSS JOIN (
        SELECT coalesce ( $1, '' ) AS table_name
        ) AS args
    WHERE m.type = 'table'
        AND m.tbl_name NOT LIKE 'sqlite_%'
        AND ( args.table_name = '' OR args.table_name = m.name )
`

	rows, err := db.Query(q, tableName)
	if err != nil {
		return r, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for rows.Next() {

		var name sql.NullString
		var query sql.NullString

		err = rows.Scan(&name,
			&query,
		)
		if err != nil {
			return r, err
		}

		for _, v := range strings.Split(query.String, "\n") {

			tv := strings.Trim(strings.TrimRight(v, "\n\r ,"), " ")
			if strings.HasPrefix(tv, "CONSTRAINT ") {
				if strings.Contains(tv, " UNIQUE ") {

					u := m.UniqueConstraint{
						TableCatalog: catName,
						TableName:    name,
					}

					u.TableSchema.String = schemaName
					u.TableSchema.Valid = true

					pts := strings.SplitN(tv, "(", 2)

					u.ConstraintName.String = strings.Trim(strings.TrimPrefix(pts[0], "CONSTRAINT"), " ")
					u.ConstraintName.Valid = true
					u.ConstraintColumns.String = strings.TrimSuffix(pts[1], ")")
					u.ConstraintColumns.Valid = true

					r = append(r, u)

				}
			}
		}
	}
	return r, err
}
