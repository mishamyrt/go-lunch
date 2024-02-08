package lunch

import (
	"errors"
	"os"
	"path"
)

// AgentsDirectory represents directory where macOS stores launch agents.
const AgentsDirectory = "Library/LaunchAgents"

// ErrEmptyBasename is returned when provided basename is empty.
var ErrEmptyBasename = errors.New("basename is empty")

// UserPath generates path for current user's package login item.
func UserPath(basename string) (string, error) {
	if len(basename) == 0 {
		return "", ErrEmptyBasename
	}
	userDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(userDir, AgentsDirectory, basename+".plist"), nil
}

// SharedPath generates path for system login item.
func SharedPath(basename string) (string, error) {
	if len(basename) == 0 {
		return "", ErrEmptyBasename
	}
	return path.Join("/", AgentsDirectory, basename+".plist"), nil
}
