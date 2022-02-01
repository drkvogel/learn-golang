
# lf

[gokcehan/lf: Terminal file manager ](https://github.com/gokcehan/lf)

installed from [Releases · gokcehan/lf ](https://github.com/gokcehan/lf/releases), put in `~/.local/bin/lf`
shows still `permission denied` in home directory (`~`)
so uninstalled (moved) and installed `lf` with `go install`:

go install github.com/gokcehan/lf@latest
```
2022-01-31 06:16:23 kvogel@kvogel-surface-ubuntu:~ 
❯ go install github.com/gokcehan/lf@latest
```
```
2022-01-31 06:25:05 kvogel@kvogel-surface-ubuntu:~ 
❯ vi ~/.zshrc
❯ . $_
❯ echo $GOPATH 
/home/kvogel/.asdf/installs/golang/1.17.6/packages
❯ which lf
/home/kvogel/.local/bin/lf
❯ mv ~/.local/bin/lf ~/tmp
❯ which lf
/home/kvogel/.asdf/installs/golang/1.17.6/packages/bin/lf
❯ lf
```
still `permission denied` in `~`...
ok, fuck this, lets go and chill out
