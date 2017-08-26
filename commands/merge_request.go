package commands

import (
	"fmt"
	"github.com/hiyamamo/lab/git"
	"github.com/hiyamamo/lab/gitlab"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func init() {
	CmdRunner.Use(mrCommand)
}

var (
	target string
	source string
	title  string
)
var mrCommand = cli.Command{
	Name:    "merge-request",
	Aliases: []string{"mr"},
	Usage:   "create merge-request",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "title, t",
			Usage:       "Title",
			Destination: &title,
		},
		cli.StringFlag{
			Name:        "target, T",
			Usage:       "Target branch",
			Destination: &target,
		},
	},
	Action: action,
}

func action(c *cli.Context) error {
	config := gitlab.CurrentConfig()
	ru, err := git.RemoteUrl()
	errorCheck(err)
	project, _ := gitlab.NewProjectFromUrl(ru)
	source := git.CurrentBranch()
	values := url.Values{}
	values.Add("source_branch", source)
	values.Add("target_branch", target)
	values.Add("title", title)
	var uri = config.Host + "/api/v4/projects/" + project.UrlString() + "/merge_requests"
	req, err := http.NewRequest("POST", uri, strings.NewReader(values.Encode()))
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("PRIVATE-Token", config.PrivateToken)
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(res.Status)
	if res.Status != "201 Created" {
		body, _ := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		fmt.Printf("%s", body)
		return fmt.Errorf("%s", body)
	}
	return nil
}

func errorCheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
