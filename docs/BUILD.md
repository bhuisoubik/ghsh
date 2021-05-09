## Build the source

### Requirement

* Go sdk 1.16
* Git

### Get the source
```shell
mkdir $GOPATH/src/github.com/soubikbhuiwk007
cd $GOPATH/src/github.com/soubikbhuiwk007
git clone https://www.github.com/soubikbhuiwk007/ghsh
```

### Build

Execute the build script.

* Windows (cmd)
```batch
x\build\build -b
```

* UNIX/Linux (bash)

```shell
./x/build/build.bash -b
```

This will also fetch all the dependencies. If it doesn't, you can run the same command but instead of using `-b` flag, use the `-d` flag.

### Run

**Execute `go run main.go` in the root directory of the project.**

***