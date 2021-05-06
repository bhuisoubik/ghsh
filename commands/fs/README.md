### Commands: `fs`

Read / Write Repository files

#### Usage

`fs <command> <argument>..`

#### Commands

* `-r` flag to read a file. `File name` should be supplied as an `argument`

```shell
$ ghsh: @username /repo/path/ (branch) > fs -r <file-name>
```

* `-w` flag to write a file. `File name`, `File path` should be supplied as arguments.

```shell
$ ghsh: @username /repo/path (branch) > fs -w <repo-file-name> <local-file-path>
```

***