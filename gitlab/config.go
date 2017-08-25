package gitlab

import (
	"github.com/mitchellh/go-homedir"
	"path/filepath"
)

var (
	defaultConfigsFile = ""
)

func init() {
	homeDir, _ := homedir.Dir()
	defaultConfigsFile = filepath.Join(homeDir, ".config", "lab")
}

type Config struct {
	Host         string `yaml:"host"`
	PrivateToken string `yaml:"private-token"`
}

var currentConfig *Config
var configLoadedFrom = ""

func CurrentConfig() *Config {
	filename := defaultConfigsFile
	if currentConfig == nil {
		currentConfig = &Config{}
		newConfigService().Load(filename, currentConfig)
		configLoadedFrom = filename
	}

	return currentConfig
}
