
### GOPATH

[GOPATH Explained ](https://flaviocopes.com/go-gopath/)

>The `GOPATH` environment variable specifies the location of your workspace. That’s where you find the Go tools, where you develop and where you install 3rd party packages and binaries. As of *Go 1.8*, if you don’t set a `GOPATH`, the default will be used. In older releases had to set it explicitly, but for ease of use, a default has been introduced. By default the value of `GOPATH` is `$HOME/go` on Unix-like systems

```
2021-06-13 14:57:25 kvogel-surface:~
❯ go version
go version go1.13.8 linux/amd64
❯ echo $GOPATH
/home/kvogel/go
❯ go env GOPATH
/home/kvogel/go
```

[Understanding the GOPATH  DigitalOcean ](https://www.digitalocean.com/community/tutorials/understanding-the-gopath)

[gp - The GOPATH Manager : golang ](https://www.reddit.com/r/golang/comments/691fyl/gp_the_gopath_manager/)


### GOROOT

>`$GOPATH` Is Not `$GOROOT`. The `$GOROOT` is where Go’s code, compiler, and tooling lives — this is not our source code. The `$GOROOT` is usually something like `/usr/local/go`. Our `$GOPATH` is usually something like `$HOME/go`.

```
2021-06-13 15:01:32 kvogel-surface:~
❯ echo $GOROOT
❯ go env GOROOT
/usr/lib/go-1.13
```

[golang no src dir](https://www.google.com/search?q=golang+no+src+dir&ie=UTF-8)
[No src folder although the official guide tells to? · Issue #48 · golang-standards/project-layout ](https://github.com/golang-standards/project-layout/issues/48)
[go - src/golang.org directory missing](https://stackoverflow.com/questions/46926380/src-golang-org-directory-missing)
[Why Go put all code in $GOPATH/src](https://stackoverflow.com/questions/38987866/why-go-put-all-code-in-gopath-src)
>Golang is going to introduce an official tool called `dep`, which is like `pip` for python or `npm` for node.js.



```
21/04/10 18:20:15 homer:~
❯ which go
/usr/local/go/bin/go
❯ vi ~/.zshrc
❯ . ~/.zshrc
❯ echo $GOPATH
/usr/local/go
```


```
21/04/10 18:20:25 homer:~
❯ go get github.com/motemen/gore/cmd/gore
warning: GOPATH set to GOROOT (/usr/local/go) has no effect
package github.com/motemen/gore/cmd/gore: cannot download, $GOPATH must not be set to $GOROOT. For more details see: 'go help gopath'

❯ echo $GOROOT

❯ go help gopath
```
>The Go path is used to resolve import statements. It is implemented by and documented in the go/build package.
>The GOPATH environment variable lists places to look for Go code. On Unix, the value is a colon-separated string. On Windows, the value is a semicolon-separated string. On Plan 9, the value is a list. If the environment variable is unset, GOPATH defaults to a subdirectory named "go" in the user's home directory ($HOME/go on Unix, %USERPROFILE%\go on Windows), unless that directory holds a Go distribution. Run "go env GOPATH" to see the current GOPATH.
>See https://golang.org/wiki/SettingGOPATH to set a custom GOPATH.
>Each directory listed in GOPATH must have a prescribed structure:
>The `src` directory holds source code. The path below src determines the import path or executable name.
>The `pkg` directory holds installed package objects. As in the Go tree, each target operating system and architecture pair has its own subdirectory of pkg (`pkg/GOOS_GOARCH`).
>If `DIR` is a directory listed in the `GOPATH`, a package with source in `DIR/src/foo/bar` can be imported as `foo/bar` and has its compiled form installed to `DIR/pkg/GOOS_GOARCH/foo/bar.a`
>The `bin` directory holds compiled commands. Each command is named for its source directory, but only the final element, not the entire path. That is, the command with source in DIR/src/foo/quux is installed into DIR/bin/quux, not DIR/bin/foo/quux. The "foo/" prefix is stripped so that you can add DIR/bin to your PATH to get at the installed commands. If the GOBIN environment variable is set, commands are installed to the directory it names instead of DIR/bin. GOBIN must be an absolute path.
>Here's an example directory layout:
```
    GOPATH=/home/user/go

    /home/user/go/
        src/
            foo/
                bar/               (go code in package bar)
                    x.go
                quux/              (go code in package main)
                    y.go
        bin/
            quux                   (installed command)
        pkg/
            linux_amd64/
                foo/
                    bar.a          (installed package object)
```
>Go searches each directory listed in GOPATH to find source code, but new packages are always downloaded into the first directory in the list. See https://golang.org/doc/code.html for an example.

### GOPATH and Modules
>*When using modules, `GOPATH` is no longer used* for resolving imports. However, it is still used to store downloaded source code (in `GOPATH/pkg/mod`) and compiled commands (in `GOPATH/bin`).

### Internal Directories
>Code in or below a directory named "internal" is importable only by code in the directory tree rooted at the parent of "internal". Here's an extended version of the directory layout above:
```
    /home/user/go/
        src/
            crash/
                bang/              (go code in package bang)
                    b.go
            foo/                   (go code in package foo)
                f.go
                bar/               (go code in package bar)
                    x.go
                internal/
                    baz/           (go code in package baz)
                        z.go
                quux/              (go code in package main)
                    y.go
```

>The code in z.go is imported as "foo/internal/baz", but that import statement can only appear in source files in the subtree rooted at foo. The source files foo/f.go, foo/bar/x.go, and foo/quux/y.go can all import "foo/internal/baz", but the source file crash/bang/b.go cannot. See https://golang.org/s/go14internal for details.

### Vendor Directories
>Go 1.6 includes support for using local copies of external dependencies to satisfy imports of those dependencies, often referred to as vendoring. Code below a directory named "vendor" is importable only by code in the directory tree rooted at the parent of "vendor", and only using an import path that omits the prefix up to and including the vendor element. Here's the example from the previous section, but with the "internal" directory renamed to "vendor" and a new foo/vendor/crash/bang directory added:
```
    /home/user/go/
        src/
            crash/
                bang/              (go code in package bang)
                    b.go
            foo/                   (go code in package foo)
                f.go
                bar/               (go code in package bar)
                    x.go
                vendor/
                    crash/
                        bang/      (go code in package bang)
                            b.go
                    baz/           (go code in package baz)
                        z.go
                quux/              (go code in package main)
                    y.go
```
>The same visibility rules apply as for internal, but the code in z.go is imported as "baz", not as "foo/vendor/baz".
>Code in vendor directories deeper in the source tree shadows code in higher directories. Within the subtree rooted at foo, an import of "crash/bang" resolves to "foo/vendor/crash/bang", not the top-level "crash/bang".
>Code in vendor directories is not subject to import path checking (see 'go help importpath').
>When 'go get' checks out or updates a git repository, it now also updates submodules.
>Vendor directories do not affect the placement of new repositories being checked out for the first time by 'go get': those are always placed in the main GOPATH, never in a vendor subtree. See https://golang.org/s/go15vendor for details.


[motemen/gore: Yet another Go REPL that works nicely. Featured with line editing, code completion, and more. ](https://github.com/motemen/gore)
[cho45/KeyCast: Record keystroke for screencast ](https://github.com/cho45/KeyCast)
[SettingGOPATH · golang/go Wiki ](https://github.com/golang/go/wiki/SettingGOPATH)
[bash_profile zsh](https://www.google.com/search?q=bash_profile+zsh)

```
21/04/10 18:26:07 homer:~
❯ ls go
bin/  pkg/
❯ vi ~/.zshrc
```
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
```
21/04/10 18:28:00 homer:~
❯ . ~/.zshrc
❯ echo $GOROOT

❯ echo $GOPATH
/home/kvogel/go
❯ go get github.com/motemen/gore/cmd/gore
...
21/04/10 18:34:47 homer:~
❯ gore
gore version 0.5.2  :help for help
gore> :import fmt
gore>
gore> fmt.Println("Hello, World!")
Hello, World!
14
nil
```

2022-01-31 05:50 Mon / 11:20 IST / 12:50 ICT

### GOPATH confusion

trying to install `lf` file manager got errors
```
2022-01-31 02:23:20 kvogel@kvogel-surface-ubuntu:~/go 
❯ go version
zsh: command not found: go
❯ which go
go not found

2022-01-31 02:35:20 kvogel@kvogel-surface-ubuntu:~ 
❯ ls /usr/local/go/bin 
go*  gofmt*
❯ ls go/bin 
gopls*  hello*
❯ vi ~/.profile
```
```diff
 export PATH=$PATH:$HOME/go/bin/
+export PATH=$PATH:/usr/local/go/bin
```
??
OK, reinstall via `asdf`
```
2022-01-31 03:43:31 kvogel@kvogel-surface-ubuntu:~/Downloads 
❯ asdf plugin add golang
❯ asdf install golang latest
Platform 'linux' supported!
/home/kvogel/.asdf/downloads/golang/1.17.6/archive.tar.gz: OK
❯ which go
/usr/local/go/bin/go
❯ which -a -- go
/home/kvogel/.asdf/shims/go
/usr/local/go/bin/go
❯ du -sh /usr/local/go 
423M	/usr/local/go

❯ sudo rm -rf /usr/local/go
❯ sudo rm -rf ~/go
❯ lf
❯ go
No version set for command go
Consider adding one of the following versions in your config file at /home/kvogel/.tool-versions
golang 1.17.6
❯ asdf global golang latest
❯ go
Go is a tool for managing Go source code.
...
```
