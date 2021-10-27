
[How to Write Go Code - The Go Programming Language ](https://golang.org/doc/code)


```
2021-06-14 23:07:54 kvogel-surface:~/go
❯ go version
go version go1.16.5 linux/amd64
❯ z learn-go
2021-06-14 23:24:00 kvogel-surface:~/projects/general/dev/learn/learn-golang ±(master) ✗
❯ cd tutorials
❯ mkdir first
❯ cd first
❯ mkdir hello
❯ cd $_
❯ go version
go version go1.16.5 linux/amd64
❯ go mod init example.com/user/hello
go: creating new go.mod: module example.com/user/hello
❯ cat go.mod
module example.com/user/hello

go 1.16
❯ ls
go.mod
❯ vi hello.go
❯ go install example.com/user/hello
❯ which hello
/home/kvogel/go/bin/hello
❯ hello
Hello, world.
```

>The install directory is controlled by the GOPATH and GOBIN environment variables. If GOBIN is set, binaries are installed to that directory. If GOPATH is set, binaries are installed to the bin subdirectory of the first directory in the GOPATH list. Otherwise, binaries are installed to the bin subdirectory of the default GOPATH ($HOME/go or %USERPROFILE%\go).

try install when not in dir with `.mod` file:
```
2021-06-14 23:27:50 kvogel-surface:~/projects/general/dev/learn/learn-golang/tutorials/first/hello ±(master) ✗
❯ cd ..
❯ go install example.com/user/hello
go install: version is required when current directory is not in a module
        Try 'go install example.com/user/hello@latest' to install the latest version
```
try installing `@latest`:
```
2021-06-14 23:29:36 kvogel-surface:~/projects/general/dev/learn/learn-golang/tutorials/first/hello ±(master) ✗
❯ go install example.com/user/hello@latest
go install example.com/user/hello@latest: unrecognized import path "example.com/user/hello": https fetch: Get "https://example.com/user/hello?go-get=1": dial tcp: lookup example.com on 172.26.16.1:53: no such host
```


>You can use the `go env` command to portably set the default value for an environment variable for future go commands:
```
$ go env -w GOBIN=/somewhere/else/bin
```
>To unset a variable previously set by go env -w, use go env -u:
```
$ go env -u GOBIN
```
>For convenience, go commands accept paths relative to the working directory, and default to the package in the current working directory if no other path is given. So in our working directory, the following commands are all equivalent:
```
$ go install example.com/user/hello
$ go install .
$ go install
```

```
$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))    # ???
$ hello
Hello, world.
```
?

>Because our `ReverseRunes` function *begins with an upper-case letter, it is exported*, and can be used in other packages that import our `morestrings` package.

```
$ go build
```
>This won't produce an output file. Instead it saves the compiled package in the local build cache.

```
2021-06-14 23:48:49 kvogel-surface:~/projects/general/dev/learn/learn-golang/tutorials/first/hello ±(master) ✗
❯ go run example.com/user/hello
.dlrow ,olleH
```

>Now that you have a dependency on an external module, you need to download that module and record its version in your go.mod file. The go mod tidy command adds missing module requirements for imported packages and removes requirements on modules that aren't used anymore.

```
2021-06-14 23:53:04 kvogel-surface:~/projects/general/dev/learn/learn-golang/tutorials/first/hello ±(master) ✗
❯ go run example.com/user/hello
hello.go:6:2: no required module provides package github.com/google/go-cmp/cmp; to add it:
        go get github.com/google/go-cmp/cmp
❯ go install example.com/user/hello
hello.go:6:2: no required module provides package github.com/google/go-cmp/cmp; to add it:
        go get github.com/google/go-cmp/cmp
❯ go mod tidy
go: finding module for package github.com/google/go-cmp/cmp
go: downloading github.com/google/go-cmp v0.5.6
go: found github.com/google/go-cmp/cmp in github.com/google/go-cmp v0.5.6
go: downloading golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
❯ go run example.com/user/hello
Hello, Go!
  string(
-       "Hello, World",
+       "Hello Go",
  )
```

>Module dependencies are automatically downloaded to the pkg/mod subdirectory of the directory indicated by the GOPATH environment variable. The downloaded contents for a given version of a module are shared among all other modules that require that version, so the go command marks those files and directories as read-only. To remove all downloaded modules, you can pass the -modcache flag to go clean:
```
$ go clean -modcache
```

### Testing

>Go has a lightweight test framework composed of the `go test` command and the `testing` package.
>You write a test by creating a file with a name ending in `_test.go` that contains functions named `TestXXX` with signature `func (t *testing.T)`. The test framework runs each such function; if the function calls a failure function such as `t.Error` or `t.Fail`, the test is considered to have failed.

```
2021-06-15 00:27:50 kvogel-surface:~/projects/general/dev/learn/learn-golang/tutorials/first/hello/morestrings ±(master) ✗
❯ go test
PASS
ok      example.com/user/hello/morestrings      0.002s
```


### What's next

>Subscribe to the `golang-announce` mailing list to be notified when a new stable version of Go is released.
>See [Effective Go](https://golang.org/doc/effective_go.html) for tips on writing clear, idiomatic Go code.
>Take [A Tour of Go](https://tour.golang.org/) to learn the language proper.
>Visit the [documentation page](https://golang.org/doc/#articles) for a set of in-depth articles about the Go language and its libraries and tools.

### Getting help

>For real-time help, ask the helpful gophers in the community-run gophers Slack server (grab an invite here).
>The official mailing list for discussion of the Go language is Go Nuts.
>Report bugs using the Go issue tracker.
