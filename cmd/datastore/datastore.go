package datastore

import (
	"context"

	"github.com/meysam81/tarzan/cmd/config"
	"github.com/meysam81/tarzan/cmd/models"
)

type Datastore interface {
	List(context.Context, ...func(*models.Post)) (*[]models.Post, error)
	Insert(context.Context, *models.InboundEmail) error
	Close() error
}

type BuildDatastore interface {
	NewDatastore(ctx context.Context, cfg *config.Config) (Datastore, error)
}
