package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameNormalizer_Normalize(t *testing.T) {
	n := NewNameNormalizer()
	assert.Equal(t, "typescript", n.Normalize("TypeScript"))
	assert.Equal(t, "visual basic", n.Normalize("Visual basic"))
	assert.Equal(t, "c++", n.Normalize("C/C++"))
}

// Should return a synonym
func TestNameNormalizer_GetSynonym(t *testing.T) {
	n := NewNameNormalizer()
	assert.Equal(t, "golang", n.GetSynonym("go"))
	assert.Equal(t, "go", n.GetSynonym("golang"))
	assert.Equal(t, "", n.GetSynonym("invalid"))
}

// Should return a new object
func TestNewNameNormalizer(t *testing.T) {
	n := NewNameNormalizer()
	assert.Equal(t, "golang", n.synonyms["go"])
	assert.Equal(t, "go", n.synonyms["golang"])
	assert.Equal(t, "typescript", n.synonyms["ts"])
	assert.Equal(t, "ts", n.synonyms["typescript"])
	assert.Equal(t, "c++", n.replacements["c/c++"])
}
