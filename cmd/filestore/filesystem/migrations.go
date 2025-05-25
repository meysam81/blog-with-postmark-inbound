package filesystem

import (
	"os"
)

func (fs *fileSystem) runMigrations() error {
	err := os.MkdirAll(fs.AttachmentPath, 0755)
	return err
}
