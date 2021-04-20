package vm

import (
	"fmt"
	"time"

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
		arg := GetArgs(command)
		exec.Exec(arg...)
		fmt.Println("Exit Status 0")
		time.Sleep(1e+9)
	}
}