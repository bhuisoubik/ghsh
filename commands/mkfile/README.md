### Command: `mkfile`

Use the `mkfile` command to create a new Gist/File

#### Usage
`mkfile <argument>`

* If the command is used outside a repository then it will prompt to create a new gist. No arguments needed to supply.
```shell
$ ghsh: @username / > mkfile
```
```
Description:
Type (public|secret):
Path [0]:               # To stop the path prompt press enter at the next path prompt of the last one.
```
* If the command is used inside a repository then a new file will be created. The `file-name` should be supplied as an argument.
```shell
$ ghsh: @username /repo/path > mkfile <file-name> 
```

***