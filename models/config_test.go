package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Should return the config object
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Contains(t, c.scrappers, "tiobe")
	assert.Contains(t, c.scrappers, "pypl")
	assert.Contains(t, c.scrappers, "redmonk")
}
