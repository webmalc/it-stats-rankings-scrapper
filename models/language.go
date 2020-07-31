package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/qor/validations"
	"github.com/webmalc/it-stats-rankings-scrapper/services"
)

// Language is the service struct
type Language struct {
	gorm.Model
	Source        string        `gorm:"size:255;not null;index:source;" valid:"required"`
	Title         string        `gorm:"size:255;not null;index:title" valid:"required"`
	Synonym       string        `gorm:"size:255;"`
	Position      uint          `gorm:"not null;index:position" valid:"type(uint)"`
	nameProcessor nameProcessor `gorm:"-"`
}

// Validate validates struct
func (l *Language) Validate(db *gorm.DB) {
	if !govalidator.IsIn(l.Source, l.GetAvailableSources()...) {
		_ = db.AddError(
			validations.NewError(l, "Source", "the source is inccorrect"),
		)
	}
}

// GetAvailableSources returns the available sources
func (l *Language) GetAvailableSources() []string {
	config := NewConfig()
	return config.scrappers
}

// InitNameProcessor creates a name normalizer
func (l *Language) InitNameProcessor() {
	if l.nameProcessor == nil {
		l.nameProcessor = services.NewNameNormalizer()
	}
}

// BeforeSave runs before saving the object
func (l *Language) BeforeSave() (err error) {
	l.InitNameProcessor()
	l.Title = l.nameProcessor.Normalize(l.Title)
	l.Synonym = l.nameProcessor.GetSynonym(l.Title)
	return nil
}

// Migrate migrates the DB
func Migrate(migrater AutoMigrater) {
	migrater.AutoMigrate(&Language{})
}
