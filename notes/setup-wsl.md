
[Download and install - The Go Programming Language ](https://golang.org/doc/install)
[Install Golang on WSL - sal.as ](https://sal.as/post/install-golan-on-wsl/)

[command not found after go build](https://stackoverflow.com/questions/39816579/command-not-found-after-go-build)

```
21/02/2 19:29:03 kvogel-elitebook:~/go
❯ sudo mkdir -p /usr/local/go/bin
[sudo] password for kvogel:
21/02/2 19:29:16 kvogel-elitebook:~/go
❯ vi ~/.zshrc
```
```bash
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go;
export PATH=$PATH:$GOPATH/bin;
```
```
21/02/2 19:30:12 kvogel-elitebook:~/go
❯ vi ~/.zshrc && source $_
This is /home/kvogel/.zshrc
EOF (ctrl+d) is disabled
ctrl+r is ignored
21/02/2 19:30:27 kvogel-elitebook:~/go
❯ echo $GOPATH/
/home/kvogel/go/
21/02/2 19:31:15 kvogel-elitebook:~/go
❯ ls $GOPATH/
pkg/
```

```
21/02/2 19:33:52 kvogel-elitebook:~/po/duf ±(master) ✗
❯ sudo cp duf /usr/local/go/bin
❯ duf
╭──────────────────────────────────────────────────────────────────────────────────────────────────────────────────╮
│ 16 local devices                                                                                                 │
├────────────────────────────────┬────────┬────────┬────────┬───────────────────────────────┬─────────┬────────────┤
│ MOUNTED ON                     │   SIZE │   USED │  AVAIL │              USE%             │ TYPE    │ FILESYSTEM │
├────────────────────────────────┼────────┼────────┼────────┼───────────────────────────────┼─────────┼────────────┤
│ /                              │ 251.0G │   7.5G │ 230.7G │ [....................]   3.0% │ ext4    │ /dev/sdb   │
│ /init                          │ 221.3G │ 201.8G │  19.4G │ [##################..]  91.2% │ 9p      │ tools      │
│ /mnt/c                         │ 221.3G │ 201.8G │  19.4G │ [##################..]  91.2% │ 9p      │ C:\        │
│ /mnt/d                         │  13.3G │   2.0G │  11.4G │ [##..................]  14.7% │ 9p      │ D:\        │
│ /mnt/e                         │   2.0G │  72.6M │   1.9G │ [....................]   3.6% │ 9p      │ E:\        │
│ /mnt/g                         │ 465.7G │ 318.7G │ 147.1G │ [#############.......]  68.4% │ 9p      │ G:\        │
│ /mnt/wsl/docker-desktop-bind-m │ 251.0G │   7.5G │ 230.7G │ [....................]   3.0% │ ext4    │ /dev/sdb   │
│ ounts/Ubuntu/19f074b726d2257c2 │        │        │        │                               │         │            │
│ 7e6915e9f61956d079846330351bd2 │        │        │        │                               │         │            │
│ 4e6bf447018de5191              │        │        │        │                               │         │            │
│ /mnt/wsl/docker-desktop-bind-m │ 251.0G │   7.5G │ 230.7G │ [....................]   3.0% │ ext4    │ /dev/sdb   │
│ ounts/Ubuntu/1dbf25b0c3439176d │        │        │        │                               │         │            │
│ a674412e4ef4df0c22d335d385bfc4 │        │        │        │                               │         │            │
│ d5be847bc780e3eee              │        │        │        │                               │         │            │
│ /mnt/wsl/docker-desktop-bind-m │ 251.0G │   7.5G │ 230.7G │ [....................]   3.0% │ ext4    │ /dev/sdb   │
│ ounts/Ubuntu/52c1b51354eb86437 │        │        │        │                               │         │            │
│ fd784fb96fa86f14b19759860bfa01 │        │        │        │                               │         │            │
│ 9b30ee700c74962e9              │        │        │        │                               │         │            │
│ /mnt/wsl/docker-desktop-bind-m │ 251.0G │   7.5G │ 230.7G │ [....................]   3.0% │ ext4    │ /dev/sdb   │
│ ounts/Ubuntu/78a2a4cb466ed4f14 │        │        │        │                               │         │            │
│ 9efbd79c2eb94d1226d7be5c50a906 │        │        │        │                               │         │            │
│ 02a8f00f560ca2972              │        │        │        │                               │         │            │
│ /mnt/wsl/docker-desktop-bind-m │ 251.0G │   7.5G │ 230.7G │ [....................]   3.0% │ ext4    │ /dev/sdb   │
│ ounts/Ubuntu/9896cb67c9a0ef971 │        │        │        │                               │         │            │
│ f09f4361566e0cf2ecfa64e8301a3c │        │        │        │                               │         │            │
│ 351ed03cb36a68067              │        │        │        │                               │         │            │
│ /mnt/wsl/docker-desktop-bind-m │ 251.0G │   7.5G │ 230.7G │ [....................]   3.0% │ ext4    │ /dev/sdb   │
│ ounts/Ubuntu/ca9ffd7bb50477662 │        │        │        │                               │         │            │
│ ad177dad057ebbff1bcd2a44cca739 │        │        │        │                               │         │            │
│ a495e3fe8f41648b3              │        │        │        │                               │         │            │
│ /mnt/wsl/docker-desktop-bind-m │ 251.0G │   7.5G │ 230.7G │ [....................]   3.0% │ ext4    │ /dev/sdb   │
│ ounts/Ubuntu/ed4b20127bcffe750 │        │        │        │                               │         │            │
│ 0d540d3824872cb0e446e4768d97a4 │        │        │        │                               │         │            │
│ 1b7cef01cf9026d7c              │        │        │        │                               │         │            │
│ /mnt/wsl/docker-desktop-data/i │ 251.0G │   7.6G │ 230.5G │ [....................]   3.0% │ ext4    │ /dev/sdd   │
│ socache                        │        │        │        │                               │         │            │
│ /mnt/wsl/docker-desktop/cli-to │ 383.3M │ 383.3M │     0B │ [####################] 100.0% │ iso9660 │ /dev/loop0 │
│ ols                            │        │        │        │                               │         │            │
│ /mnt/wsl/docker-desktop/docker │ 251.0G │ 125.8M │ 238.0G │ [....................]   0.0% │ ext4    │ /dev/sdc   │
│ -desktop-proxy                 │        │        │        │                               │         │            │
╰────────────────────────────────┴────────┴────────┴────────┴───────────────────────────────┴─────────┴────────────╯
╭────────────────────────────────────────────────────────────────────────────────────────────────────────────────╮
│ 9 special devices                                                                                              │
├────────────────────────────────┬──────┬────────┬───────┬───────────────────────────────┬──────────┬────────────┤
│ MOUNTED ON                     │ SIZE │   USED │ AVAIL │              USE%             │ TYPE     │ FILESYSTEM │
├────────────────────────────────┼──────┼────────┼───────┼───────────────────────────────┼──────────┼────────────┤
│ /dev                           │ 6.0G │     0B │  6.0G │                               │ devtmpfs │ none       │
│ /mnt/wsl                       │ 6.0G │ 398.8M │  5.6G │ [#...................]   6.5% │ tmpfs    │ tmpfs      │
│ /mnt/wsl/docker-desktop/shared │ 6.0G │  12.0K │  6.0G │ [....................]   0.0% │ tmpfs    │ none       │
│ -sockets/guest-services        │      │        │       │                               │          │            │
│ /mnt/wsl/docker-desktop/shared │ 6.0G │  12.0K │  6.0G │ [....................]   0.0% │ tmpfs    │ none       │
│ -sockets/host-services         │      │        │       │                               │          │            │
│ /run                           │ 6.0G │  12.0K │  6.0G │ [....................]   0.0% │ tmpfs    │ none       │
│ /run/lock                      │ 6.0G │     0B │  6.0G │                               │ tmpfs    │ none       │
│ /run/shm                       │ 6.0G │     0B │  6.0G │                               │ tmpfs    │ none       │
│ /run/user                      │ 6.0G │     0B │  6.0G │                               │ tmpfs    │ none       │
│ /sys/fs/cgroup                 │ 6.0G │     0B │  6.0G │                               │ tmpfs    │ tmpfs      │
╰────────────────────────────────┴──────┴────────┴───────┴───────────────────────────────┴──────────┴────────────╯
21/02/2 19:33:58 kvogel-elitebook:~/po/duf ±(master) ✗
❯
```
