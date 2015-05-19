#Got

**Git + Go = Got**

Simple git utility functions.

##Install

```bash
$ go get github.com/morcmarc/got
```

##Documentation

[![GoDoc](https://godoc.org/github.com/morcmarc/got?status.svg)](https://godoc.org/github.com/morcmarc/got)

##Supported Commands (Work in Progress)

- `git init`
- `git clone`
- `git remote {add, --get, set-url}`

##Example

```go
import (
    "fmt"
    "github.com/morcmarc/got"
)

// Use Interface to make testing / mocking easier
var gotCli got.Interface

// Initialize a new GotClient at the given path
gotCli = got.NewGotClient("~/Code/got")

// Clone a repository into ~/Code/got
err := gotCli.Clone("git@github.com:morcmarc/got.git")
// ... handle errors

// Update origin to use bitbucket instead
err := gotCli.SetRemote("origin", "git@bitbucket.org:morcmarc/got.git")
// ... handle errors

// Confirm changes
remote, err := gotCli.GetRemote("origin")
// ... handle errors
fmt.Println(remote)
```

##Unit Tests

```bash
$ make test
```