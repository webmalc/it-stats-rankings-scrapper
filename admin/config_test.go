package admin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/it-stats-rankings-scrapper/common/test"
)

// Should return the config object
func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.Equal(t, ":9000", c.AdminURL)
	assert.Equal(t, "", c.AdminPath)
}

// Setups the tests
func TestMain(m *testing.M) {
	test.Run(m)
}
