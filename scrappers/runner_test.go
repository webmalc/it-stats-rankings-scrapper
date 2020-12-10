package scrappers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/it-stats-rankings-scrapper/common/test"
	"github.com/webmalc/it-stats-rankings-scrapper/scrappers/mocks"
)

func TestRunner_getSelectedScrappers(t *testing.T) {
	// TODO:  TEST IT FIRST !!!!!!
}

func TestRunner_Run(t *testing.T) {
}

func TestRunner_save(t *testing.T) {
}

func TestRunner_fanIn(t *testing.T) {
}

func TestNewRunner(t *testing.T) {
	log := &mocks.Logger{}
	adder := &mocks.LanguageAdder{}
	scrappers := map[string]scrapper{
		"one": &scrapperMock{}, "two": &scrapperMock{},
	}
	runner := NewRunner(adder, log)
	assert.Equal(t, log, runner.logger)
	assert.Equal(t, adder, runner.adder)
	assert.NotEmpty(t, runner.scrappers)

	runner.scrappers = scrappers
	assert.Equal(t, runner.scrappers, scrappers)
}

// Setups the tests
func TestMain(m *testing.M) {
	test.Run(m)
}
