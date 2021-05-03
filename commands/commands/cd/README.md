### Command: `cd`

Change directory / repository

#### Usage
`cd <dir_name/repo_name>`

> Only 1 directory can be changed at a time.

* If `cd` command is used outside `repository` then it will move to the `repository` specified in the `argument`
```shell
$ ghsh: @username / > cd <repo-name>
```
* If `cd` command is used inside a `repository` then it will move to the `directory` specified in the `argument`
```shell
$ ghsh: @username /repo/path/ (branch-name) > cd <dir-name>
```

***