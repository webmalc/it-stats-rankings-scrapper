package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/qor/validations"
)

// Language is the service struct
type Language struct {
	gorm.Model
	Source        string `gorm:"size:255;not null;index:source;" valid:"required"`
	Title         string `gorm:"size:255;not null;index:title" valid:"required"`
	Synonym       string `gorm:"size:255;"`
	Position      uint   `gorm:"not null;index:position" valid:"type(uint)"`
	nameProcessor nameProcessor
}

// Validate validates struct
func (l *Language) Validate(db *gorm.DB) {
	config := NewConfig()
	if !govalidator.IsIn(l.Source, config.scrappers...) {
		_ = db.AddError(
			validations.NewError(l, "Source", "the source is inccorrect"),
		)
	}
}

// BeforeSave runs before saving the object
func (l *Language) BeforeSave() (err error) {
	l.Title = l.nameProcessor.Normalize(l.Title)
	l.Synonym = l.nameProcessor.GetSynonym(l.Title)
	return nil
}

// Migrate migrates the DB
func Migrate(migrater AutoMigrater) {
	migrater.AutoMigrate(&Language{})
}
