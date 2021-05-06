### Command: `pr`

Manage Pull Request for current Repository

#### Usage

`pr <command> <argument>`

> Every command should be used inside a repository

#### Commands

* `new <base.ref>` to create a new pull request with the argument supplied as a base.

```shell
$ ghsh: @username repo/path/ (branch) > pr new <base.ref>
```

* `list <state>` to list all the **PR** of a specified state. `state` can only be `closed` , `open`, `all`. `state` must be supplied as an argument.

```shell
$ ghsh: @username repo/path/ (branch) > pr list <state>
```

*  `close <pull-request-number>` to close a specified pull request.

```shell
$ ghsh: @username repo/path/ (branch) > pr close <pr-number>
```

* `open <pull-request-number>` to re-open a specified pull request.

```shell
$ ghsh: @username repo/path/ (branch) > pr open <pr-number>
```

* `state <pr-number>` to get the state (`open` or `closed`) of a specific pr.

```shell
$ ghsh: @username repo/path/ (branch) > pr state <pr-number>
```

* `get <pr-number>` to show a specific pr.

```shell
$ ghsh: @username repo/path/ (branch) > pr get <pr-number>
```

* `merge <pr-number>` to merge a specific pr. It will prompt for a commit message.

```shell
$ ghsh: @username repo/path/ (branch) > pr merge <pr-number>
```

* `updt <command> <pr-number>` to update a pull request.

    * `-t` to update Title
    * `-b` to update Body
    * `-m` to update MaintainerCanModify (Yes | No)
    * `-br` tp update the Base Reference

```shell
$ ghsh: @username repo/path/ (branch) > pr updt <command> <argument>
```

* `diff <pr-number>` tp print diff count.

```shell
$ ghsh: @username repo/path/ (branch) > pr diff <pr-number>
```

* `commits <pr-number>` to print all the commits for a pr.

```shell
$ ghsh: @username repo/path/ (branch) > pr commits <pr-number>
```

***