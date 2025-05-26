package datastore

import (
	"context"

	"github.com/meysam81/tarzan/cmd/config"
	"github.com/meysam81/tarzan/cmd/models"
)

type Datastore interface {
	ListPosts(context.Context, ...func(*models.Post)) (*[]models.Post, error)
	FetchPost(context.Context, int, ...func(*models.Post)) (*models.Post, error)
	InsertEmail(context.Context, *models.EmailInsertDB) error
	FindAuthorizedSenderByEmail(ctx context.Context, email string) (*models.AuthorizedSender, error)
	Close() error
}

type BuildDatastore interface {
	NewDatastore(ctx context.Context, cfg *config.Config) (Datastore, error)
}
