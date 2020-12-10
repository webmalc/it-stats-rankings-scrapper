package scrappers

import (
	"context"
	"fmt"
)

// Runner runs scrappers
type Runner struct {
	adder     languageAdder
	logger    logger
	scrappers map[string]scrapper
}

// TODO: make it method
// fanIn sources
func fanIn(inputs ...langChan) langChan {
	c := make(langChan)
	for i := 0; i < len(inputs); i++ {
		in := inputs[i]
		go func() {
			c <- <-in
		}()
	}
	return c
}

// TODO: test it
func (r *Runner) getSelectedScrappers(names []string) []scrapper {
	result := make([]scrapper, 0, len(r.scrappers))
	if len(names) == 0 {
		for _, scrapper := range r.scrappers {
			result = append(result, scrapper)
		}
		return result
	}
	for _, id := range names {
		if scrapper, ok := r.scrappers[id]; ok {
			result = append(result, scrapper)
		}
	}
	return result
}

// TODO: test it
func (r *Runner) getChannels(names []string) []langChan {
	scrappers := r.getSelectedScrappers(names)
	channels := make([]langChan, 0)

	c := make(langChan)
	for _, scrapper := range scrappers {
		scrapper := scrapper
		channels = append(channels, func() langChan {
			go func() {
				c <- scrapper.scrap()
			}()
			return c
		}())
	}
	return channels
}

// save the results
func (r *Runner) save(langs []*language) {
	for _, lang := range langs {
		errors := r.adder.AddLanguage(
			lang.source,
			lang.title,
			lang.position,
		)
		if len(errors) > 0 {
			r.logger.Error(errors)
		}
	}
}

// Run runs scrappers
func (r *Runner) Run(ctx context.Context, names []string) {
	channels := r.getChannels(names)
	c := fanIn(channels...)
	for i := 0; i < len(channels); i++ {
		select {
		case result := <-c:
			fmt.Println(result)
			r.save(result)
		case <-ctx.Done():
			r.logger.Error("The parser timed out")
		}
	}
}

// NewRunner creates a new object.
func NewRunner(adder languageAdder, logger logger) *Runner {
	t := newTiobe()
	p := newPypl()
	return &Runner{
		adder:  adder,
		logger: logger,
		scrappers: map[string]scrapper{
			t.getName(): t,
			p.getName(): p,
		}}
}
