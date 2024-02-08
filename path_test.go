package lunch_test

import (
	"errors"
	"os"
	"testing"

	"github.com/mishamyrt/go-lunch"
)

func TestSharedPath(t *testing.T) {
	t.Parallel()
	t.Run("with non-empty basename", func(t *testing.T) {
		t.Parallel()
		basename := "testAgent"
		expected := "/" + lunch.AgentsDirectory + "/" + basename + ".plist"
		got, err := lunch.SharedPath(basename)
		if err != nil {
			t.Errorf("SharedPath() error = %v, wantErr false", err)
		}
		if got != expected {
			t.Errorf("SharedPath() = %v, want %v", got, expected)
		}
	})

	t.Run("with empty basename", func(t *testing.T) {
		t.Parallel()
		_, err := lunch.SharedPath("")
		if err == nil {
			t.Errorf("SharedPath() error = nil, wantErr true")
		}
		if !errors.Is(err, lunch.ErrEmptyBasename) {
			t.Errorf("SharedPath() error = %v, want %v", err, lunch.ErrEmptyBasename)
		}
	})
}

func TestUserPath(t *testing.T) {
	t.Parallel()
	mockHomeDir := "/Users/test"
	// Setup environment
	homeEnv := createMock("HOME")
	homeEnv.mock(mockHomeDir)
	defer homeEnv.clear()

	t.Run("with valid basename", func(t *testing.T) {
		basename := "testAgent"
		expected := mockHomeDir + "/" + lunch.AgentsDirectory + "/" + basename + ".plist"
		got, err := lunch.UserPath(basename)
		if err != nil {
			t.Errorf("Unexpected error on UserPath(): %v", err)
		}
		if got != expected {
			t.Errorf("Unexpected UserPath() return value: got %v, want %v", got, expected)
		}
	})

	t.Run("with empty basename", func(t *testing.T) {
		_, err := lunch.UserPath("")
		if err == nil {
			t.Errorf("Unexpected nil error on UserPath(\"\")")
		}
		if !errors.Is(err, lunch.ErrEmptyBasename) {
			t.Errorf("Unexpected error on UserPath(): got %v, want %v", err, lunch.ErrEmptyBasename)
		}
	})

	t.Run("with empty $HOME", func(t *testing.T) {
		os.Setenv("HOME", "")
		_, err := lunch.UserPath("testAgent")
		if err == nil {
			t.Errorf("Unexpected nil error on UserPath(): want \"$HOME is not defined\"")
		}
	})
}

type envMock struct {
	variable      string
	originalValue string
}

func createMock(variable string) envMock {
	return envMock{
		variable: variable,
	}
}

func (e *envMock) mock(value string) {
	e.originalValue = os.Getenv(e.variable)
	os.Setenv(e.variable, value)
}

func (e *envMock) clear() {
	os.Setenv(e.variable, e.originalValue)
}
