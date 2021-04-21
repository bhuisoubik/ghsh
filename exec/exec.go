package exec

import (
	"fmt"

	"github.com/soubikbhuiwk007/ghve/reg"
)

func Exec(c string) {
	args := GetArgs(c)
	if len(args) > 0 {
		_, ok := reg.REGISTERED_COMMANDS[args[0]];
		if ok {
			reg.REGISTERED_COMMANDS[args[0]](args)
		} else {
			fmt.Println("Command not Found")
		}
	}
}