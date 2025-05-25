package sqlite

import (
	"context"
	"time"

	"github.com/meysam81/tarzan/cmd/config"
	"github.com/meysam81/x/sqlite"
)

func NewDatastore(ctx context.Context, cfg *config.Config) (*Sqlite, error) {
	db, err := createDB(ctx, cfg)
	if err != nil {
		return nil, err
	}

	db.runMigrations(ctx)

	return db, err
}

func createDB(ctx context.Context, cfg *config.Config) (*Sqlite, error) {
	ctxT, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	db, err := sqlite.NewDB(ctxT, cfg.DbPath, sqlite.WithJournalMode(""), sqlite.WithMode(""))
	if err != nil {
		return nil, err
	}

	return &Sqlite{DB: db}, nil
}

func (s *Sqlite) Close() error {
	return s.DB.Close()
}
