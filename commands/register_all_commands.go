package commands

import (
	"github.com/soubikbhuiwk007/ghve/commands/exit" // command: exit
	"github.com/soubikbhuiwk007/ghve/commands/ptoken" // command: get-token
	"github.com/soubikbhuiwk007/ghve/commands/pwd" // command: pwd
	"github.com/soubikbhuiwk007/ghve/commands/ls" // command: ls <repo|org|gist> <userid>
	"github.com/soubikbhuiwk007/ghve/commands/cd" // command: cd <subdirectory>
	"github.com/soubikbhuiwk007/ghve/commands/dir" // command: dir
	"github.com/soubikbhuiwk007/ghve/commands/mkdir" // command: mkdir <> || <subdirectory>
	"github.com/soubikbhuiwk007/ghve/commands/mkfile" // command: mkfile <> || <filename>
	"github.com/soubikbhuiwk007/ghve/commands/rmdir" // command: rmdir <subdirectory|repository>
	"github.com/soubikbhuiwk007/ghve/commands/rmfile" // command: rmdir <filename|gist-id>
)

func RegisterAll() {
	pwd.Register() // Print Working Directory
	ptoken.Register() // Print Token
	exit.Register() // Exit with Code 0
	ls.Register() // List Organistion/Repository/Gist for Any User
	cd.Register() // Change Directory
	dir.Register() // List Content inside Directory
	mkdir.Register() // Create a new Directory/Repository
	mkfile.Register() // Create a new File/Gist
	rmdir.Register() // Remove an existing Directory/Repository
	rmfile.Register() // Remove an existing File/Gist
}