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
	fmt.Println("Github Vitual Environment (1.0.0)\nCopyright Soubik Bhui 2021")
	for {
		if config.Branch == "" {
			fmt.Print("\n$ ghve: @" + config.UserName + " " + config.CWD + " > ")
		} else {
			fmt.Print("\n$ ghve: @" + config.UserName + " " + config.CWD + " (" + config.Branch + ") > ")
		}
		command, _ := rd.ReadString('\n') 
		exec.Exec(command)
	}
}