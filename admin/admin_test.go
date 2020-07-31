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
