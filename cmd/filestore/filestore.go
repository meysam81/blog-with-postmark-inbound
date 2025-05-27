package filestore

import (
	"context"

	"github.com/meysam81/tarzan/cmd/config"
)

type Filestore interface {
	Save(filename string, data []byte) error
	Load(filename string) (string, error)
	LoadBytes(filename string) ([]byte, error)
}

type BuildFilestore interface {
	NewFilestore(ctx context.Context, cfg *config.Config) (Filestore, error)
}
