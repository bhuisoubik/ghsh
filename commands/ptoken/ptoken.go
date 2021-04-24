package ptoken

import (
	"fmt"

	"github.com/soubikbhuiwk007/ghve/reg"
	"github.com/soubikbhuiwk007/ghve/vm/config"
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
	reg.RegisterNewCommand("get-token",Ptoken)
}