package sqlite

import "database/sql"

type Sqlite struct {
	DB *sql.DB
}

type Builder struct{}
