package ptoken

import (
	"fmt"
	"os"

	"github.com/soubikbhuiwk007/ghve/reg"
)

func Ptoken(args ...string) {
	fmt.Println(os.Getenv("GHVE_USER_AUTH_TOKEN"))
}

func Register() {
	reg.RegisterNewCommand("get-token", Ptoken)
}