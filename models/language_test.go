package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/it-stats-rankings-scrapper/common/db"
	"github.com/webmalc/it-stats-rankings-scrapper/common/test"
	"github.com/webmalc/it-stats-rankings-scrapper/models/mocks"
)

// Should migrate the DB.
func TestMigrate(t *testing.T) {
	am := &mocks.AutoMigrater{}
	conn := db.NewConnection()
	defer conn.Close()
	args := []interface{}{&Language{}}
	am.On("AutoMigrate", args...).Return(conn.DB).Once()
	Migrate(am)
	am.AssertExpectations(t)
}

// Should Validate the object
func TestLanguage_Validate(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	Migrate(conn)
	processor := &mocks.NameProcessor{}
	repo := NewLanguageRepository(conn, processor)
	l := repo.NewLanguage("invalid", "php", 12)
	errors := conn.Create(l).GetErrors()
	assert.NotEmpty(t, errors)
	l = repo.NewLanguage("tiobe", "go", 1)
	processor.On("Normalize", "go").Return("go").Once()
	processor.On("GetSynonym", "go").Return("go").Once()
	errors = conn.Create(l).GetErrors()
	assert.Empty(t, errors)
	var count int
	conn.Find(&Language{}).Count(&count)
	assert.Equal(t, 1, count)
	processor.AssertExpectations(t)
}

// Should create a name processor
func TestLanguage_InitNameProcessor(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	Migrate(conn)
	lang := &Language{}
	assert.Nil(t, lang.nameProcessor)
	lang.InitNameProcessor()
	assert.NotNil(t, lang.nameProcessor)
}

// Should run the hook
func TestService_BeforeSave(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	Migrate(conn)
	processor := &mocks.NameProcessor{}
	repo := NewLanguageRepository(conn, processor)
	processor.On("Normalize", "php").Return("php").Once()
	processor.On("GetSynonym", "php").Return("php").Once()
	repo.AddLanguage("tiobe", "php", 12)
	processor.AssertExpectations(t)
}

// Should return the list of available sources
func TestLanguage_GetAvailableSources(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	Migrate(conn)
	lang := &Language{}
	assert.Contains(t, lang.GetAvailableSources(), "tiobe")
	assert.Contains(t, lang.GetAvailableSources(), "pypl")
}

// Setups the tests.
func TestMain(m *testing.M) {
	test.Run(m)
}
