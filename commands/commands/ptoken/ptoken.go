// Command: get-token
// (c) Soubik Bhui <@soubikbhuiwk007> 2021

package ptoken

import (
	"fmt"

	"github.com/soubikbhuiwk007/ghsh/reg"
	"github.com/soubikbhuiwk007/ghsh/vm/config"
)

func Ptoken(args []string) {
	token := config.AuthToken
	if token == "" {
		fmt.Println("No Authorised Token")
	} else {
		fmt.Println(token)
	}
}

func Register() {
	reg.RegisterNewCommand("get-token", Ptoken)
}
