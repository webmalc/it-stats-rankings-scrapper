package services

import "strings"

// NameNormalizer is the service repository struct
type NameNormalizer struct {
	synonyms     map[string]string
	replacements map[string]string
}

// Normalize normalizes the provided name
func (p *NameNormalizer) Normalize(name string) string {
	result := strings.ToLower(name)
	if val, ok := p.replacements[result]; ok {
		return val
	}
	return result
}

// GetSynonym returns a synonym of the provided name
func (p *NameNormalizer) GetSynonym(name string) string {
	return p.synonyms[name]
}

// NewNameNormalizer returns a new object
func NewNameNormalizer() *NameNormalizer {
	n := &NameNormalizer{
		synonyms: map[string]string{
			"go":                "golang",
			"visual basic":      "vba",
			"assembly language": "assembler",
			"typescript":        "ts",
			"javascript":        "js",
		},
		replacements: map[string]string{"c/c++": "c++"},
	}
	for k, v := range n.synonyms {
		n.synonyms[v] = k
	}
	return n
}
