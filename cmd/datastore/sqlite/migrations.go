package sqlite

import (
	"context"
)

func (s *Sqlite) runMigrations(ctx context.Context) error {
	_, err := s.DB.Exec(`CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        content TEXT,
        author_email TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    )`)
	if err != nil {
		return err
	}

	return nil
}
