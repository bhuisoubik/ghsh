package vm

import (
	"bufio"
	"fmt"
	"os"

	"github.com/soubikbhuiwk007/ghsh/exec"
	"github.com/soubikbhuiwk007/ghsh/vm/config"
)

func Shell() {
	var rd = bufio.NewReader(os.Stdin)
	for {
		if config.Branch == "" {
			fmt.Print("\n$ ghsh: @" + config.UserName + " " + config.CWD + " > ")
		} else {
			fmt.Print("\n$ ghsh: @" + config.UserName + " " + config.CWD + " (" + config.Branch + ") > ")
		}
		command, _ := rd.ReadString('\n') 
		exec.Exec(exec.GetArgs(command))
	}
}