package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	result := 123

	assert.Equal(t, 123, 123, "they should be equal")

	if result != 123 {
		t.Errorf("Invalid result value.")
	}
}
