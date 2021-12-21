package ora

import (
	"database/sql"
	"fmt"
	"os/user"

	_ "github.com/godror/godror"

	m "github.com/gsiems/go-db-meta/model"
)

// OpenDB opens a database connection and returns a DB reference.
func OpenDB(c *m.ConnectInfo) (*m.DB, error) {

	var osUser string
	usr, err := user.Current()
	if err == nil {
		osUser = usr.Username
	}

	c.Username = m.Coalesce(c.Username, osUser)
	c.Host = m.Coalesce(c.Host, "localhost")
	c.Port = m.Coalesce(c.Port, "1521")

	connStr := fmt.Sprintf("%s/%s@%s", c.Username, c.Password, c.DbName)
	db, err := sql.Open("godror", connStr)

	if err != nil {
		return nil, err
	}
	return &m.DB{db}, db.Ping()
}
