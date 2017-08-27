package commands

import (
	"fmt"
	"github.com/hiyamamo/lab/git"
	"github.com/hiyamamo/lab/gitlab"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
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
	project, err := config.FindProjectFromURL(ru)

	errorCheck(err)

	client := gitlab.NewClient(config)
	if target == "" {
		if project.DefaultBranch == "" {
			project, err = client.Project(project.Owner.Name, project.Name)
			errorCheck(err)
			config.AddProject(project)
			config.Save()
		}
		target = project.DefaultBranch
	}

	source := git.CurrentBranch()
	params := map[string]interface{}{
		"source_branch": source,
		"target_branch": target,
		"title":         title,
	}
	res, err := client.CreateMergeRequest(project, params)

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
