package gitlab

import (
	"github.com/mitchellh/go-homedir"
	"net/url"
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
	Host         string    `yaml:"host"`
	PrivateToken string    `yaml:"private-token"`
	Projects     []Project `yaml:"projects"`
}

var currentConfig *Config
var configLoadedFrom = ""

func (c *Config) FindProjectFromURL(url *url.URL) (p *Project, err error) {
	p, err = NewProjectFromUrl(url)
	if err != nil {
		return
	}
	proj := c.findProject(p.Owner.Name, p.Name)
	if proj != nil {
		p = proj
	}
	return
}

func (c *Config) findProject(owner, name string) (p *Project) {
	for _, proj := range c.Projects {
		if proj.Owner.Name == owner && proj.Name == name {
			p = &proj
			return
		}
	}
	return
}

func (c *Config) AddProject(p *Project) {
	c.Projects = append(c.Projects, *p)
}

func (c *Config) Save() error {
	filename := defaultConfigsFile
	return newConfigService().Save(filename, c)
}

func CurrentConfig() *Config {
	filename := defaultConfigsFile
	if currentConfig == nil {
		currentConfig = &Config{}
		newConfigService().Load(filename, currentConfig)
		configLoadedFrom = filename
	}

	return currentConfig
}
