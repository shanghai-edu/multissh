 ## multissh

一个简单的并行 SSH 工具，可以批量的对主机通过 SSH 执行命令组合。

#### 编译
```
go get ./...
go build
```


#### 命令体系
```
./multissh -h
  -cmd string
        cmds // 需要执行的命令组合，多条命令以 ； 分割
  -cmdfile string
        cmdfile path //需要执行的命令组合文件，文件内命令按行分割
  -hostfile string
        hostfile path // 需要执行的主机列表文件，主机列表在文件内按行分割
  -hosts string
        host address list //需要执行的主机列表，多个主机以 ; 分割
  -ipfile string
        hostfile path //需要执行的主机（IP）列表文件，IP可以以地址段的方式逐行写在文本内
  -p string
        password // 主机的 SSH 密码
  -port int
        ssh port (default 22) //主机的 SSH 端口，默认 22
  -u string
        username //主机的 SSH 用户名
```
**cmdfile 示例**
```
show clock
exit
```
**hostfile 示例**
```
10.10.15.101
10.10.15.102
```
**ipfile 示例**
```
10.10.15.101-10.10.15.102
```

## 用法
#### cmd string & host string
```
./multissh -cmd "show clock;exit" -hosts "10.10.15.101;10.10.15.102" -u admin -p admin

10.10.15.101  ssh start
sw-1#show clock
05:26:40.649 UTC Tue Jun 6 2017
sw-1#exit

10.10.15.101  ssh end

10.10.15.102  ssh start
sw-2#show clock
05:24:38.708 UTC Tue Jun 6 2017
sw-2#exit

10.10.15.102  ssh end
```

#### cmdfile & hostfile
```
./multissh -cmdfile cmd.txt -hostfile host.txt -u admin -p admin

10.10.15.101  ssh start
sw-1#show clock
05:29:43.269 UTC Tue Jun 6 2017
sw-1#exit

10.10.15.101  ssh end

10.10.15.102  ssh start
sw-2#show clock
05:27:41.332 UTC Tue Jun 6 2017
sw-2#exit

10.10.15.102  ssh end
```

#### ipfile
```
./multissh -cmdfile cmd.txt -ipfile ip.txt -u admin -p admin

10.10.15.101  ssh start
sw-1#show clock
05:29:43.269 UTC Tue Jun 6 2017
sw-1#exit

10.10.15.101  ssh end

10.10.15.102  ssh start
sw-2#show clock
05:27:41.332 UTC Tue Jun 6 2017
sw-2#exit

10.10.15.102  ssh end
```

