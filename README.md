# GitHub Shell

<img alt="about" src="docs/images/about.png">

`ghsh` (or Github shell) is a command line tool available for windows, linux and macos that lets you use github as a shell. It is not another `github cli` ofc. It is extensible so you can easily create & add your own commands. You can go into specific folder and edit files right there. It ships as a single executable.

## Installation

### Using Go sdk

```shell
go install github.com/soubikbhuiwk007/ghsh/@latest
```

### Using Package Manager

|OS|Command|
|:-:|:----:|
|`windows`|`choco install ghsh`|
|`darwin`|`brew install ghsh`|
|`linux`|`apt-get install ghsh`|

## Get Started
Run the following to authorise `ghsh`
```shell
$ ghsh auth -login
```

Checkout the list of all the available [Commands](docs/COMMANDS.md) & the [Documentation](docs/README.md)

There are 2 ways to execute commands

* From command line arguments. You can run this from your terminal.

```shell
$ ghsh [command-name] [arguments...]
```

* Inside `ghsh`

```
$ ghsh: @username / > command-name [arguments...]
```

> Some Commands can only be used inside a repository

### **This Project licensed under [MIT](./LICENSE)**
