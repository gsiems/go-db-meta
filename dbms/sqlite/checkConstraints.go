package sqlite

import (
	"database/sql"
	"strings"

	m "github.com/gsiems/go-db-meta/model"
)

// CheckConstraints returns an empty set it is unclear how to extract
// check constraints from SQLite other than parsing the sqlite_master.sql
// column
func CheckConstraints(db *sql.DB, schemaName, tableName string) ([]m.CheckConstraint, error) {

	var r []m.CheckConstraint

	catName, err := catalogName(db)
	if err != nil {
		return r, err
	}

	q := `
SELECT m.name AS table_name,
        m.sql
    FROM sqlite_master m
    CROSS JOIN (
        SELECT coalesce ( $1, '' ) AS table_name
        ) AS args
    WHERE m.type = 'table'
        AND substr ( m.tbl_name, 1, 7 ) <>  'sqlite'
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
				if strings.Contains(tv, " CHECK ") {

					u := m.CheckConstraint{
						TableCatalog: catName,
						TableName:    name,
					}

					u.TableSchema.String = schemaName
					u.TableSchema.Valid = true

					pts := strings.SplitN(tv, "(", 2)

					u.ConstraintName.String = strings.Trim(strings.TrimPrefix(pts[0], "CONSTRAINT"), " ")
					u.ConstraintName.Valid = true
					u.CheckClause.String = strings.TrimSuffix(pts[1], ")")
					u.CheckClause.Valid = true

					r = append(r, u)

				}
			}
		}
	}
	return r, err

}
