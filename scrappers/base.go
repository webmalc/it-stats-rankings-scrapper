package scrappers

// baseScrapper is the base scrapper struct
type baseScrapper struct {
	name string
}

// returnLang returns a created language structure
func (b *baseScrapper) returnLang(title string, position uint) *language {
	return &language{
		source:   b.getName(),
		title:    title,
		position: position,
	}
}

// getName returns the name
func (b *baseScrapper) getName() string {
	return b.name
}

// language is the language struct
type language struct {
	source   string
	title    string
	position uint
}
