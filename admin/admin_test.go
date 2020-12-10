package admin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/it-stats-rankings-scrapper/common/db"
)

// Should run an admin
func TestAdmin_Run(t *testing.T) {
}

// Should create a new admin
func TestNewAdmin(t *testing.T) {
	conn := db.NewConnection()
	adm := NewAdmin(conn.DB)
	assert.Equal(t, adm.db, conn.DB)
	assert.NotNil(t, adm.config)
	assert.NotNil(t, adm.admin)
}

// Should return a new startup message
func TestAdmin_startupMessage(t *testing.T) {
	conn := db.NewConnection()
	adm := NewAdmin(conn.DB)
	assert.Contains(t, adm.getStartupMessage(), adm.config.AdminPath)
	assert.Contains(t, adm.getStartupMessage(), adm.config.AdminURL)
	assert.Contains(t, adm.getStartupMessage(), "http")
	adm.config.AdminSSL = true
	adm.config.AdminPath = "/new"
	assert.Contains(t, adm.getStartupMessage(), "https")
	assert.Contains(t, adm.getStartupMessage(), "new")
}
