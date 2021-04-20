package exec

import (
	"fmt"

	"github.com/soubikbhuiwk007/ghve/reg"
)

func Exec(args ...string) {
	if _, ok := reg.REGISTERED_COMMANDS[args[0]]; ok {
		reg.REGISTERED_COMMANDS[args[0]](args...)
	} else {
		fmt.Println("Command not Found")
	}
}