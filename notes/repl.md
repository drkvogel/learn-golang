


[motemen/gore: Yet another Go REPL that works nicely. Featured with line editing, code completion, and more. ](https://github.com/motemen/gore)

```
21/04/10 18:20:25 homer:~
❯ go get github.com/motemen/gore/cmd/gore
warning: GOPATH set to GOROOT (/usr/local/go) has no effect
package github.com/motemen/gore/cmd/gore: cannot download, $GOPATH must not be set to $GOROOT. For more details see: 'go help gopath'

❯ echo $GOROOT

❯ sudo apt remove golang-go
❯ asdf install golang latest
Platform 'linux' supported!
/home/kvogel/.asdf/downloads/golang/1.16.5/archive.tar.gz: OK
❯ vi ~/.tool-versions
❯ go version
go version go1.16.5 linux/amd64
```
now `gore` works:
```
❯ gore
gore version 0.5.2  :help for help
gore> 
```
