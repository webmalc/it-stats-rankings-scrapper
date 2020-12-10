package cmd

import (
	"time"

	"github.com/spf13/viper"
)

// Config is the cmd configuration struct.
type Config struct {
	scrappers []string
	timeout   time.Duration
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	config := &Config{
		scrappers: viper.GetStringSlice("scrappers"),
		timeout:   10 * time.Minute,
	}
	return config
}
