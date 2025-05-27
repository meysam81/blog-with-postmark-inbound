package filesystem

import (
	"context"
	"encoding/base64"
	"os"
	"path/filepath"

	"github.com/meysam81/tarzan/cmd/config"
	"github.com/meysam81/tarzan/cmd/filestore"
)

func (*Builder) NewFilestore(ctx context.Context, cfg *config.Config) (filestore.Filestore, error) {
	attachmentPath := filepath.Join(cfg.StoragePath, "attachments")
	fs := &fileSystem{AttachmentPath: attachmentPath}
	err := fs.runMigrations()
	if err != nil {
		return nil, err
	}
	return fs, nil
}

func (fs *fileSystem) Save(filename string, data []byte) error {
	filepath := filepath.Join(fs.AttachmentPath, filename)

	if err := os.WriteFile(filepath, data, 0644); err != nil {
		return err
	}

	return nil
}
func (fs *fileSystem) Load(filename string) (string, error) {
	path := filepath.Join(fs.AttachmentPath, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", nil
	}
	return base64.StdEncoding.EncodeToString(data), nil

}
