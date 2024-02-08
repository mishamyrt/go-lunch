// Package lunch provides tools for working with autorun on macOS.
package lunch

import (
	"errors"
	"os"
	"strings"

	"github.com/mishamyrt/go-lunch/plist"
)

const devNull = "/dev/null"

// Agent represent macOS launch agent.
type Agent struct {
	PackageName       string
	KeepAlive         bool
	Command           string
	StandardOutPath   string
	StandardErrorPath string
	Path              string
}

// Exists checks if login item is exists at the Path.
func (a *Agent) Exists() (bool, error) {
	if len(a.Path) == 0 {
		return false, ErrEmptyPath
	}
	info, err := os.Stat(a.Path)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return !info.IsDir(), nil
}

// Remove login item.
func (a *Agent) Remove() error {
	exists, err := a.Exists()
	if err != nil {
		return err
	}
	if !exists {
		return os.ErrNotExist
	}
	return os.Remove(a.Path)
}

// Write login item.
func (a *Agent) Write() error {
	if len(a.Path) == 0 {
		return ErrEmptyPath
	}
	content := []byte(a.String())
	err := os.WriteFile(a.Path, content, 0644) //nolint:gomnd
	if err != nil {
		return err
	}
	return nil
}

// String returns Launch Agent plist content
func (a *Agent) String() string {
	outPath := a.StandardOutPath
	errPath := a.StandardErrorPath
	if len(outPath) == 0 {
		outPath = devNull
	}
	if len(errPath) == 0 {
		errPath = devNull
	}
	arguments := strings.Split(a.Command, " ")
	props := plist.New().
		AddBool("KeepAlive", a.KeepAlive).
		AddString("Label", a.PackageName).
		AddStringArray("ProgramArguments", arguments).
		AddBool("RunAtLoad", true).
		AddString("StandardErrorPath", errPath).
		AddString("StandardOutPath", outPath)
	return props.String()
}

// ErrEmptyPath is returned when Agent has an empty path
var ErrEmptyPath = errors.New("agents 'path' field is empty")
