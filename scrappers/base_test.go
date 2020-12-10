package scrappers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Should return the scrapper name
func Test_baseScrapper_getName(t *testing.T) {
	bs := &baseScrapper{name: "test name"}
	assert.Equal(t, bs.getName(), "test name")
}

// Should return a created language
func Test_baseScrapper_returnLang(t *testing.T) {
	code := "foobar"
	bs := &baseScrapper{name: code}
	lang := bs.returnLang("brainfuck", 42)

	assert.Equal(t, lang.position, uint(42))
	assert.Equal(t, lang.title, "brainfuck")
	assert.Equal(t, lang.source, code)
}
