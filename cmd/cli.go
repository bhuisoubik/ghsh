package cmd

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/soubikbhuiwk007/ghsh/commands"
	"github.com/soubikbhuiwk007/ghsh/exec"
	"github.com/soubikbhuiwk007/ghsh/vm"
)

func isTokenAuthorised() bool {
	user, _ := user.Current()
	authFname := filepath.Join(user.HomeDir, ".ghsh_auth_token")
	byteToken, err := ioutil.ReadFile(authFname)

	if err != nil || len(byteToken) < 2 {
		return false
	}
	return true
}

func Cli(args []string) {
	if isTokenAuthorised() {
		commands.RegisterAll()
		if len(args) < 2 {
			vm.Shell()
		} else {
			exec.Exec(exec.GetArgs(strings.Join(args[1:], " ")))
		}
	} else {
		if len(args) > 1 {
			if args[1] == "auth" && args[2] == "-login" {
				commands.RegisterAll()
				exec.Exec(exec.GetArgs(strings.Join(args[1:], " ")))
			} else {
				fmt.Println("Not Authorised\nTry running `ghsh auth -login`")
			}
		} else {
			fmt.Println("Not Authorised\nTry running `ghsh auth -login`")
		}
	}
}