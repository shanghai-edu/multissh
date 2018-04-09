 ## multissh

一个简单的并行 SSH 工具，可以批量的对主机通过 SSH 执行命令组合。

支持：
- 并发执行
- 单次执行多条命令
- ip 地址段自动匹配主机（192.168.0.1-192.168.0.100）
- ssh 用户名/密码认证
- ssh key 认证
- json 格式输出
- 输出到文本，文件名为 host.txt

#### 编译
```
go get ./...
go build
```

#### release
可以直接下载编译好的 release 版本

提供 win64 和 linux64 两个平台的可执行文件

https://github.com/shanghai-edu/multissh/releases/

#### 命令体系
```
# ./multissh -h
Usage of ./multissh:
  -c string
        cfg File Path
  -ciphers string
        ciphers
  -cmdfile string
        cmdfile path
  -cmds string
        cmds
  -f string
        write file locate
  -hostfile string
        hostfile path
  -hosts string
        host address list
  -ipfile string
        ipfile path
  -ips string
        ip address list
  -j    print output in json format
  -k string
        ssh private key
  -l    In linux mode,multi command combine with && ,such as date&&cd /opt&&ls
  -n int
        max execute number (default 20)
  -outTxt
        write result into txt
  -p string
        password
  -port int
        ssh port (default 22)
  -t int
        max timeout (default 30)
  -u string
        username
  -v    show version
```
**cmdfile 示例**
```
show clock
```
**hostfile 示例**
```
192.168.31.21
192.168.15.102
```
**ipfile 示例**
```
192.168.15.101-192.168.15.103
192.168.31.21-192.168.31.22
```

**ssh.json 示例**
```
{
    "SshHosts": [
        {
            "Host": "192.168.31.51",
            "Port": 22,
            "Username": "admin",
            "Password": "admin",
			"cmds":"show clock;show clock"
        },
        {
            "Host": "192.168.80.131",
            "Port": 22,
            "Username": "root",
            "Password": "",
            "key": "./server.key",
            "linuxMode": true,
            "CmdFile": "cmd2.txt.example"
        }
    ]
}
```

## 用法
#### cmd string & host string
```
# ./multissh -cmds "show clock" -hosts "192.168.31.21;192.168.15.102" -u admin -p password
2018/01/17 14:01:28 Multissh start
2018/01/17 14:01:31 Multissh finished. Process time 2.867808673s. Number of active ip is 2
host:  192.168.31.21
========= Result =========

******************************************************************************
* Copyright (c) 2004-2016 Hangzhou H3C Tech. Co., Ltd. All rights reserved.  *
* Without the owner's prior written consent,                                 *
* no decompiling or reverse-engineering shall be allowed.                    *
******************************************************************************

<sw-h3c>show clock
14:01:31 CN Wed 01/17/2018
Time Zone : CN add 08:00:00
<sw-h3c>exit

host:  192.168.15.102
========= Result =========

sw-cisco#show clock
05:50:24.935 UTC Wed Jan 17 2018
sw-cisco#exit

```

#### cmdfile & hostfile
```
# ./multissh -cmdfile cmd1.txt.example -hostfile host.txt.example -u admin -p password
2018/01/17 14:01:28 Multissh start
2018/01/17 14:01:31 Multissh finished. Process time 2.867808673s. Number of active ip is 2
host:  192.168.31.21
========= Result =========

******************************************************************************
* Copyright (c) 2004-2016 Hangzhou H3C Tech. Co., Ltd. All rights reserved.  *
* Without the owner's prior written consent,                                 *
* no decompiling or reverse-engineering shall be allowed.                    *
******************************************************************************

<sw-h3c>show clock
14:01:31 CN Wed 01/17/2018
Time Zone : CN add 08:00:00
<sw-h3c>exit

host:  192.168.15.102
========= Result =========

sw-cisco#show clock
05:50:24.935 UTC Wed Jan 17 2018
sw-cisco#exit

```

#### ipfile
```
# ./multissh -cmdfile cmd1.txt.example -ipfile ip.txt.example -u admin -p password
2018/01/17 14:25:26 Multissh start
2018/01/17 14:25:29 Multissh finished. Process time 2.847347642s. Number of active ip is 5
host:  192.168.15.101
========= Result =========

sw-cisco-1#show clock
06:17:49.422 UTC Wed Jan 17 2018
sw-cisco-1#exit

host:  192.168.15.102
========= Result =========
sw-cisco-2#show clock
06:14:22.445 UTC Wed Jan 17 2018
sw-cisco-2#exit

host:  192.168.15.103
========= Result =========
sw-cisco-3#show clock
06:19:14.487 UTC Wed Jan 17 2018
sw-cisco-3#exit

host:  192.168.31.21
========= Result =========

******************************************************************************
* Copyright (c) 2004-2016 Hangzhou H3C Tech. Co., Ltd. All rights reserved.  *
* Without the owner's prior written consent,                                 *
* no decompiling or reverse-engineering shall be allowed.                    *
******************************************************************************

<sw-h3c>show clock
14:25:29 CN Wed 01/17/2018
Time Zone : CN add 08:00:00
<sw-h3c>exit

host:  192.168.31.22
========= Result =========

sw-cisco-4#show clock
14:25:27.639 beijing Wed Jan 17 2018
sw-cisco-4#exit
```
#### ssh key-based Auth and linuxMode
```
# ./multissh -hosts "192.168.80.131" -cmds "date;cd /opt;ls" -u root -k "server.key"
2018/01/17 14:33:55 Multissh start
2018/01/17 14:33:56 Multissh finished. Process time 960.367764ms. Number of active ip is 1
host:  192.168.80.131
========= Result =========
Welcome to Ubuntu 16.04.3 LTS (GNU/Linux 4.4.0-98-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

  System information as of Wed Jan 17 14:33:55 CST 2018

  System load:  0.0                Processes:              335
  Usage of /:   10.0% of 90.18GB   Users logged in:        0
  Memory usage: 2%                 IP address for eth0:    192.168.80.131
  Swap usage:   0%                 IP address for docker0: 172.17.0.1

  Graph this data and manage this system at:
    https://landscape.canonical.com/

0 个可升级软件包。
0 个安全更新。

New release '17.10' available.
Run 'do-release-upgrade' to upgrade to it.

You have new mail.
Last login: Wed Jan 17 14:29:39 2018 from 202.120.80.201
root@ubuntu-docker-node3:~# 201817:33:56 CST
root@ubuntu-docker-node3:~# root@ubuntu-docker-node3:/opt# cisco
composer.json
composer.phar
example-oauth2-server
getting-started-with-mmdb
gitlab
gitlab-ce_8.0.4-ce.1_amd64.deb
oauth2-demo-php
oauth2-server-php
python_test
rsyslog-maxminddb
root@ubuntu-docker-node3:/opt# 注销

# ./multissh -hosts "192.168.80.131" -cmds "date;cd /opt;ls" -u root -k "server.key" -l
2018/01/17 14:34:02 Multissh start
2018/01/17 14:34:02 Multissh finished. Process time 842.465643ms. Number of active ip is 1
host:  192.168.80.131
========= Result =========
201817:34:02 CST
cisco
composer.json
composer.phar
example-oauth2-server
getting-started-with-mmdb
gitlab
gitlab-ce_8.0.4-ce.1_amd64.deb
oauth2-demo-php
oauth2-server-php
python_test
rsyslog-maxminddb

```

#### ssh.json
```
./multissh -c ssh.json.example
2018/01/17 14:29:38 Multissh start
2018/01/17 14:29:41 Multissh finished. Process time 2.922928532s. Number of active ip is 2
host:  192.168.31.51
========= Result =========

******************************************************************************
* Copyright (c) 2004-2016 Hangzhou H3C Tech. Co., Ltd. All rights reserved.  *
* Without the owner's prior written consent,                                 *
* no decompiling or reverse-engineering shall be allowed.                    *
******************************************************************************

<sw-h3c>show clock
14:29:41 CN Wed 01/17/2018
Time Zone : CN add 08:00:00
<WenKe-5F-Stack-2>show clock
14:29:41 CN Wed 01/17/2018
Time Zone : CN add 08:00:00
<WenKe-5F-Stack-2>exit

host: 192.168.80.131
========= Result =========
cisco
composer.json
composer.phar
example-oauth2-server
getting-started-with-mmdb
gitlab
gitlab-ce_8.0.4-ce.1_amd64.deb
oauth2-demo-php
oauth2-server-php
python_test
rsyslog-maxminddb
```

#### LICENSE
Apache License 2.0