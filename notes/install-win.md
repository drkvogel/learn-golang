


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