package model

import (
	"strings"
)

// csvToMap splits a comma-separated list into a map
func csvToMap(s string) map[string]int {

	l := make(map[string]int)

	for i, v := range strings.Split(s, ",") {
		l[v] = i
	}
	return l
}

func (db *m.DB) QSingleString(q string) (sql.NullString, error) {

	var v sql.NullString
	rows, err := db.Query(q)
	if err != nil {
		return v, err
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	if rows.Next() {
		err = rows.Scan(&v)
	}
	return v, err
}
