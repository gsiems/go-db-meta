package model

import (
	"database/sql"
)

/*

| Table Name                            | Column Name                       | Position | Matches                                 | Qty |
| ------------------------------------- | --------------------------------- | -------- | --------------------------------------- | --- |
| DOMAINS                               | DOMAIN_CATALOG                    | 1        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| DOMAINS                               | DOMAIN_SCHEMA                     | 2        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| DOMAINS                               | DOMAIN_NAME                       | 3        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| DOMAINS                               | DATA_TYPE                         | 4        | sql2003, pg, mssql, hsqldb, h2          | 5   |
| DOMAINS                               | CHARACTER_MAXIMUM_LENGTH          | 5        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | CHARACTER_OCTET_LENGTH            | 6        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | CHARACTER_SET_CATALOG             | 7        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | CHARACTER_SET_SCHEMA              | 8        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | CHARACTER_SET_NAME                | 9        | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | COLLATION_CATALOG                 | 10       | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | COLLATION_SCHEMA                  | 11       | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | COLLATION_NAME                    | 12       | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | NUMERIC_PRECISION                 | 13       | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | NUMERIC_PRECISION_RADIX           | 14       | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | NUMERIC_SCALE                     | 15       | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | DATETIME_PRECISION                | 16       | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | INTERVAL_TYPE                     | 17       | sql2003, pg, hsqldb                     | 3   |
| DOMAINS                               | INTERVAL_PRECISION                | 18       | sql2003, pg, hsqldb                     | 3   |
| DOMAINS                               | DOMAIN_DEFAULT                    | 19       | sql2003, pg, mssql, hsqldb              | 4   |
| DOMAINS                               | MAXIMUM_CARDINALITY               | 20       | sql2003, pg, hsqldb                     | 3   |
| DOMAINS                               | DTD_IDENTIFIER                    | 21       | sql2003, pg, hsqldb                     | 3   |

*/

// Domain contains details for (user defined) domains
type Domain struct {
	DomainCatalog sql.NullString `json:"domainCatalog"`
	DomainSchema  sql.NullString `json:"domainSchema"`
	DomainName    sql.NullString `json:"domainName"`
	DomainOwner   sql.NullString `json:"domainOwner"`
	DataType      sql.NullString `json:"dataType"`
	DomainDefault sql.NullString `json:"domainDefault"`
	Comment       sql.NullString `json:"comment"`
}

// Domains returns a slice of Domains for the (schema) parameter
func (db *DB) Domains(q, schema string) ([]Domain, error) {

	var d []Domain

	if q == "" {
		return d, nil
	}

	rows, err := db.Query(q, schema)
	if err != nil {
		return d, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	for rows.Next() {
		var u Domain
		err = rows.Scan(&u.DomainCatalog,
			&u.DomainSchema,
			&u.DomainName,
			&u.DomainOwner,
			&u.DataType,
			&u.DomainDefault,
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
