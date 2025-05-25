package sqlite

import (
	"context"
	"log"
	"time"

	"github.com/meysam81/tarzan/cmd/models"
)

func (s *Sqlite) List(ctx context.Context, transformers ...func(*models.Post)) (*[]models.Post, error) {
	ctxT, cancelT := context.WithTimeout(ctx, 3*time.Second)
	defer cancelT()
	rows, err := s.DB.QueryContext(ctxT, "SELECT id, title, content, author_email, author_name, created_at FROM posts ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}()

	var posts []models.Post
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.AuthorEmail, &p.AuthorName, &p.CreatedAt); err != nil {
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

func (s *Sqlite) Insert(ctx context.Context, email *models.EmailInsertDB) error {
	ctxT, cancelT := context.WithTimeout(ctx, 3*time.Second)
	defer cancelT()

	_, err := s.DB.ExecContext(ctxT, "INSERT INTO posts (title, content, author_email, author_name) VALUES (?, ?, ?, ?)",
		email.Subject, email.Content, email.From, email.FromName)
	return err
}

func (s *Sqlite) FindAuthorizedSenderByEmail(ctx context.Context, email string) (*models.AuthorizedSender, error) {
	ctxT, cancelT := context.WithTimeout(ctx, 3*time.Second)
	defer cancelT()

	row := s.DB.QueryRowContext(ctxT, "SELECT id, email, created_at FROM authorized_senders WHERE email = ?", email)
	var sender models.AuthorizedSender
	err := row.Scan(&sender.ID, &sender.Email, &sender.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &sender, nil
}
