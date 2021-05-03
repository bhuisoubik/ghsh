package about

import (
	"fmt"

	"github.com/soubikbhuiwk007/ghsh/reg"
)

func About(args []string) {
	fmt.Println(`
  _____ _ _   _           _        _____ _          _ _ 
 / ____(_) | | |         | |      / ____| |        | | |       Version: ghsh1.0.0
| |  __ _| |_| |__  _   _| |__   | (___ | |__   ___| | |       Build: 1
| | |_ | | __| '_ \| | | | '_ \   \___ \| '_ \ / _ \ | |       License: MIT
| |__| | | |_| | | | |_| | |_) |  ____) | | | |  __/ | |       Author: Soubik Bhui (@soubikbhuiwk007)
 \_____|_|\__|_| |_|\__,_|_.__/  |_____/|_| |_|\___|_|_|`)
}

func Register() {
	reg.RegisterNewCommand("about", About)
}