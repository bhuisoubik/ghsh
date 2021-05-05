package about

import (
	"fmt"

	"github.com/soubikbhuiwk007/ghsh/reg"
	"github.com/soubikbhuiwk007/ghsh/version"
)

func About(args []string) {
	fmt.Printf(`
  _____ _ _   _           _        _____ _          _ _ 
 / ____(_) | | |         | |      / ____| |        | | |       Version: ghsh%v
| |  __ _| |_| |__  _   _| |__   | (___ | |__   ___| | |       Build: 1
| | |_ | | __| '_ \| | | | '_ \   \___ \| '_ \ / _ \ | |       License: MIT
| |__| | | |_| | | | |_| | |_) |  ____) | | | |  __/ | |       Author: Soubik Bhui (@soubikbhuiwk007)
 \_____|_|\__|_| |_|\__,_|_.__/  |_____/|_| |_|\___|_|_|`, version.CurrentVersion)
}

func Register() {
	reg.RegisterNewCommand("about", About)
}