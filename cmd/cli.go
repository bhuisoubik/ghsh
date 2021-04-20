package cmd

import (
	"fmt"

	"github.com/soubikbhuiwk007/ghve/auth"
	"github.com/soubikbhuiwk007/ghve/vm"
)

var (
	VERSION string = "1.0.0" 
)

func help() {
	fmt.Println(`Usage: ghve [arguments]
	
Options:
	auth [argument]	:	Authenticate Ghve
	 	-login	:	Login using Personal Token Access
	 	-logout	:	Logout from the previous login
	-v, version	:	Print Current Version
	-h, help	:	Print Help (this)
VE Commands:`)
}
func cmdNotFound(c string){
	fmt.Printf("\nCommand '%v' not found\nTry running 'ghve help'", c)
}

func Cli(args []string) {
	if len(args) < 2 {
		vm.VirtualEnv()
	} else {
		if args[1] == "auth" {
			if args[2] == "-login" {
				auth.Login()
			} else if args[2] == "-logout" {
				auth.Logout()
			} else {
				cmdNotFound(args[2])
			}
		} else if args[1] == "-v" || args[1] == "version" {
			fmt.Printf("\nGhve: Version: %v\n", VERSION)
		} else if args[1] == "-h" || args[1] == "help" {
			help()
		} else {
			cmdNotFound(args[1])
		}
	}
}