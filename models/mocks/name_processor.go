package mocks

import (
	"github.com/stretchr/testify/mock"
)

// NameProcessor logs errors.
type NameProcessor struct {
	mock.Mock
}

// Normalize is a method mock
func (m *NameProcessor) Normalize(name string) string {
	arg := m.Called(name)
	return arg.Get(0).(string)
}

// GetSynonym is a method mock
func (m *NameProcessor) GetSynonym(name string) string {
	arg := m.Called(name)
	return arg.Get(0).(string)
}
