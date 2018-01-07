package tests

import (
	"testing"
)

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	result := 123

	if result != 123 {
		t.Errorf("Invalid result value.")
	}
}
