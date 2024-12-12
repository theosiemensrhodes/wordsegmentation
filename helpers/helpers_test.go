package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanString(t *testing.T) {
	assert.Equal(t, CleanString("  abc  "), "abc")
	assert.Equal(t, CleanString("HElLo"), "hello")
	assert.Equal(t, CleanString("éàï"), "eai")
	assert.Equal(t, CleanString("a@3£4&^>?d"), "a34d")
}
