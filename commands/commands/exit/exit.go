// Command: exit
// (c) Soubik Bhui <@soubikbhuiwk007> 2021

package exit

import (
	"os"

	"github.com/soubikbhuiwk007/ghsh/reg"
)

func Exit(args []string) {
	os.Exit(0)
}

func Register() {
	reg.RegisterNewCommand("exit", Exit)
}
