// Command: pwd
// (c) Soubik Bhui <@soubikbhuiwk007> 2021

package pwd

import (
	"fmt"

	"github.com/soubikbhuiwk007/ghsh/reg"
	"github.com/soubikbhuiwk007/ghsh/vm/config"
)

func Pwd(args []string) {
	fmt.Println(config.CWD)
}

func Register() {
	reg.RegisterNewCommand("pwd", Pwd)
}
