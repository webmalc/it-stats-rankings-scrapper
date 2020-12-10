package admin

import (
	"github.com/spf13/viper"
)

// Config is the admin configuration struct.
type Config struct {
	AdminURL  string
	AdminPath string
	AdminSSL  bool
}

// setDefaults sets the default values
func setDefaults() {
	viper.SetDefault("admin_url", ":9000")
	viper.SetDefault("admin_path", "")
	viper.SetDefault("admin_ssl", false)
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	setDefaults()
	config := &Config{
		AdminURL:  viper.GetString("admin_url"),
		AdminPath: viper.GetString("admin_path"),
		AdminSSL:  viper.GetBool("admin_ssl"),
	}
	return config
}
