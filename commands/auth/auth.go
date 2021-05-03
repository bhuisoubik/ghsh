// Command: auth
// (c) Soubik Bhui <@soubikbhuiwk007> 2020

package auth

import (
	"fmt"

	au "github.com/soubikbhuiwk007/ghsh/auth"
	"github.com/soubikbhuiwk007/ghsh/reg"
)

func Auth(args []string) {
	if len(args) > 1 {
		if args[1] == "-login" {
			au.Login()
		} else if args[1] == "-logout" {
			au.Logout()
		} else {
			fmt.Println("Invalid Argument")
		}
	} else {
		fmt.Println("Invalid Argument")
	}
}

func Register() {
	reg.RegisterNewCommand("auth", Auth)
}