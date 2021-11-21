package model

import (
	"database/sql"
)

/*

| Table Name                            | Column Name                       | Position | Matches                                 | Qty |
| ------------------------------------- | --------------------------------- | -------- | --------------------------------------- | --- |
| REFERENTIAL_CONSTRAINTS               | CONSTRAINT_CATALOG                | 1        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | CONSTRAINT_SCHEMA                 | 2        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | CONSTRAINT_NAME                   | 3        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | UNIQUE_CONSTRAINT_CATALOG         | 4        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | UNIQUE_CONSTRAINT_SCHEMA          | 5        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | UNIQUE_CONSTRAINT_NAME            | 6        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | MATCH_OPTION                      | 7        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | UPDATE_RULE                       | 8        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| REFERENTIAL_CONSTRAINTS               | DELETE_RULE                       | 9        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |

| CONSTRAINT_COLUMN_USAGE               | TABLE_CATALOG                     | 1        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_COLUMN_USAGE               | TABLE_SCHEMA                      | 2        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_COLUMN_USAGE               | TABLE_NAME                        | 3        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_COLUMN_USAGE               | COLUMN_NAME                       | 4        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_COLUMN_USAGE               | CONSTRAINT_CATALOG                | 5        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_COLUMN_USAGE               | CONSTRAINT_SCHEMA                 | 6        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_COLUMN_USAGE               | CONSTRAINT_NAME                   | 7        | sql2003, pg, mssql, hsqldb              | 4   |

| CONSTRAINT_TABLE_USAGE                | TABLE_CATALOG                     | 1        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_TABLE_USAGE                | TABLE_SCHEMA                      | 2        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_TABLE_USAGE                | TABLE_NAME                        | 3        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_TABLE_USAGE                | CONSTRAINT_CATALOG                | 4        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_TABLE_USAGE                | CONSTRAINT_SCHEMA                 | 5        | sql2003, pg, mssql, hsqldb              | 4   |
| CONSTRAINT_TABLE_USAGE                | CONSTRAINT_NAME                   | 6        | sql2003, pg, mssql, hsqldb              | 4   |

| KEY_COLUMN_USAGE                      | CONSTRAINT_CATALOG                | 1        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | CONSTRAINT_SCHEMA                 | 2        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | CONSTRAINT_NAME                   | 3        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | TABLE_CATALOG                     | 4        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | TABLE_SCHEMA                      | 5        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | TABLE_NAME                        | 6        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | COLUMN_NAME                       | 7        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | ORDINAL_POSITION                  | 8        | sql2003, pg, mssql, mariadb, hsqldb, h2 | 6   |
| KEY_COLUMN_USAGE                      | POSITION_IN_UNIQUE_CONSTRAINT     | 9        | sql2003, pg, mariadb, hsqldb, h2        | 5   |
| KEY_COLUMN_USAGE                      | REFERENCED_COLUMN_NAME            | (null)   | mariadb                                 | 1   |
| KEY_COLUMN_USAGE                      | REFERENCED_TABLE_NAME             | (null)   | mariadb                                 | 1   |
| KEY_COLUMN_USAGE                      | REFERENCED_TABLE_SCHEMA           | (null)   | mariadb                                 | 1   |

*/

// ReferentialConstraint contains details for referential constraints
type ReferentialConstraint struct {
	TableCatalog      sql.NullString `json:"tableCatalog"`
	TableSchema       sql.NullString `json:"tableSchema"`
	TableName         sql.NullString `json:"tableName"`
	TableColumns      sql.NullString `json:"tableColumns"`
	ConstraintName    sql.NullString `json:"constraintName"`
	RefTableCatalog   sql.NullString `json:"refTableCatalog"`
	RefTableSchema    sql.NullString `json:"refTableSchema"`
	RefTableName      sql.NullString `json:"refTableName"`
	RefTableColumns   sql.NullString `json:"refTableColumns"`
	RefConstraintName sql.NullString `json:"refConstraintName"`
	MatchOption       sql.NullString `json:"matchOption"`
	UpdateRule        sql.NullString `json:"updateRule"`
	DeleteRule        sql.NullString `json:"deleteRule"`
	IsEnforced        sql.NullString `json:"isEnforced"`
	//is_deferrable
	//initially_deferred
	Comment sql.NullString `json:"comment"`
}

// ReferentialConstraints returns a slice of Referential Constraints
// for the (tableSchema, tableName) parameters
func (db *m.DB) ReferentialConstraints(q, tableSchema, tableName string) ([]ReferentialConstraint, error) {

	var d []ReferentialConstraint

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
		var u ReferentialConstraint
		err = rows.Scan(&u.TableCatalog,
			&u.TableSchema,
			&u.TableName,
			&u.TableColumns,
			&u.ConstraintName,
			&u.RefTableCatalog,
			&u.RefTableSchema,
			&u.RefTableName,
			&u.RefTableColumns,
			&u.RefConstraintName,
			&u.MatchOption,
			&u.UpdateRule,
			&u.DeleteRule,
			&u.IsEnforced,
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
