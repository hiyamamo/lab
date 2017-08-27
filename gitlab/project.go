package gitlab

import (
	"fmt"
	"net/url"
	"strings"
)

type Owner struct {
	Name string
}

type Project struct {
	Owner         Owner  `json:"owner" yaml:"owner"`
	Name          string `json:"name" yaml:"name"`
	DefaultBranch string `json:"default_branch" yaml:"default-branch"`
}

func (p *Project) String() string {
	return "Onwer: " + p.Owner.Name + ", Name: " + p.Name + ", Default Branch: " + p.DefaultBranch
}

func (p *Project) UrlString() string {
	return url.PathEscape(fmt.Sprintf("%s/%s", p.Owner.Name, p.Name))
}

func NewProjectFromURL(url *url.URL) (*Project, error) {
	parts := strings.SplitN(url.Path, "/", 4)
	if len(parts) <= 2 {
		err := fmt.Errorf("Invalid URL: %s", url)
		return nil, err
	}

	name := strings.TrimSuffix(parts[2], ".git")
	owner := Owner{
		Name: parts[1],
	}
	p := newProject(owner, name, "")

	return p, nil
}

func newProject(owner Owner, name, defaultBranch string) *Project {
	return &Project{
		Owner:         owner,
		Name:          name,
		DefaultBranch: defaultBranch,
	}
}
