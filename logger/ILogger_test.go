package logger

import (
	"testing"
)

func TestLogger_ReturnsDefaultLogger(t *testing.T) {
	logger := Logger()
	if logger == nil {
		t.Errorf("Expected logger to have a value!")
	}
}

func TestUseLogger_UsesInitializedLogger(t *testing.T) {
	result := ""
	UseLogger(MockLogger{logMessage: &result})
	Logger().Log("", "foobar")

	if result != "foobar" {
		t.Errorf("Expected LogResult to be foobar!")
	}
}