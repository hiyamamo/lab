package gitlab

import (
	"fmt"
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
	URL          string    `yaml:"url"`
	PrivateToken string    `yaml:"private-token"`
	Projects     []Project `yaml:"projects"`
}

var currentConfig *Config
var configLoadedFrom = ""

func (c *Config) OpenPrompt() error {
	if c.URL == "" {
		var url, privateToken string
		fmt.Printf("Enter GitLab URL: ")
		fmt.Scanln(&url)
		fmt.Printf("Enter your Private token: ")
		fmt.Scanln(&privateToken)
		c = NewConfig(url, privateToken)
		return c.Save()
	}
	return nil
}

func (c *Config) FindProjectFromURL(url *url.URL) (p *Project, err error) {
	p, err = NewProjectFromURL(url)
	if err != nil {
		return
	}
	proj := c.findProject(p.NameSpace.Path, p.Name)
	if proj != nil {
		p = proj
	}
	return
}

func (c *Config) findProject(namespace, name string) (p *Project) {
	for _, proj := range c.Projects {
		if proj.NameSpace.Path == namespace && proj.Name == name {
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

func NewConfig(url, privateToken string) *Config {
	return &Config{
		URL:          url,
		PrivateToken: privateToken,
	}
}
