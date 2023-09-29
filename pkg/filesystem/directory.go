package filesystem

import (
	"fmt"
	"os"
)

func CreateDirectoryIfNotExists(path string) error {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)

		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}

		return nil
	} else if err != nil {
		return fmt.Errorf("failed to check directory existence: %v", err)
	}

	return nil
}
