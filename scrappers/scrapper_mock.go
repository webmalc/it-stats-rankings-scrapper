package scrappers

import (
	"github.com/stretchr/testify/mock"
)

// Scrapper is a mock object
type scrapperMock struct {
	mock.Mock
}

// getName is method mock
func (m *scrapperMock) getName() string {
	arg := m.Called()
	return arg.Get(0).(string)
}

// scrap is method mock
func (m *scrapperMock) scrap() []*language {
	arg := m.Called()
	return arg.Get(0).([]*language)
}
