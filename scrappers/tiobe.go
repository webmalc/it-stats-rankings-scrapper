package scrappers

// TODO: test it
type tiobe struct {
	baseScrapper
}

// scrap scraps the pypl
func (p *tiobe) scrap() []*language {
	return []*language{
		p.returnLang("python", 1),
		p.returnLang("go", 2),
	}
}

// newTiobe creates a new tiobe struct
func newTiobe() *tiobe {
	return &tiobe{baseScrapper: baseScrapper{name: "tiobe"}}
}
