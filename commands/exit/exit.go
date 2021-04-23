package exit

import (
	"fmt"
	"os"

	"github.com/soubikbhuiwk007/ghve/reg"
)

func Exit(args []string) {
	os.Exit(0)

	if args[1] == "man" {
		fmt.Println("Exit Environment")
	}
}

func Register() {
	reg.RegisterNewCommand("exit", Exit)
}