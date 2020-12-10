package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/webmalc/it-stats-rankings-scrapper/cmd/mocks"
	"github.com/webmalc/it-stats-rankings-scrapper/common/test"
)

// Should run the root command and log an error.
func TestCommandRouter_Run(t *testing.T) {
	m := &mocks.ErrorLogger{}
	r := &mocks.Runner{}
	c := &mocks.ContextRunner{}
	cr := NewCommandRouter(m, r, r, c)
	os.Args = []string{"invalid", "invalid"}
	m.On("Error", mock.Anything).Return(nil).Once()
	cr.Run()
	m.AssertExpectations(t)
}

// Should create a command router object.
func TestNewCommandRouter(t *testing.T) {
	l := &mocks.ErrorLogger{}
	r := &mocks.Runner{}
	c := &mocks.ContextRunner{}
	cr := NewCommandRouter(l, r, r, c)
	assert.Equal(t, l, cr.logger)
	assert.Equal(t, r, cr.adminRunner)
	assert.Equal(t, c, cr.scrappersRunner)
	assert.NotNil(t, cr.rootCmd)
}

func TestCommandRouter_scrap(t *testing.T) {
	r := &mocks.Runner{}
	a := &mocks.Runner{}
	scrapper := &mocks.ContextRunner{}
	cr := NewCommandRouter(&mocks.ErrorLogger{}, r, a, scrapper)
	args := []string{"pypl", "tiobe"}
	scrapper.On("Run", args, mock.Anything).Return(nil).Once()
	cr.scrap(&cobra.Command{}, args)
	scrapper.AssertExpectations(t)
	cr.timeout = 0
	timeout, _ := cr.rootCmd.Flags().GetDuration("timeout")
	assert.Equal(t, cr.timeout, timeout)
}

func TestCommandRouter_admin(t *testing.T) {
	bindata := &mocks.Runner{}
	admin := &mocks.Runner{}
	scrapper := &mocks.ContextRunner{}
	cr := NewCommandRouter(&mocks.ErrorLogger{}, admin, bindata, scrapper)
	admin.On("Run", mock.Anything).Return(nil).Once()
	cr.admin(&cobra.Command{}, []string{})
	admin.AssertExpectations(t)
}

func TestCommandRouter_bindatafs(t *testing.T) {
	bindata := &mocks.Runner{}
	admin := &mocks.Runner{}
	scrapper := &mocks.ContextRunner{}
	cr := NewCommandRouter(&mocks.ErrorLogger{}, admin, bindata, scrapper)
	bindata.On("Run", mock.Anything).Return(nil).Once()
	cr.bindatafs(&cobra.Command{}, []string{})
	bindata.AssertExpectations(t)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
