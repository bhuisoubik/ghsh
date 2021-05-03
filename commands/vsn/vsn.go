// Command: vsn
// (c) Soubik Bhui <@soubikbhuiwk007> 2020

package vsn

import (
	"fmt"

	"github.com/soubikbhuiwk007/ghsh/reg"
)

func Vsn(args []string) {
	fmt.Println("ghsh: Version: 1.0.0")
}

func Register() {
	reg.RegisterNewCommand("vsn", Vsn)
}
