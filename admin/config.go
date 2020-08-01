package admin

import (
	"github.com/spf13/viper"
)

// Config is the admin configuration struct.
type Config struct {
	AdminURL  string
	AdminPath string
}

// setDefaults sets the default values
func setDefaults() {
	viper.SetDefault("admin_url", ":9000")
	viper.SetDefault("admin_path", "")
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	config := &Config{
		AdminURL:  viper.GetString("admin_url"),
		AdminPath: viper.GetString("admin_path"),
	}
	return config
}
