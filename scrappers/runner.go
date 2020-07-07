package scrappers

// Runner runs scrappers
type Runner struct{}

// Run runs scrappers
func (r *Runner) Run(names []string) {}

// NewRunner creates a new object.
func NewRunner() *Runner {
	return &Runner{}
}
