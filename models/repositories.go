package models

// LanguageRepository is the language repository struct
type LanguageRepository struct {
	db            Database
	nameProcessor nameProcessor
}

// NewLanguage return a service object
func (r *LanguageRepository) NewLanguage(
	source, title string, position uint,
) *Language {
	return &Language{
		Source:        source,
		Title:         title,
		Position:      position,
		nameProcessor: r.nameProcessor,
	}
}

// AddLanguage creates a new language entry in the DB
func (r *LanguageRepository) AddLanguage(
	source, title string, position uint,
) []error {
	l := r.NewLanguage(source, title, position)
	return r.CreateLanguage(l)
}

// CreateLanguage creates a new language
func (r *LanguageRepository) CreateLanguage(lang *Language) []error {
	errors := r.db.Create(lang).GetErrors()
	return errors
}

// NewLanguageRepository return a new LanguageRepository
func NewLanguageRepository(
	database Database, nameProcessor nameProcessor,
) *LanguageRepository {
	return &LanguageRepository{db: database, nameProcessor: nameProcessor}
}
