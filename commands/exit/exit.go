package exit

import (
	"os"

	"github.com/soubikbhuiwk007/ghve/reg"
)

func Exit(args []string) {
	os.Exit(0)
}

func Register() {
	reg.RegisterNewCommand("exit", Exit)
}