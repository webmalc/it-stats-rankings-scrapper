package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/webmalc/it-stats-rankings-scrapper/models/mocks"
	"github.com/webmalc/it-stats-rankings-scrapper/services"
	"github.com/webmalc/services-scrapper/common/db"
)

// Should return a new language object
func TestLanguageRepository_NewLanguage(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	repo := NewLanguageRepository(conn, &mocks.NameProcessor{})
	s := repo.NewLanguage("tiobe", "python", 1)
	assert.Equal(t, "python", s.Title)
	assert.Equal(t, "tiobe", s.Source)
	assert.Equal(t, uint(1), s.Position)
}

// Should return a new LanguageRepository
func TestNewLanguageRepository(t *testing.T) {
	conn := db.NewConnection()
	processor := &mocks.NameProcessor{}
	defer conn.Close()
	repo := NewLanguageRepository(conn, processor)
	assert.Equal(t, conn, repo.db)
	assert.Equal(t, processor, repo.nameProcessor)
}

// Should create a new language
func TestLanguageRepository_CreateLanguage(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	Migrate(conn)
	repo := NewLanguageRepository(conn, services.NewNameNormalizer())
	l := repo.NewLanguage("tiobe", "python", 1)
	errors := repo.CreateLanguage(l)
	assert.Empty(t, errors)
	l = &Language{}
	conn.Set("gorm:auto_preload", true).First(l)
	assert.Equal(t, "tiobe", l.Source)
	assert.Equal(t, "python", l.Title)
	assert.Equal(t, uint(1), l.Position)

	l = repo.NewLanguage("invalid", "php", 9)
	errors = repo.CreateLanguage(l)
	assert.NotEmpty(t, errors)
}

// Should create a new language entry in the DB
func TestLanguageRepository_AddLanguage(t *testing.T) {
	conn := db.NewConnection()
	defer conn.Close()
	Migrate(conn)
	repo := NewLanguageRepository(conn, services.NewNameNormalizer())
	errors := repo.AddLanguage("tiobe", "python", 10)
	l := &Language{}
	conn.Set("gorm:auto_preload", true).First(l)

	assert.Empty(t, errors)
	assert.Equal(t, "tiobe", l.Source)
	assert.Equal(t, "python", l.Title)
	assert.Equal(t, uint(10), l.Position)
}
