package ptoken

import (
	"fmt"
	"io/ioutil"

	"github.com/soubikbhuiwk007/ghve/reg"
)

func Ptoken(args []string) {
	if args[1] == "man" {
		fmt.Println("Print Authorised token")
	} else {
		tok, _ := ioutil.ReadFile("auth/.gh_access_token")
		if string(tok) == "" {
			fmt.Println("No Authorised Token")
		} else {
			fmt.Println(string(tok))
		}
	}
}

func Register() {
	reg.RegisterNewCommand("get-token",Ptoken)
}