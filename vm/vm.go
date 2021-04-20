package vm

import (
	"fmt"

	"github.com/soubikbhuiwk007/ghve/commands"
	"github.com/soubikbhuiwk007/ghve/exec"
	"github.com/soubikbhuiwk007/ghve/vm/config"
)

func VirtualEnv() {
	commands.RegisterAll()
	var command string
	pmpt := "\n$ ghve: " + config.UserName + " " + config.CWD + " > "

	for {
		fmt.Print(pmpt)
		fmt.Scanln(&command)
		exec.Exec(command)
	}
}