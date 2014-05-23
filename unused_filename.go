package pathutils

import (
	"fmt"
	"path/filepath"
)

func UnusedFilename(exists func(string) bool, path string, maxRetries int) (newPath string, err error) {
	dir, fileName := filepath.Split(path)

	name, ext := SplitNameExt(fileName)

	newPath = path

	var newName string

	for i := 1; i <= maxRetries; i++ {
		if !exists(newPath) {
			return newPath, nil
		}

		newName = JoinNameExt(fmt.Sprintf("%s (%d)", name, i), ext)
		newPath = filepath.Join(dir, newName)
	}

	err = fmt.Errorf("too many retries UnusedFilename(_, %s, %d)", path, maxRetries)

	return
}
