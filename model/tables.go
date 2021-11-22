package model

import (
	"database/sql"
)

/*

| Table Name                            | Column Name                       | Position | Matches                                 | Qty |
| ------------------------------------- | --------------------------------- | -------- | --------------------------------------- | --- |
| TABLES                                | TABLE_CATALOG                     | 1        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLES                                | TABLE_SCHEMA                      | 2        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLES                                | TABLE_NAME                        | 3        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLES                                | TABLE_TYPE                        | 4        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| TABLES                                | SELF_REFERENCING_COLUMN_NAME      | 5        | sql2003, pg, hsqldb                     | 3   |
| TABLES                                | REFERENCE_GENERATION              | 6        | sql2003, pg, hsqldb                     | 3   |
| TABLES                                | USER_DEFINED_TYPE_CATALOG         | 7        | sql2003, pg, hsqldb                     | 3   |
| TABLES                                | USER_DEFINED_TYPE_SCHEMA          | 8        | sql2003, pg, hsqldb                     | 3   |
| TABLES                                | USER_DEFINED_TYPE_NAME            | 9        | sql2003, pg, hsqldb                     | 3   |
| TABLES                                | IS_INSERTABLE_INTO                | 10       | sql2003, pg, hsqldb                     | 3   |
| TABLES                                | IS_TYPED                          | 11       | sql2003, pg, hsqldb                     | 3   |
| TABLES                                | COMMIT_ACTION                     | 12       | sql2003, pg, hsqldb                     | 3   |
| VIEWS                                 | TABLE_CATALOG                     | 1        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| VIEWS                                 | TABLE_SCHEMA                      | 2        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| VIEWS                                 | TABLE_NAME                        | 3        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| VIEWS                                 | VIEW_DEFINITION                   | 4        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| VIEWS                                 | CHECK_OPTION                      | 5        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| VIEWS                                 | IS_UPDATABLE                      | 6        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| VIEWS                                 | INSERTABLE_INTO                   | 7        | sql2003, hsqldb                         | 2   |
| VIEWS                                 | IS_TRIGGER_DELETABLE              |          | pg, hsqldb                              | 2   |
| VIEWS                                 | IS_TRIGGER_INSERTABLE_INTO        |          | pg, hsqldb                              | 2   |
| VIEWS                                 | IS_TRIGGER_UPDATABLE              |          | pg, hsqldb                              | 2   |

*/

// Table contains details for tables and views
type Table struct {
	TableCatalog   sql.NullString `json:"tableCatalog"`
	TableSchema    sql.NullString `json:"tableSchema"`
	TableName      sql.NullString `json:"tableName"`
	TableOwner     sql.NullString `json:"tableOwner"`
	TableType      sql.NullString `json:"tableType"`
	RowCount       sql.NullInt64  `json:"rowCount"`
	Comment        sql.NullString `json:"comment"`
	ViewDefinition sql.NullString `json:"viewDefinition"`
}

// Tables returns a slice of Tables for the (schema) parameter
func (db *DB) Tables(q, tableSchema string) ([]Table, error) {

	var d []Table

	if q == "" {
		return d, nil
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
