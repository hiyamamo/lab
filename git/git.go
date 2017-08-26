package git

import (
	"fmt"
	"net/url"
	"os/exec"
	"regexp"
	"strings"
)

func RemoteUrl() (u *url.URL, err error) {
	re := regexp.MustCompile(`(.+)\s+(.+)\s+\((push|fetch)\)`)

	rs := runCmd("remote", "-v")

	// build the remotes map
	r := rs[0]
	if re.MatchString(r) {
		match := re.FindStringSubmatch(r)
		rawURL := strings.TrimSpace(match[2])
		u, err = parseURL(rawURL)
	} else {
		err = fmt.Errorf("Invalid git remote %s", r)
	}
	return
}

func parseURL(rawURL string) (u *url.URL, err error) {
	protocolRe := regexp.MustCompile("^[a-zA-Z_+-]+://")
	if !protocolRe.MatchString(rawURL) &&
		strings.Contains(rawURL, ":") &&
		// not a Windows path
		!strings.Contains(rawURL, "\\") {
		rawURL = "ssh://" + strings.Replace(rawURL, ":", "/", 1)
	}
	u, err = url.Parse(rawURL)
	return
}

func CurrentBranch() string {
	outputs := runCmd("rev-parse", "--abbrev-ref", "HEAD")
	return outputs[0]
}

func runCmd(arg ...string) (outputs []string) {
	out, _ := exec.Command("git", arg...).CombinedOutput()
	for _, line := range strings.Split(string(out), "\n") {
		if strings.TrimSpace(line) != "" {
			outputs = append(outputs, string(line))
		}
	}

	return
}
