

[The Wails Project  Wails ](https://wails.io/)

>Wails is a project that enables you to write desktop apps using Go and web technologies. Consider it a lightweight and fast Electron alternative for Go. You can easily build applications with the flexibility and power of Go, combined with a rich, modern frontend.




```
22/10/24 3:37:32 kvogel@kvogel-macbook:~/Projects/general/dev/learn/learn-golang/wails ±(master) ✗ 
❯ wails init -n svelte-ts -t svelte-ts
Wails CLI v2.1.0


22/10/24 3:40:43 kvogel@kvogel-macbook:~/Projects/general/dev/learn/learn-golang/wails/svelte-ts ±(master) ✗ 
❯ wails dev
Wails CLI v2.1.0

Executing: go mod tidy
  - Generating bindings: Done.
  - Installing frontend dependencies: Done.
  - Compiling frontend: Done.

> frontend@0.0.0 dev /Users/kvogel/Projects/general/dev/learn/learn-golang/wails/svelte-ts/frontend
> vite

Vite Server URL: http://127.0.0.1:5173/

  VITE v3.1.8  ready in 809 ms

  ➜  Local:   http://127.0.0.1:5173/
  ➜  Network: use --host to expose
Running frontend DevWatcher command: 'npm run dev'
Building application for development...
  - Generating bindings: Done.
  - Compiling application: Done.
  - Packaging application: Done.

INF | Serving assets from frontend DevServer URL: http://127.0.0.1:5173/
DEB | [DevWebServer] Serving DevServer at http://localhost:34115
watching: /Users/kvogel/Projects/general/dev/learn/learn-golang/wails/svelte-ts
...
```



[UPX: the Ultimate Packer for eXecutables - Homepage ](https://upx.github.io/)


[Installation  Wails ](https://wails.io/docs/gettingstarted/installation)

```
22/10/24 3:28:34 kvogel@kvogel-macbook:~/Projects/general/dev/learn/learn-golang ±(master) 
❯ go install github.com/wailsapp/wails/v2/cmd/wails@latest

go: downloading github.com/yuin/goldmark-emoji v1.0.1
go: downloading github.com/go-git/go-billy/v5 v5.2.0
go: downloading github.com/imdario/mergo v0.3.12
go: downloading golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
go: downloading github.com/rivo/uniseg v0.2.0
go: downloading github.com/alecthomas/chroma v0.10.0
go: downloading github.com/microcosm-cc/bluemonday v1.0.17
go: downloading github.com/muesli/reflow v0.3.0
go: downloading github.com/tidwall/pretty v1.1.0
go: downloading github.com/tidwall/match v1.0.3
go: downloading github.com/lucasb-eyer/go-colorful v1.2.0
go: downloading github.com/mitchellh/go-homedir v1.1.0
go: downloading github.com/emirpasic/gods v1.12.0
go: downloading github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99
go: downloading github.com/aymerick/douceur v0.2.0
go: downloading golang.org/x/net v0.0.0-20220722155237-a158d28d115b
go: downloading github.com/go-git/gcfg v1.5.0
go: downloading github.com/kevinburke/ssh_config v0.0.0-20201106050909-4977a11b4351
go: downloading github.com/xanzy/ssh-agent v0.3.0
go: downloading github.com/dlclark/regexp2 v1.4.0
go: downloading github.com/gorilla/css v1.0.0
go: downloading gopkg.in/warnings.v0 v0.1.2

❯ wails doctor
Wails CLI v2.1.0
...
Wails
------
Version:        v2.1.0

Dependency                      Package Name    Status          Version
----------                      ------------    ------          -------
xcode command line tools        N/A             Installed       2354
npm                             N/A             Installed       6.14.8
*upx                            N/A             Available 
*nsis                           N/A             Available 
* - Optional Dependency

Diagnosis
---------
Your system is ready for Wails development!
Optional package(s) installation details: 
  - upx : Available at https://upx.github.io/
  - nsis : Available at https://nsis.sourceforge.io/Download
```
