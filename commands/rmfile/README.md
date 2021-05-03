### Commands: `rmfile`

Delete an existing Gist/File

#### Usage
`rmfile <argument>`

* If the command is used outside a repository then it will delete a gist. The `gist-id` must be supplied as an argument.
```shell
$ ghsh: @username / > rmfile <gist-id>
```
```
Gist of ID '<gist-id>' deleted successfully
```

* If the command is used inside a repository then it will delete a file. The `filename` must be supplied as an argument.
```shell
$ ghsh: @username /repo/path/ (branch) > rmfile <filename>
```
```
File '<filename>' deleted successfully
```

***
