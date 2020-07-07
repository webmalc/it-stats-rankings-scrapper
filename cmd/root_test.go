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
	cr := NewCommandRouter(m, r)
	os.Args = []string{"invalid", "invalid"}
	m.On("Error", mock.Anything).Return(nil).Once()
	cr.Run()
	m.AssertExpectations(t)
}

// Should create a command router object.
func TestNewCommandRouter(t *testing.T) {
	m := &mocks.ErrorLogger{}
	r := &mocks.Runner{}
	cr := NewCommandRouter(m, r)
	assert.Equal(t, m, cr.logger)
	assert.Equal(t, r, cr.scrappersRunner)
	assert.NotNil(t, cr.rootCmd)
}

func TestCommandRouter_scrap(t *testing.T) {
	r := &mocks.Runner{}
	cr := NewCommandRouter(&mocks.ErrorLogger{}, r)
	args := []string{"redmonk", "tiobe"}
	r.On("Run", args).Return(nil).Once()
	cr.scrap(&cobra.Command{}, args)
	r.AssertExpectations(t)
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
