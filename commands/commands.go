package commands

import (
	"github.com/urfave/cli"
)

var CmdRunner = newRunner()

type Runner struct {
	commands []cli.Command
}

func newRunner() *Runner {
	return &Runner{
		commands: []cli.Command{},
	}
}

func (r *Runner) Execute(args []string) {
	app := cli.NewApp()

	app.Commands = r.commands
	app.Run(args)
}

func (r *Runner) Use(com cli.Command) {
	r.commands = append(r.commands, com)
}
