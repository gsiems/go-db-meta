package sqlite

import (
	"database/sql"
	"fmt"

	m "github.com/gsiems/go-db-meta/model"
)

// Columns obtains the list of columns
// for the (tableSchema, tableName) parameters and returns the results
func Columns(db *m.DB, tableSchema, tableName string) ([]m.Column, error) {

	var d []m.Column

	q := `
SELECT '%s' AS table_catalog,
        args.table_schema,
        m.name AS table_name,
        cols.name AS column_name,
        cols.cid AS ordinal_position,
        CASE
            WHEN cols."notnull" = 1 THEN 'N'
            ELSE 'Y'
            END AS is_nullable,
        cols."type" AS data_type,
        NULL AS domain_catalog,
        NULL AS domain_schema,
        NULL AS domain_name,
        NULL AS comment
    FROM sqlite_master AS m
    JOIN pragma_table_info ( m.name ) AS cols
    CROSS JOIN (
        SELECT coalesce ( $1, '' ) AS table_schema,
                coalesce ( $2, '' ) AS table_name
        ) AS args
    WHERE m.type IN ( 'table', 'view' )
        AND m.tbl_name NOT LIKE 'sqlite_%'
        AND ( args.table_name = '' OR args.table_name = m.name )
`

	catName, err := catalogName(db)
	if err != nil {
		return d, err
	}

	q2 := fmt.Sprintf(q, catName)
	return db.Columns(q2, tableSchema, tableName)
}

/*
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

	var d []m.Column

	catName, err := catalogName(db)
	if err != nil {
		return d, err
	}

	q := `
SELECT "cid" AS ordinal_position,
        "name" AS column_name,
        "type" AS dataType,
        "notnull" AS not_null,
        dflt_value AS default_value
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

			var u m.Column

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
*/
