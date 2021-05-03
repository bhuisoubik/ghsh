package cmd

import (
	"strings"

	"github.com/soubikbhuiwk007/ghsh/commands"
	"github.com/soubikbhuiwk007/ghsh/exec"
	"github.com/soubikbhuiwk007/ghsh/vm"
)

func Cli(args []string) {
	commands.RegisterAll()
	if len(args) < 2 {
		vm.Shell()
	} else {
		exec.Exec(exec.GetArgs(strings.Join(args[1:], " ")))
	}
}