### Commands: `issue`

Manage Repository Issues

#### Usage
`issue <command> <argument>`

> Every Command should be used inside a repo

#### Commands

* `close <issue-number>` to close an issue. The `issue-number` should be supplied as an argument.

```shell
$ ghsh: @username /repo/path/ (branch) > issue close <issue-number>
```

* `open <issue-number>` to re-open an issue. The `issue-number` should be supplied as an argument

```shell
$ ghsh: @username /repo/path/ (branch) > issue open <issue-number>
```

* `get <issue-number>` to get a specific issue. The `issue-number` should be supplied as an argument.

```shell
$ ghsh: @username /repo/path/ (branch) > issue get <issue-number>
```

* `state <issue-number>` to get the state (`open` or `closed`)of a specific issue.

```shell
$ ghsh: @username /repo/path/ (branch) > issue state <issue-number>
```

* `list <state(open|closed|all)>` to list all issues of specific repository. `all` will list all every issue (`open` and `closed`).

```shell
$ ghsh: @username /repo/path/ (branch) > issue list <state>
```

* `updt <command> <issue-number>` to update/edit a specific issue. `issue-number` should be supplied as argument.

    * `updt -t <issue-number>` to update the Title
    * `updt -b <issue-number>` to update the Body
    * `updt -l <issue-number>` to update the Labels
    * `updt -a <issue-number>` to update the Assignees

```shell
$ ghsh: @username /repo/path/ (branch) > issue updt <command> <argument>
```

* `new` to create a new Issue.

```
$ ghsh: @username /repo/path/ (branch) > issue new
```

* `lock <issue-number>` to Lock a Issue. It will prompt for Lock-Reason (`off-topic`, `too-heated`, `resolved`, `spam`)

```
$ ghsh: @username /repo/path/ (branch) > issue lock <issue-number>
```

* `comment <command> <argument>` to manage issue comment. Use the following commands:

    * `comment del <comment-id>` to delete a specific comment
    * `comment get <comment-id>` to get a specific comment
    * `comment new <issue-number>` to create a new comment.
    * `comment list <issue-number>` to list all the comments for an issue
    * `comment updt <comment-id>` to update a comment

***