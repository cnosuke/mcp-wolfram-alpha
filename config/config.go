package config

import (
	"github.com/jinzhu/configor"
)

// Config - Application configuration
type Config struct {
	Log    string `yaml:"log" default:"" env:"LOG_PATH"`
	Debug  bool   `yaml:"debug" default:"false" env:"DEBUG"`
	Wolfram struct {
		AppID           string `yaml:"app_id" env:"WOLFRAM_APP_ID"`
		Timeout         int    `yaml:"timeout" default:"30" env:"WOLFRAM_TIMEOUT"`
		UseBearer       bool   `yaml:"use_bearer" default:"false" env:"WOLFRAM_USE_BEARER"`
		DefaultMaxChars int    `yaml:"default_max_chars" default:"2000" env:"WOLFRAM_DEFAULT_MAX_CHARS"`
	} `yaml:"wolfram"`
}

// LoadConfig - Load configuration file
func LoadConfig(path string) (*Config, error) {
	cfg := &Config{}
	err := configor.New(&configor.Config{
		Debug:      false,
		Verbose:    false,
		Silent:     true,
		AutoReload: false,
	}).Load(cfg, path)
	return cfg, err
}
