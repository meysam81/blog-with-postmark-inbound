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
		createCSPViolationsTable,
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

const createCSPViolationsTable = `
CREATE TABLE IF NOT EXISTS csp_violations (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	document_uri TEXT NOT NULL,
	referrer TEXT,
	violated_directive TEXT NOT NULL,
	effective_directive TEXT,
	original_policy TEXT,
	disposition TEXT,
	blocked_uri TEXT,
	line_number INTEGER DEFAULT 0,
	source_file TEXT,
	status_code INTEGER DEFAULT 0,
	script_sample TEXT,
	user_agent TEXT,
	ip_address TEXT,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_csp_violations_created_at ON csp_violations(created_at);
CREATE INDEX IF NOT EXISTS idx_csp_violations_violated_directive ON csp_violations(violated_directive);
CREATE INDEX IF NOT EXISTS idx_csp_violations_blocked_uri ON csp_violations(blocked_uri);
`
