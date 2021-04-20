package crt

import (
	"fmt"

	"github.com/soubikbhuiwk007/ghve/reg"
)

func Crt(args []string) {
	fmt.Println("Create Commmand")
}

func Register() {
	reg.RegisterNewCommand("crt", Crt)
}