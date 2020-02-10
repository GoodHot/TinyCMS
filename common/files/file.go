package files

import (
	"os"
)

func FileExists(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !file.IsDir()
}

func DirExists(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return file.IsDir()
}
