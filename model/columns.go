package model

import (
	"database/sql"
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
	Comment         sql.NullString `json:"comment"`
}

// Columns returns a slice of Columns for the (tableSchema, tableName) parameters
func Columns(db *sql.DB, q, tableSchema, tableName string) ([]Column, error) {

	var d []Column

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
		var u Column
		err = rows.Scan(&u.TableCatalog,
			&u.TableSchema,
			&u.TableName,
			&u.ColumnName,
			&u.OrdinalPosition,
			&u.DataType,
			&u.IsNullable,
			&u.ColumnDefault,
			&u.DomainCatalog,
			&u.DomainSchema,
			&u.DomainName,
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

/*

| Table Name                            | Column Name                       | Position | Matches                                 | Qty |
| ------------------------------------- | --------------------------------- | -------- | --------------------------------------- | --- |
| COLUMNS                               | TABLE_CATALOG                     | 1        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | TABLE_SCHEMA                      | 2        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | TABLE_NAME                        | 3        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | COLUMN_NAME                       | 4        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | ORDINAL_POSITION                  | 5        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | COLUMN_DEFAULT                    | 6        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | IS_NULLABLE                       | 7        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | DATA_TYPE                         | 8        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | CHARACTER_MAXIMUM_LENGTH          | 9        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | CHARACTER_OCTET_LENGTH            | 10       | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | NUMERIC_PRECISION                 | 11       | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | NUMERIC_PRECISION_RADIX           | 12       | sql2003, pg, mssql, hsqldb, h2          | 5   |
| COLUMNS                               | NUMERIC_SCALE                     | 13       | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | DATETIME_PRECISION                | 14       | sql2003, pg, mssql, mariadb, hsqldb     | 5   |
| COLUMNS                               | INTERVAL_TYPE                     | 15       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | INTERVAL_PRECISION                | 16       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | CHARACTER_SET_CATALOG             | 17       | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMNS                               | CHARACTER_SET_SCHEMA              | 18       | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMNS                               | CHARACTER_SET_NAME                | 19       | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | COLLATION_CATALOG                 | 20       | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMNS                               | COLLATION_SCHEMA                  | 21       | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMNS                               | COLLATION_NAME                    | 22       | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| COLUMNS                               | DOMAIN_CATALOG                    | 23       | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMNS                               | DOMAIN_SCHEMA                     | 24       | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMNS                               | DOMAIN_NAME                       | 25       | sql2003, pg, mssql, hsqldb              | 4   |
| COLUMNS                               | UDT_CATALOG                       | 26       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | UDT_SCHEMA                        | 27       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | UDT_NAME                          | 28       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | SCOPE_CATALOG                     | 29       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | SCOPE_SCHEMA                      | 30       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | SCOPE_NAME                        | 31       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | MAXIMUM_CARDINALITY               | 32       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | DTD_IDENTIFIER                    | 33       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IS_SELF_REFERENCING               | 34       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IS_IDENTITY                       | 35       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IDENTITY_GENERATION               | 36       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IDENTITY_START                    | 37       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IDENTITY_INCREMENT                | 38       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IDENTITY_MAXIMUM                  | 39       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IDENTITY_MINIMUM                  | 40       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IDENTITY_CYCLE                    | 41       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | IS_GENERATED                      | 42       | sql2003, pg, mariadb, hsqldb            | 4   |
| COLUMNS                               | GENERATION_EXPRESSION             | 43       | sql2003, pg, mariadb, hsqldb            | 4   |
| COLUMNS                               | IS_UPDATABLE                      | 44       | sql2003, pg, hsqldb                     | 3   |
| COLUMNS                               | COLUMN_TYPE                       |          | mariadb, h2                             | 2   |

*/
