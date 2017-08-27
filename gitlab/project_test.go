package gitlab

import (
	"net/url"
	"testing"

	"github.com/bmizerany/assert"
)

func TestProject_UrlString(t *testing.T) {
	p := &Project{
		Owner: Owner{Name: "Suzukaze-Aoba"},
		Name:  "Fairie's-Story",
	}

	assert.Equal(t, url.PathEscape("Suzukaze-Aoba/Fairie's-Story"), p.UrlString())
}

func TestProject_NewProjectFromURL(t *testing.T) {
	u, _ := url.Parse("ssh://git@gitlab.com/Suzukaze-Aoba/Fairie's-Story.git")
	p, err := NewProjectFromURL(u)

	assert.Equal(t, nil, err)
	assert.Equal(t, "Fairie's-Story", p.Name)
	assert.Equal(t, "Suzukaze-Aoba", p.Owner.Name)
	assert.Equal(t, "", p.DefaultBranch)

	u, _ = url.Parse("git://gitlab.com/Suzukaze-Aoba/Fairie's-Story.git")
	p, err = NewProjectFromURL(u)

	assert.Equal(t, nil, err)
	assert.Equal(t, "Fairie's-Story", p.Name)
	assert.Equal(t, "Suzukaze-Aoba", p.Owner.Name)
	assert.Equal(t, "", p.DefaultBranch)

	u, _ = url.Parse("https://gitlab.com/Suzukaze-Aoba/Fairie's-Story")
	p, err = NewProjectFromURL(u)

	assert.Equal(t, nil, err)
	assert.Equal(t, "Fairie's-Story", p.Name)
	assert.Equal(t, "Suzukaze-Aoba", p.Owner.Name)
	assert.Equal(t, "", p.DefaultBranch)

	u, _ = url.Parse("origin/master")
	_, err = NewProjectFromURL(u)

	assert.NotEqual(t, nil, err)
}
