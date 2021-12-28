package dbms

import "fmt"

func unsupportedDBErr(id int) error {
	return fmt.Errorf("unknown, or unsupported, database ID (%d) specified", id)
}
