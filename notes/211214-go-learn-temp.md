
### golang
```
2021-12-12 09:13:44 kvogel-surface-ubuntu:~/Downloads 
❯ which go
/usr/local/go/bin/go
2021-12-12 09:13:47 kvogel-surface-ubuntu:~/Downloads 
❯ go version
go version go1.17.5 linux/amd64
```

```
2021-12-12 14:10:51 kvogel-surface-ubuntu:~/projects/tmp 
❯ mkdir hello_go
❯ cd hello_go 
2021-12-12 14:11:04 kvogel-surface-ubuntu:~/projects/tmp/hello_go 
❯ go mod init kvogel.net/hello
go: creating new go.mod: module kvogel.net/hello
❯ ls
go.mod
❯ file go.mod 
go.mod: ASCII text
❯ view go.mod 
```
```go
module kvogel.net/hello

go 1.17
```
```
❯ vi hello.go
```
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
```
❯ go run .
Hello, World!
❯ go build
❯ lr
total 1.7M
-rwxrwxr-x 1 kvogel kvogel 1.7M Dec 12 14:13 hello*
-rw-rw-r-- 1 kvogel kvogel   78 Dec 12 14:12 hello.go
-rw-rw-r-- 1 kvogel kvogel   33 Dec 12 14:11 go.mod
❯ ./hello 
Hello, World!
❯ go list
kvogel.net/hello
❯ go list -f '{{.Target}}'
/home/kvogel/go/bin/hello
❯ ls ~/go/        
pkg/
❯ go install
❯ ls ~/go/  
bin/  pkg/
❯ hello
zsh: command not found: hello
❯ ls ~/go/bin
hello*
❯ vi ~/.profile 
❯ . ~/.profile
❯ hello
Hello, World!
❯ ls /usr/local/go/bin 
go*  gofmt*
```


