package gitlab

import (
	"fmt"
	"net/url"
	"strings"
)

type NameSpace struct {
	Path string
}

type Project struct {
	NameSpace     NameSpace `json:"namespace" yaml:"namespace"`
	Name          string    `json:"name" yaml:"name"`
	DefaultBranch string    `json:"default_branch" yaml:"default-branch"`
}

func (p *Project) String() string {
	return "Onwer: " + p.NameSpace.Path + ", Name: " + p.Name + ", Default Branch: " + p.DefaultBranch
}

func (p *Project) UrlString() string {
	return url.PathEscape(fmt.Sprintf("%s/%s", p.NameSpace.Path, p.Name))
}

func NewProjectFromURL(url *url.URL) (*Project, error) {
	parts := strings.SplitN(url.Path, "/", 4)
	if len(parts) <= 2 {
		err := fmt.Errorf("Invalid URL: %s", url)
		return nil, err
	}

	name := strings.TrimSuffix(parts[2], ".git")
	namespace := NameSpace{
		Path: parts[1],
	}
	p := newProject(namespace, name, "")

	return p, nil
}

func newProject(namespace NameSpace, name, defaultBranch string) *Project {
	return &Project{
		NameSpace:     namespace,
		Name:          name,
		DefaultBranch: defaultBranch,
	}
}
