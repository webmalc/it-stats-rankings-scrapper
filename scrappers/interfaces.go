package scrappers

type langChan = chan []*language

// Logger logs the information
type logger interface {
	Infof(format string, args ...interface{})
	Error(args ...interface{})
}

// languageAdder add a language to the DB
type languageAdder interface {
	AddLanguage(source, title string, position uint) []error
}

// Scrapper scraps the source website
type scrapper interface {
	getName() string
	scrap() []*language
}
