package sqlite

import (
	"database/sql"
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

type pto struct {
	cid        sql.NullInt32
	colName    sql.NullString
	colType    sql.NullString
	notnull    sql.NullString
	colDefault sql.NullString
}

// Columns obtains the list of columns
// for the (tableSchema, tableName) parameters and returns the results
func Columns(db *m.DB, tableSchema, tableName string) ([]m.Column, error) {

	switch tableName {
	case "":
		var d []m.Column

		tables, err := Tables(db, tableSchema)
		if err != nil {
			return d, err
		}

		for _, v := range tables {
			d2, cerr := tableColumns(db, tableSchema, v.TableName)
			if cerr != nil {
				return d, cerr
			}
			d = append(d, d2)
		}

		return d, err
	}

	d, err := tableColumns(db, tableSchema, tableName)

	return d, err
}

func tableColumns(db *m.DB, tableSchema, tableName string) ([]m.Column, error) {

	var d []Column

	catName, err := catalogName(db)
	if err != nil {
		return d, err
	}

	q := `
SELECT "cid" AS ordinal_position,
        "name" AS column_name,
        "type" AS dataType,
        "notnull" AS not_null
        dflt_value AS default_value,
    FROM pragma_table_info('%s')
    ORDER BY cid
`

	rows, err := db.Query(fmt.Sprintf(q, tableName))
	if err != nil {
		return d, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for rows.Next() {
		var u2 pto
		err = rows.Scan(&u2.cid,
			&u2.colName,
			&u2.colType,
			&u2.notnull,
			&u2.colDefault,
		)
		if err != nil {
			return d, err
		} else {

			var u Column

			u.TableCatalog = catName
			u.TableSchema = schemaName
			u.TableName = tableName
			u.ColumnName = u2.colName
			u.OrdinalPosition = u2.cid
			u.ColumnDefault = u2.colDefault
			switch u2.notnull {
			case 1:
				u.IsNullable = "N"
			default:
				u.IsNullable = "Y"
			}
			u.DataType = u2.colType

			d = append(d, u)
		}
	}

	return d, err
}
