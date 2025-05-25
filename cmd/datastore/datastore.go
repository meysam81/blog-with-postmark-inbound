package datastore

import (
	"context"

	"github.com/meysam81/tarzan/cmd/models"
)

type Datastore interface {
	List(context.Context, ...func(*models.Post)) (*[]models.Post, error)
	Insert(context.Context, *models.InboundEmail) error
	Close() error
}
