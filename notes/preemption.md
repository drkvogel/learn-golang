
The Go Playground (https://play.golang.org/p/KvN8W_Tupqg)


[non-cooperative preemption golang](https://www.google.com/search?q=non-cooperative+preemption+golang&gs_lcp=Cgdnd3Mtd2l6EAM6BwgAEEcQsAM6BggAEBYQHjoFCCEQoAFQjy5YxDpggz1oAXACeACAAYkBiAHBBZIBAzYuMpgBAKABAaoBB2d3cy13aXrIAQjAAQE&sclient=gws-wiz&uact=5)
[proposal/24543-non-cooperative-preemption.md at master · golang/proposal ](https://github.com/golang/proposal/blob/master/design/24543-non-cooperative-preemption.md)
[Go: Goroutine and Preemption. ℹ️ This article is based on Go 1.13.  by Vincent Blanchon  A Journey With Go  Medium ](https://medium.com/a-journey-with-go/go-goroutine-and-preemption-d6bc2aa2f4b7)
[Go: Asynchronous Preemption. ℹ️ This article is based on Go 1.14.  by Vincent Blanchon  A Journey With Go  Medium ](https://medium.com/a-journey-with-go/go-asynchronous-preemption-b5194227371c)


On Ubuntu/WSL on hp, go 1.13.8:
```
2021-03-30 16:44:01 kvogel-elitebook:~/projects/general/dev/learn/learn-golang ±(master) 
❯ go version  
go version go1.13.8 linux/amd64
❯ go run preemption.go 
2021/03/30 16:56:54 NumCpu 4
Starting Goroutine 1
Starting Goroutine 99
Starting Goroutine 0
Starting Goroutine 36
^Csignal: interrupt     # hangs here
```

```
2021-03-30 16:55:02 kvogel-elitebook:/mnt/c/Users/cbird
❯ htop
12589 kvogel     20   0  100M  5260  1020 R 393.  0.0  5:12.73 preemption
```
393% CPU!

In [The Go Playground ](https://play.golang.org/p/KvN8W_Tupqg):
```
2009/11/10 23:00:00 NumCpu 8
timeout running program
Starting Goroutine 10
Starting Goroutine 64
Starting Goroutine 46
Starting Goroutine 36
Starting Goroutine 51
Starting Goroutine 26
Starting Goroutine 81
Starting Goroutine 21

Program exited: status 1.
```

[Release History - The Go Programming Language ](https://golang.org/doc/devel/release.html)

>go1.16 (released 2021/02/16)
>go1.15 (released 2020/08/11)
>go1.14 (released 2020/02/25)
>go1.13 (released 2019/09/03)
