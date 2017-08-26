package gitlab

import (
	"fmt"
	"net/url"
	"strings"
)

type Project struct {
	Owner string
	Name  string
}

func (p *Project) UrlString() string {
	return url.PathEscape(fmt.Sprintf("%s/%s", p.Owner, p.Name))
}

func NewProjectFromUrl(url *url.URL) (*Project, error) {
	parts := strings.SplitN(url.Path, "/", 4)
	if len(parts) <= 2 {
		err := fmt.Errorf("Invalid URL: %s", url)
		return nil, err
	}

	name := strings.TrimSuffix(parts[2], ".git")
	p := newProject(parts[1], name, url.Host, url.Scheme)

	return p, nil
}

func newProject(owner, name, host, protocol string) *Project {
	return &Project{
		Owner: owner,
		Name:  name,
	}
}
