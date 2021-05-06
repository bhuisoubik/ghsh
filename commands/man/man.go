package man

import (
	"github.com/soubikbhuiwk007/ghsh/reg"
)

func Man(args []string) {

}

func Register() {
	reg.RegisterNewCommand("man", Man)
}