package models

import "github.com/jinzhu/gorm"

// AutoMigrater auto migrate the DB
type AutoMigrater interface {
	AutoMigrate(values ...interface{}) *gorm.DB
}

// Creator interface
type Creator interface {
	Create(value interface{}) *gorm.DB
}

// Finder interface
type Finder interface {
	Model(value interface{}) *gorm.DB
}

// Database interface
type Database interface {
	Creator
	Finder
}

// normalizer interface
type normalizer interface {
	Normalize(name string) string
}

// synonymGetter interface
type synonymGetter interface {
	GetSynonym(name string) string
}

// nameProcessor interface
type nameProcessor interface {
	normalizer
	synonymGetter
}
