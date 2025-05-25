package sqlite

import (
	"context"
	"time"
)

func (s *Sqlite) runMigrations(ctx context.Context) error {
	migrations := []string{
		`CREATE TABLE IF NOT EXISTS posts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        content TEXT,
        author_email TEXT,
				author_name TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    )`,
		`CREATE INDEX IF NOT EXISTS idx_posts_created_at ON posts(created_at)`,
		`CREATE INDEX IF NOT EXISTS idx_posts_author_email_title ON posts(author_email, title)`,
		`CREATE TABLE IF NOT EXISTS authorized_senders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT UNIQUE NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE INDEX IF NOT EXISTS idx_authorized_senders_created_at ON authorized_senders(created_at)`,
		`CREATE INDEX IF NOT EXISTS idx_authorized_senders_email ON authorized_senders(email)`,
	}

	for _, migration := range migrations {
		ctxT, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		_, err := s.DB.ExecContext(ctxT, migration)
		if err != nil {
			return err
		}
	}

	return nil
}
