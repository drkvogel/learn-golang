
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


