package sqlite

import (
	"context"
	"log"
	"time"

	. "github.com/meysam81/tarzan/cmd/models"
)

func (s *Sqlite) List(ctx context.Context, transformers ...func(*Post)) (*[]Post, error) {
	ctxT, cancelT := context.WithTimeout(ctx, 3*time.Second)
	defer cancelT()
	rows, err := s.DB.QueryContext(ctxT, "SELECT id, title, content, author_email, created_at FROM posts ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

	var posts []Post
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.AuthorEmail, &p.CreatedAt); err != nil {
			log.Printf("Error scanning post: %v", err)
			continue
		}
		for _, t := range transformers {
			t(&p)
		}
		posts = append(posts, p)
	}

	return &posts, nil
}

func (s *Sqlite) Insert(ctx context.Context, email *InboundEmail) error {
	content := email.HtmlBody
	if content == "" {
		content = email.TextBody
	}

	ctxT, cancelT := context.WithTimeout(ctx, 3*time.Second)
	defer cancelT()

	_, err := s.DB.ExecContext(ctxT, "INSERT INTO posts (title, content, author_email) VALUES (?, ?, ?)",
		email.Subject, content, email.From)
	return err
}
