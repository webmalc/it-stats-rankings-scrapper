package scrappers

// TODO: test it
type pypl struct {
	baseScrapper
}

// scrap scraps the pypl
func (p *pypl) scrap() []*language {
	return []*language{
		p.returnLang("python", 1),
		p.returnLang("go", 2),
	}
}

// newTiobe creates a new tiobe struct
func newPypl() *pypl {
	return &pypl{baseScrapper: baseScrapper{name: "pypl"}}
}
