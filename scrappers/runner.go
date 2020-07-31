package scrappers

import "fmt"

// Runner runs scrappers
type Runner struct{}

// Run runs scrappers
func (r *Runner) Run(names []string) {
	fmt.Println(names)
}

// NewRunner creates a new object.
func NewRunner() *Runner {
	return &Runner{}
}
