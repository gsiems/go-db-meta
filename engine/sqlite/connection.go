package sqlite

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"

	e "github.com/gsiems/go-db-meta/engine"
	m "github.com/gsiems/go-db-meta/model"
)

// OpenDB opens a database connection and returns a DB reference.
func OpenDB(c *e.ConnectInfo) (*m.DB, error) {

	_, err := os.Stat(c.File)
	if err != nil {
		return nil, err
	}

	// Options can be given using the following format: KEYWORD=VALUE and multiple options can be combined with the & ampersand.
	// mode=ro

	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=ro", c.File))
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
