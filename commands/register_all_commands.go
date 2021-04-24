package commands

import (
	"github.com/soubikbhuiwk007/ghve/commands/exit" // Exit
	"github.com/soubikbhuiwk007/ghve/commands/ptoken" // PToken
	"github.com/soubikbhuiwk007/ghve/commands/pwd" // Pwd
	"github.com/soubikbhuiwk007/ghve/commands/ls" // List
	"github.com/soubikbhuiwk007/ghve/commands/cd" // Cd
	"github.com/soubikbhuiwk007/ghve/commands/dir" // Dir
	"github.com/soubikbhuiwk007/ghve/commands/mkdir" // Dir
)

func RegisterAll() {
	pwd.Register() // Print Working Directory
	ptoken.Register() // Print Token
	exit.Register() // Exit with Code 0
	ls.Register() // List Organistion/Repository/Gist for Any User
	cd.Register() // Change Directory
	dir.Register() // List Content inside Directory
	mkdir.Register() // Make Directory
}