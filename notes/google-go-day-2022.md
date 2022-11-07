


[Home - Go Day 2022 on Google Open Source Live ](https://opensourcelive.withgoogle.com/events/go-day-2022)
[Cameron Balahan](https://www.google.com/search?q=Cameron+Balahan&ie=UTF-8)

[Full Event  Go Day on Google Open Source Live](https://www.youtube.com/watch?v=PkCrgxJBmAs)


[Watch - Opening for Go Day 2022 on Google Open Source Live ](https://opensourcelive.withgoogle.com/events/go-day-2022/watch?talk=opening)

#GoogleOSLive ?

### Compatibility

Compatibiity promise
Go is boring, and that's good
Go 1, 2?
```
22/11/2 1:00:47 kvogel@kvogel-macbook:~/p
❯ go version
go version go1.19 darwin/amd64
```
future "point" releases
compatibility at source code level - have to recompile
keep Go boring
API checking - can't take away API 
os.Stdout.WriteString
can't change type
tagged literals, e.g.:
composite literals
Go 1.1 times nanosecond precision
save() and load() - roundtrip sth

why programs 
compatibility problems generally on of: output changes, input changes, protocol changes

sort order - algorithm could change and order equally scored elements differently
gzip output format

deal with range of outputs
or fork code to insulate from

input changes - validate input before calling converters

protocol changes
e.g. HTTP/2
cloud Gophers
can disable HTTP/2 from Go 1.6

Transport.TLSNextProto, ?
no-one should be using SHA1-based certificates
GODEBUG
crypto/TLS
NAT HTTP

Go vet?

### Go Security

vulnerable leaf node dependencies

Go Vulnerability Management
vulnerability database, shared 
pkg.go.dev
CVEs, GHSA, go package maintainers

go vulndb client
pkg.go.dev/vuln
vulncheck API
golang.org/x/vuln/vulncheck
govulncheck command

Finding undiscovered vuln?

### Go Fuzzing


type of automated testing which continously manipulates inputs

seed corpus entries

dividing HTTP service?

he is using vscode

[Go Fuzzing - The Go Programming Language ](https://go.dev/security/fuzz/)

### Go memory limits

Respecting Memory Limits in Go

GOGC - percentage of new heap above live heap allowed before GC kicks in
GC CPU time is almost the same each run and strongly proportional to live heap size
so fewer/less frequent runs of GC reduces CPU usage, but require more memory
tradeoff between CPU and memory - allows you to find the sweet spot for your application

Memory limit - hard limit rather than percentage of live heap?
load spike

thrashing - can cause app to stall, sometimes indefinitely
thrashing mitigation
GC CPU limit

hysterisis?

gc limit good for containers - gc off
not good for e.g. cli or desktop app where you don't control the env

SetMemoryLimit
e.g. when using C in predictable changes in external memory use

https://go.dev/doc/gc-guide

sli.do #82058
