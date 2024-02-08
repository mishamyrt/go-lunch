package lunch_test

import (
	"errors"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/mishamyrt/go-lunch"
)

func tempFile(t *testing.T, name string) string {
	tmpFile := path.Join(t.TempDir(), name+".plist")
	err := os.WriteFile(tmpFile, []byte("agent"), 0644)
	if err != nil {
		t.Fatalf("Unable to create temporary file: %v", err)
	}
	return tmpFile
}

func TestAgentExists(t *testing.T) {
	t.Parallel()

	t.Run("file exists", func(t *testing.T) {
		t.Parallel()
		tmpFile := tempFile(t, "file")
		defer os.Remove(tmpFile)

		agent := lunch.Agent{Path: tmpFile}
		exists, err := agent.Exists()
		if err != nil {
			t.Errorf("Exists() error = %v, wantErr false", err)
		}
		if !exists {
			t.Error("Exists() = false, want true for existing file")
		}
	})

	t.Run("file does not exist", func(t *testing.T) {
		t.Parallel()
		agent := lunch.Agent{Path: "/path/does/not/exist"}
		exists, err := agent.Exists()
		if err != nil {
			t.Errorf("Exists() error = %v, wantErr nil", err)
		}
		if exists {
			t.Error("Exists() = true, want false for non-existing file")
		}
	})

	t.Run("empty path", func(t *testing.T) {
		t.Parallel()
		agent := lunch.Agent{}
		_, err := agent.Exists()
		if !errors.Is(err, lunch.ErrEmptyPath) {
			t.Errorf("Exists() error = %v, want %v", err, lunch.ErrEmptyPath)
		}
	})
}

func TestAgentRemove(t *testing.T) {
	t.Parallel()

	t.Run("remove existing file", func(t *testing.T) {
		t.Parallel()
		tmpFile := tempFile(t, "file")
		defer os.Remove(tmpFile)

		agent := lunch.Agent{Path: tmpFile}
		if err := agent.Remove(); err != nil {
			t.Errorf("Remove() error = %v, wantErr nil", err)
		}
		if _, err := os.Stat(tmpFile); !os.IsNotExist(err) {
			t.Error("Remove() file still exists, want file to be removed")
		}
	})

	t.Run("remove non-existing file", func(t *testing.T) {
		t.Parallel()
		agent := lunch.Agent{Path: "/path/does/not/exist"}
		if err := agent.Remove(); !errors.Is(err, os.ErrNotExist) {
			t.Errorf("Remove() error = %v, want %v", err, os.ErrNotExist)
		}
	})

	t.Run("empty path", func(t *testing.T) {
		agent := lunch.Agent{}
		err := agent.Remove()
		if !errors.Is(err, lunch.ErrEmptyPath) {
			t.Errorf("Remove() error = %v, want %v", err, lunch.ErrEmptyPath)
		}
	})
}

func TestAgentWrite(t *testing.T) {
	t.Run("write to new file", func(t *testing.T) {
		tmpFile := tempFile(t, "file")
		defer os.Remove(tmpFile)

		agent := lunch.Agent{
			PackageName:       "com.example.test",
			KeepAlive:         true,
			Command:           "/usr/bin/example",
			StandardOutPath:   "",
			StandardErrorPath: "",
			Path:              tmpFile,
		}

		if err := agent.Write(); err != nil {
			t.Errorf("Write() error = %v, wantErr nil", err)
		}

		content, err := os.ReadFile(tmpFile)
		if err != nil {
			t.Fatalf("ReadFile() error = %v", err)
		}

		if !strings.Contains(string(content), agent.PackageName) {
			t.Errorf("Write() content does not contain PackageName, content = %s", string(content))
		}
	})

	t.Run("empty path", func(t *testing.T) {
		agent := lunch.Agent{}
		err := agent.Write()
		if !errors.Is(err, lunch.ErrEmptyPath) {
			t.Errorf("Write() error = %v, want %v", err, lunch.ErrEmptyPath)
		}
	})
}
