package tools

import (
	"os"
	"path"
)

// MkdirAllIfNotExists
// If the directory (or the directory of the file) does not exist, create it.
func MkdirAllIfNotExists(pathname string, perm os.FileMode) error {
	dir := path.Dir(pathname)
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, perm); err != nil {
				return err
			}
		}
	}
	return nil
}
