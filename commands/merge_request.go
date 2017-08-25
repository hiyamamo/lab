package commands

import (
	"fmt"
	"github.com/hiyamamo/lab/gitlab"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func init() {
	CmdRunner.Use(mrCommand)
}

var (
	target  string
	project string
	source  string
	title   string
)
var mrCommand = cli.Command{
	Name:    "merge-request",
	Aliases: []string{"mr"},
	Usage:   "create merge-request",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "project, p",
			Usage:       "Project name or id",
			Destination: &project,
		},
		cli.StringFlag{
			Name:        "title, t",
			Usage:       "Title",
			Destination: &title,
		},
		cli.StringFlag{
			Name:        "target, tg",
			Usage:       "Target branch",
			Destination: &target,
		},
		cli.StringFlag{
			Name:        "source, s",
			Usage:       "Source branch",
			Destination: &source,
		},
	},
	Action: action,
}

func action(c *cli.Context) error {
	config := gitlab.CurrentConfig()
	values := url.Values{}
	values.Add("source_branch", source)
	values.Add("target_branch", target)
	values.Add("title", title)
	var uri = config.Host + "/api/v4/projects/" + url.PathEscape(project) + "/merge_requests"
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
	return nil
}
