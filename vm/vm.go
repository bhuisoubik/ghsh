package vm

import (
	"fmt"
	"os"
	"bufio"

	"github.com/soubikbhuiwk007/ghve/commands"
	"github.com/soubikbhuiwk007/ghve/exec"
	"github.com/soubikbhuiwk007/ghve/vm/config"
)

func VirtualEnv() {
	commands.RegisterAll()
	var rd = bufio.NewReader(os.Stdin)
	pmpt := "\n$ ghve: " + config.UserName + " " + config.CWD + " > "

	for {
		fmt.Print(pmpt)
		command, _ := rd.ReadString('\n') 
		exec.Exec(command)
	}
}