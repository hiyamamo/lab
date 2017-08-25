package main

import (
	"github.com/hiyamamo/lab/commands"
	"os"
)

func main() {
	commands.CmdRunner.Execute(os.Args)
}
