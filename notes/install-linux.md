
[Download and install - The Go Programming Language ](https://golang.org/doc/install)


Install with asdf:
```
2021-06-14 22:53:32 kvogel-surface:~/go
❯ asdf plugin list --urls
elixir                       https://github.com/asdf-vm/asdf-elixir.git
erlang                       https://github.com/asdf-vm/asdf-erlang.git
golang                       https://github.com/kennyp/asdf-golang.git
2021-06-14 22:53:41 kvogel-surface:~/go
❯ asdf list golang
  No versions installed
2021-06-14 22:54:06 kvogel-surface:~/go
❯ asdf install golang latest
Platform 'linux' supported!
/home/kvogel/.asdf/downloads/golang/1.16.5/archive.tar.gz: OK
checksum verified
```

version installed with `apt` overshadows:
```
❯ go version
go version go1.13.8 linux/amd64
❯ which -a go
/home/kvogel/.asdf/shims/go
/usr/bin/go
❯ ~/.asdf/shims/go version
No version set for command go
Consider adding one of the following versions in your config file at /home/kvogel/.tool-versions
golang 1.16.5
```
remove `apt` version:
```
2021-06-14 22:55:06 kvogel-surface:~/go
❯ sudo apt remove golang-go
The following packages were automatically installed and are no longer required:
  golang-1.13-go golang-1.13-race-detector-runtime golang-1.13-src golang-race-detector-runtime golang-src wmdocker
Use 'sudo apt autoremove' to remove them.
The following packages will be REMOVED:
  golang-go
After this operation, 53.2 kB disk space will be freed.
Removing golang-go (2:1.13~1ubuntu2) ...
❯ which go
/home/kvogel/.asdf/shims/go

❯ sudo apt autoremove
The following packages will be REMOVED:
  golang-1.13-go golang-1.13-race-detector-runtime golang-1.13-src golang-race-detector-runtime golang-src wmdocker
After this operation, 324 MB disk space will be freed.
```
"No version set for command go"
```
❯ go
No version set for command go
Consider adding one of the following versions in your config file at /home/kvogel/.tool-versions
golang 1.16.5
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



on hp:

```
20/08/8 1:13:22 kvogel-elitebook:~/p/learn-golang ±(master) 
❯ sudo apt install golang-go
...
❯ which go
/usr/bin/go
```

```
20/08/8 1:16:15 kvogel-elitebook:~/p/learn-golang ±(master) ✗ 
❯ git clone --depth 1 https://github.com/junegunn/fzf.git ~/.fzf
fatal: destination path '/mnt/c/Users/cbird/.fzf' already exists and is not an empty directory.
❯ ls ~/.fzf
BUILD.md*      Dockerfile*  Makefile*       README.md*  doc/     go.sum*   main.go*  plugin/  src/   uninstall*
CHANGELOG.md*  LICENSE*     README-VIM.md*  bin/        go.mod*  install*  man/      shell/   test/
❯ cd ~/.fzf
❯ which fzf
fzf not found
❯ view install 
❯ ./install 
Downloading bin/fzf ...
  - Already exists
  - Checking fzf executable ... 0.18.0
Do you want to enable fuzzy auto-completion? ([y]/n) y
Do you want to enable key bindings? ([y]/n) y
Generate /mnt/c/Users/cbird/.fzf.bash ... OK
Generate /mnt/c/Users/cbird/.fzf.zsh ... OK
Do you want to update your shell configuration files? ([y]/n) y
Update /mnt/c/Users/cbird/.bashrc:
  - [ -f ~/.fzf.bash ] && source ~/.fzf.bash
    - Already exists: line #123 
Update /mnt/c/Users/cbird/.zshrc:
  - [ -f ~/.fzf.zsh ] && source ~/.fzf.zsh
    + Added
Finished. Restart your shell or reload config file.
   source ~/.bashrc  # bash
   source ~/.zshrc   # zsh
Use uninstall script to remove fzf.
For more information, see: https://github.com/junegunn/fzf
❯ . ~/.zshrc 
/mnt/c/Users/cbird/.fzf/shell/key-bindings.zsh:86: parse error near `\n'
❯ which fzf
/mnt/c/Users/cbird/.fzf/bin/fzf
❯ fzf
# works
❯ code ~/.fzf/shell/key-bindings.zsh 
```


[Install Go in Mac - asciinema ](https://asciinema.org/a/138113)
[How to: Install Go 1.11.2 on Ubuntu  by Patrick Dahlke  Medium ](https://medium.com/@patdhlk/how-to-install-go-1-9-1-on-ubuntu-16-04-ee64c073cd79)
[About - asciinema ](https://asciinema.org/about)

[Shell script to download and install latest golang · GitHub ](https://gist.github.com/Zate/b3c8e18cbb2bbac2976d79525d95f893)

