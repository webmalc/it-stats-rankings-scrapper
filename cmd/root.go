package cmd

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// CommandRouter is the main commands router.
type CommandRouter struct {
	logger          ErrorLogger
	rootCmd         *cobra.Command
	config          *Config
	scrappersRunner ContextRunner
	adminRunner     Runner
	bindatafsRunner Runner
	timeout         time.Duration
}

// scrap runs scrappers.
func (r *CommandRouter) scrap(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	r.scrappersRunner.Run(ctx, args)
}

// admin runs admin server.
func (r *CommandRouter) admin(cmd *cobra.Command, args []string) {
	r.adminRunner.Run(args)
}

// bindatafs runs bindatafs generator.
func (r *CommandRouter) bindatafs(cmd *cobra.Command, args []string) {
	r.bindatafsRunner.Run(args)
}

// Run the router.
func (r *CommandRouter) Run() {
	scrapCommand := &cobra.Command{
		Use:       "scrap",
		Short:     "Run the scrappers",
		Run:       r.scrap,
		ValidArgs: r.config.scrappers,
		Args:      cobra.OnlyValidArgs,
	}
	scrapCommand.Flags().DurationVar(
		&r.timeout, "timeout", r.config.timeout, "timeout for scrapping",
	)
	r.rootCmd.AddCommand(
		scrapCommand,
		&cobra.Command{
			Use:   "admin",
			Short: "Run the admin",
			Run:   r.admin,
		},
		&cobra.Command{
			Use:   "bindatafs",
			Short: "Run the bindatafs generator",
			Run:   r.bindatafs,
		},
	)
	err := r.rootCmd.Execute()
	if err != nil {
		r.logger.Error(errors.Wrap(err, "root command"))
	}
}

// NewCommandRouter creates a new CommandRouter.
func NewCommandRouter(
	log ErrorLogger, admin, bindata Runner, scrapper ContextRunner,
) CommandRouter {
	config := NewConfig()
	return CommandRouter{
		config:          config,
		logger:          log,
		rootCmd:         &cobra.Command{Use: "its-rankings.app"},
		scrappersRunner: scrapper,
		adminRunner:     admin,
		bindatafsRunner: bindata,
		timeout:         config.timeout,
	}
}
