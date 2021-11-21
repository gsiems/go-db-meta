package model

import (
	"database/sql"
)

/*

| Table Name                            | Column Name                       | Position | Matches                                 | Qty |
| ------------------------------------- | --------------------------------- | -------- | --------------------------------------- | --- |
| USER_DEFINED_TYPES                    | USER_DEFINED_TYPE_CATALOG         | 1        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | USER_DEFINED_TYPE_SCHEMA          | 2        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | USER_DEFINED_TYPE_NAME            | 3        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | USER_DEFINED_TYPE_CATEGORY        | 4        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | IS_INSTANTIABLE                   | 5        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | IS_FINAL                          | 6        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | ORDERING_FORM                     | 7        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | ORDERING_CATEGORY                 | 8        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | ORDERING_ROUTINE_CATALOG          | 9        | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | ORDERING_ROUTINE_SCHEMA           | 10       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | ORDERING_ROUTINE_NAME             | 11       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | REFERENCE_TYPE                    | 12       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | DATA_TYPE                         | 13       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | CHARACTER_MAXIMUM_LENGTH          | 14       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | CHARACTER_OCTET_LENGTH            | 15       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | CHARACTER_SET_CATALOG             | 16       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | CHARACTER_SET_SCHEMA              | 17       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | CHARACTER_SET_NAME                | 18       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | COLLATION_CATALOG                 | 19       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | COLLATION_SCHEMA                  | 20       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | COLLATION_NAME                    | 21       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | NUMERIC_PRECISION                 | 22       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | NUMERIC_PRECISION_RADIX           | 23       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | NUMERIC_SCALE                     | 24       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | DATETIME_PRECISION                | 25       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | INTERVAL_TYPE                     | 26       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | INTERVAL_PRECISION                | 27       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | SOURCE_DTD_IDENTIFIER             | 28       | sql2003, pg, hsqldb                     | 3   |
| USER_DEFINED_TYPES                    | REF_DTD_IDENTIFIER                | 29       | sql2003, pg, hsqldb                     | 3   |

*/

// Type contains details for user defined types
type Type struct {
	TypeCatalog sql.NullString `json:"typeCatalog"`
	TypeSchema  sql.NullString `json:"typeSchema"`
	TypeName    sql.NullString `json:"typeName"`
	TypeOwner   sql.NullString `json:"typeOwner"`
	//DataType    sql.NullString `json:"dataType"`
	Comment sql.NullString `json:"comment"`
}

// Types returns a slice of Types for the (schema) parameter
func (db *m.DB) Types(q, tableSchema string) ([]Type, error) {

	var d []Type

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
		var u Type
		err = rows.Scan(&u.TypeCatalog,
			&u.TypeSchema,
			&u.TypeName,
			&u.TypeOwner,
			//&u.DataType,
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
