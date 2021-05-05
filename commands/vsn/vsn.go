// Command: vsn
// (c) Soubik Bhui <@soubikbhuiwk007> 2021

package vsn

import (
	"fmt"

	"github.com/soubikbhuiwk007/ghsh/reg"
	"github.com/soubikbhuiwk007/ghsh/version"
)

func Vsn(args []string) {
	fmt.Printf("ghsh: Version: %v\n", version.CurrentVersion)
}

func Register() {
	reg.RegisterNewCommand("vsn", Vsn)
}
