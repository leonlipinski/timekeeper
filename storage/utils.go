package storage

import (
	"os"
	"path/filepath"
)

// GetConfigPath returns the path to the configuration file in the user's home directory.
var GetConfigPath = func(filename string) (string, error) {
	// Get the user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "timekeeper", filename), nil
}

var RenameFile = func(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
