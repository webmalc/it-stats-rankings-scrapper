package mocks

import (
	"github.com/stretchr/testify/mock"
)

// LanguageAdder is a mock struct.
type LanguageAdder struct {
	mock.Mock
}

// AddLanguage is method mock
func (m *LanguageAdder) AddLanguage(
	source, title string, position uint,
) []error {
	arg := m.Called(source, title, position)
	return arg.Get(0).([]error)
}
