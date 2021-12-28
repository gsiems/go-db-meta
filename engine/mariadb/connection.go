package mariadb

import (
	"database/sql"
	"os/user"

	//e "github.com/gsiems/go-db-meta/engine"
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
	c.Port = m.Coalesce(c.Port, "3306")

	connStr := "" // fmt.Sprintf("%s/%s@%s", c.Username, c.Password, c.DbName)
	db, err := sql.Open("todo", connStr)

	if err != nil {
		return nil, err
	}
	return &m.DB{db}, db.Ping()
}

// BindConnection binds a database/sql connection to the model. This
// should allow the model to be database driver agnostic
func BindConnection(db *sql.DB) (m.DB, error) {
	return m.DB{db}, db.Ping()
}
