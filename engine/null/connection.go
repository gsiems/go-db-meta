package null

import (
	"database/sql"

	e "github.com/gsiems/go-db-meta/engine"
	m "github.com/gsiems/go-db-meta/model"
)

// OpenDB opens a database connection and returns a DB reference.
func OpenDB(c *e.ConnectInfo) (*m.DB, error) {

	db, err := sql.Open("nosuch", "")

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
