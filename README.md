# docker-haproxy-nginx

## Usage

- start: `make up`, clean: `make clean`
- test: `make t1`

## logs

`make up` logs:

```log
docker-haproxy-nginx3-1   | 172.27.0.5 - - [13/Jun/2022:09:59:30 +0000] "GET / HTTP/1.1" 200 14 "-" "curl/7.79.1" "172.27.0.1"
docker-haproxy-haproxy-1  | 2022-06-13T09:59:30+00:00 localhost haproxy[13]: {haproxy_clientIP:172.27.0.1,haproxy_clientPort:63892,haproxy_dateTime:13/Jun/2022:09:59:30.332,haproxy_frontendNameTransport:proxy01,haproxy_backend:proxy01,haproxy_serverName:nginx3,haproxy_Tw:0,haproxy_Tc:0,haproxy_Tt:4,haproxy_bytesRead:249,haproxy_terminationState:--,haproxy_actconn:1,haproxy_FrontendCurrentConn:1,haproxy_backendCurrentConn:0,haproxy_serverConcurrentConn:0,haproxy_retries:0,haproxy_srvQueue:0,haproxy_backendQueue:0,haproxy_backendSourceIP:172.27.0.5,haproxy_backendSourcePort:55514}
```

## HAProxy versions

[versions](http://www.haproxy.org/)

Branch  | Release date | End of life                   | Latest version | Changelog
--------|--------------|-------------------------------|----------------|-----------
2.7-dev | ~2022-Q4     | 2024-Q1 (dev » stable)        | 2.7-dev0       | 2022/05/31
2.6     | 2022-05-31   | 2027-Q2 (LTS)                 | 2.6.0          | 2022/05/31
2.5     | 2021-11-23   | 2023-Q1 (stable)              | 2.5.7          | 2022/05/13
2.4     | 2021-05-14   | 2026-Q2 (LTS)                 | 2.4.17         | 2022/05/13
2.2     | 2020-07-07   | 2025-Q2 (LTS)                 | 2.2.24         | 2022/05/13
2.0     | 2019-06-16   | 2024-Q2 (critical fixes only) | 2.0.29         | 2022/05/13
1.8     | 2017-11-26   | 2022-Q4 (critical fixes only) | 1.8.30         | 2021/04/12

- version 2.6 : May, 31st, 2022  the latest long-term supported release.
- version 2.5 : runtime server addition/removal, runtime CA/CRL updates, native HTTP client, simplified HTTPS logging, default TCP/HTTP rulesets, JWT validation, and more
- version 2.4 : syslog and DNS over TCP, multi-threaded Lua, full sharing of idle conns, lower latency, server-side dynamic SSL update, Opentracing, WebSocket over H2, atomic maps, Vary support, new debugging tools, even - more user-friendly CLI and configuration, lots of cleanups
- version 2.3 : syslog forwarding, better idle conn management, improved balancing with large queues, simplified SSL managment, more stats metrics, stricter config checking by default, general performance improvements
- version 2.2 : runtime certificate additions, improved idle connection management, logging over TCP, HTTP "return" directive, errorfile templates, TLSv1.2 by default, extensible health-checks
- version 2.1 : improved I/Os and multi-threading, FastCGI, runtime certificate updates, HTX-only, improved debugging, removal of obsolete keywords
- version 2.0 : gRPC, layer 7 retries, process manager, SSL peers, log load balancing/sampling, end-to-end TCP fast-open, automatic settings (maxconn, threads, HTTP reuse, pools), ...
- version 1.9 : improved multi-threading, end-to-end HTTP/2, connection pools, queue priority control, stdout logging, ...
- version 1.8 : multi-threading, HTTP/2, cache, on-the fly server addition/removal, seamless reloads, DNS SRV, hardware SSL engines, ...
- version 1.7 : added server hot reconfiguration, content processing agents, multi-type certs, ...
- version 1.6 : added DNS resolution support, HTTP connection multiplexing, full stick-table replication, stateless compression, ...
- version 1.5 : added SSL, IPv6, keep-alive, DDoS protection, ...

## Resouces

1. [Docker HAproxy 配置 & rsyslog 日志处理](https://blog.csdn.net/zhengxia19/article/details/115210843)
2. [How To Configure HAProxy Logging with Rsyslog on CentOS 8](https://www.digitalocean.com/community/tutorials/how-to-configure-haproxy-logging-with-rsyslog-on-centos-8-quickstart)

## syslog协议的Facility, Severity数字代号和PRI计算

[原文](https://www.ichenfu.com/2017/08/31/syslog-facility-and-severity/)

一条syslog信息包含三部分，PRI, HEADER和MSG，其中PRI是<>扩起来的一个数字，这个数字就代表了不同Facility和Severity的消息。
其中Facility, Severity的数字代号列表如下：

Facility 消息来源:

Code | Keyword  | Description                                | 解释
-----|----------|--------------------------------------------|------------------------------------
0    | kern     | kernel messages                            | 内核的信息
1    | user     | user-level messages                        | 本地用户的应用程序的信息
2    | mail     | mail system                                | 处理邮件的守护进程发出的信息
3    | daemon   | system daemons                             | 某些系统的守护程序的 syslog ，如由in.ftpd 产生的log
4    | auth     | security/authorization messages            | 认证系统，如login 或su ，即询问用户名和口令
5    | syslog   | messages generated internally by syslogd   |
6    | lpr      | line printer subsystem (archaic subsystem) | 打印机的信息
7    | news     | network news subsystem (archaic subsystem) | 新闻组的守护进程的信息
8    | uucp     | UUCP subsystem (archaic subsystem)         | uucp 子系统的信息
9    | clock    | daemon                                     |
10   | authpriv | security/authorization messages            |
11   | ftp      | FTP daemon                                 |
12   | -        | NTP subsystem                              |
13   | -        | log audit                                  |
14   | -        | log alert                                  |
15   | cron     | scheduling daemon                          | 系统执行定时任务时发出的信息
16   | local0   | local use 0 (local0)                       | 系统预留
17   | local1   | local use 1 (local1)                       | 系统预留
18   | local2   | local use 2 (local2)                       | 系统预留
19   | local3   | local use 3 (local3)                       | 系统预留
20   | local4   | local use 4 (local4)                       | 系统预留
21   | local5   | local use 5 (local5)                       | 系统预留
22   | local6   | local use 6 (local6)                       | 系统预留
23   | local7   | local use 7 (local7)                       | 系统预留

Severity 优先级:

Code | Keyword | Description                                                     | 解释
-----|---------|-----------------------------------------------------------------|----------
0    | emerg   | System is unusable                                              | 最高的紧急程度状态
1    | alert   | Should be corrected immediately                                 | 紧急状态
2    | crit    | Critical conditions                                             | 重要信息
3    | err     | Error conditions                                                | 警告
4    | warning | May indicate that an error will occur if action is not taken.   | 临界状态
5    | notice  | Events that are unusual, but not error conditions.              | 出现不寻常的事情
6    | info    | Normal operational messages that require no action.             | 一般性消息
7    | debug   | Information useful to developers for debugging the application. | 调试级信息

针对PRI的计算公式：`PRI = FacilityCode*8 + SeverityCode`

举个例子： local3.info的日志，它的PRI就是19*8+6=158，所以这条消息在传输中的格式为<158> {HEADER} {MEG}
再一个例子，如果看到一条PRI为14的消息，那么它实际的级别就是user.info (1*8+6=14)

参考：

- [rfc3164 The BSD syslog Protocol](https://tools.ietf.org/html/rfc3164)
- [systemd Journal](https://wiki.archlinux.org/index.php/systemd#Journal)

## /etc/rsyslog.conf配置文件的格式

该配置文件的基本格式如下所示：

```sh
authpriv.* /var/log/secure
#服务名称[连接符号]日志等级 日志记录位置
#认证相关服务.所有日志等级 记录在/var/log/secure日志中
```

日志服务连接日志等级的格式`日志服务[连接符号]日志等级 日志记录位置`

在这里，连接符号可以被识别为以下三种:

- `.` 代表只要比后面的等级高的（包含该等级）日志都记录。比如，“cron.info”代表cron服务产生的日志，只要日志等级大于等于info级别，就记录。
- `.=` 代表只记录所需等级的日志，其他等级的日志都不记录。比如，“*.=emerg”代表人和日志服务产生的日志，只要等级是emerg等级，就记录。这种用法极少见，了解就好。
- `.!` 代表不等于，也就是除该等级的日志外，其他等级的日志都记录

## 其它

rsyslog 在Linux上自带，兼容syslog语法，在syslog基础上增加了更多协议的支持，配合额外module插件可以完成很多场景的使用。

- Check rsyslog version with: `rsyslogd -version`
- Try to check rsyslog conf with: `rsyslogd -f /etc/rsyslog.conf -N 1`
- 修改配置文件后，重启服务： `sudo /etc/init.d/rsyslog restart`
- check if rsyslog started or not: `logger "hello"`
- 完全禁用SELinux或执行以下命令并重新启动rsyslog：`semanage port -a -t syslogd_port_t -p tcp 10544`


## rsyslog日志配置

```sh
# 配置haproxy日志

# 配置文件生效参数 UDP 接收日志数据 /etc/rsyslog.conf
# $ModLoad imudp
# $UDPServerRun 514
# 替换Rsyslog配置文件注释掉的参数
sed -i 's/\#$ModLoad imudp/$ModLoad imudp/' /etc/rsyslog.conf
sed -i 's/\#$UDPServerRun 514/$UDPServerRun 514/' /etc/rsyslog.conf

# 在local3收info的日志写入文件
echo "local3.*   /var/log/haproxy.log" >> /etc/rsyslog.conf
# 在local4收warning的日志写入文件
echo "local4.*   /var/log/haproxy_warning.log" >> /etc/rsyslog.conf


systemctl restart rsyslog
```

## 日志切割神器 logrotate

logrotate 是一个 linux 系统日志的管理工具。可以对单个日志文件或者某个目录下的文件按时间 / 大小进行切割，压缩操作；指定日志保存数量；还可以在切割之后运行自定义命令。

logrotate 是基于 crontab 运行的，所以这个时间点是由 crontab 控制的，具体可以查询 crontab 的配置文件 /etc/anacrontab。 系统会按照计划的频率运行 logrotate，通常是每天。在大多数的 Linux 发行版本上，计划每天运行的脚本位于 /etc/cron.daily/logrotate。

主流 Linux 发行版上都默认安装有 logrotate 包，如果你的 linux 系统中找不到 logrotate, 可以使用 apt-get 或 yum 命令来安装。

- 执行文件： /usr/sbin/logrotate
- 主配置文件: /etc/logrotate.conf
- 自定义配置文件: /etc/logrotate.d/*.conf
- 修改配置文件后，并不需要重启服务。(由于 logrotate 实际上只是一个可执行文件，不是以 daemon 运行)

常见参数

- daily ：指定转储周期为每天
- weekly ：指定转储周期为每周
- monthly ：指定转储周期为每月
- rotate count ：指定日志文件删除之前转储的次数，0 指没有备份，5 指保留 5 个备份
- tabooext [+] list：让 logrotate 不转储指定扩展名的文件，缺省的扩展名是：.rpm-orig, .rpmsave, v, 和～
- missingok：在日志轮循期间，任何错误将被忽略，例如 “文件无法找到” 之类的错误。
- size size：当日志文件到达指定的大小时才转储，bytes (缺省) 及 KB (sizek) 或 MB (sizem)
- compress： 通过 gzip 压缩转储以后的日志
- nocompress： 不压缩
- copytruncate：用于还在打开中的日志文件，把当前日志备份并截断
- nocopytruncate： 备份日志文件但是不截断
- create mode owner group ： 转储文件，使用指定的文件模式创建新的日志文件
- nocreate： 不建立新的日志文件
- delaycompress： 和 compress 一起使用时，转储的日志文件到下一次转储时才压缩
- nodelaycompress： 覆盖 delaycompress 选项，转储同时压缩。
- errors address ： 专储时的错误信息发送到指定的 Email 地址
- ifempty ：即使是空文件也转储，这个是 logrotate 的缺省选项。
- notifempty ：如果是空文件的话，不转储
- mail address ： 把转储的日志文件发送到指定的 E-mail 地址
- nomail ： 转储时不发送日志文件
- olddir directory：储后的日志文件放入指定的目录，必须和当前日志文件在同一个文件系统
- noolddir： 转储后的日志文件和当前日志文件放在同一个目录下
- prerotate/endscript： 在转储以前需要执行的命令可以放入这个对，这两个关键字必须单独成行

```sh
# 这个配置将会每天转储 /var/log/haproxy.log 文件
# /etc/logrotate.conf 配置文件有 include /etc/logrotate.d
# 在/etc/logrotate.d下面写入一个配置文件
cat << \EOF > /etc/logrotate.d/haproxy
/var/log/haproxy.log { # 处理的日志文件
    missingok # 在日志轮循期间，任何错误将被忽略，例如 “文件无法找到” 之类的错误
    daily # 每日
    compress # 压缩
    minsize 10M # 最小size
    rotate 7 # 保留份数
    notifempty # 如果是空文件的话，不转储
    dateext # 转储文件加日期，效果：haproxy.log-20210320.gz
}
EOF

# 同上，处理warning日志
cat << \EOF > /etc/logrotate.d/haproxy_warning
/var/log/haproxy_warning.log {
    missingok
    daily
    compress
    minsize 10M
    rotate 7
    notifempty
    dateext
}
EOF
```

## 操作系统优化

```sh
# 操作系统参数优化
cat << \EOF >> /etc/sysctl.conf
net.ipv4.tcp_tw_reuse = 1
net.ipv4.tcp_tw_recycle = 0
net.ipv4.ip_local_port_range="1024 65024"
net.ipv4.tcp_max_syn_backlog=100000
net.core.netdev_max_backlog=100000
net.ipv4.tcp_tw_reuse=1
net.ipv4.tcp_fin_timeout = 10
fs.file-max=500000
EOF

echo '*  -  nofile  500000' >> /etc/security/limits.conf
```

## install rsyslog on Alpine

[install rsyslog on Alpine](https://www.rsyslog.com/alpine-repository/)

The Adiscon Alpine Repository  supports recent rsyslog versions for Alpine Linux including necessary third party packages.

To install rsyslog on Alpine, simply execute the following commands as root from the commandline:

```sh
cd /etc/apk/keys
wget http://alpine.adiscon.com/rsyslog@lists.adiscon.com-5a55e598.rsa.pub
echo 'http://alpine.adiscon.com/3.7/stable' >> /etc/apk/repositories
apk update
apk add rsyslog
```

[HAproxy增加日志记录功能](https://blog.51cto.com/eric1/1854574)

```sh
vi haproxy.conf(在default处添加如下信息)
defaults
log global
option httplog
log 127.0.0.1 local3
```

```sh
vi /etc/rsyslog.conf(添加如下内容)
local3.* /data/logs/haproxy.log
```

```sh
vi /etc/sysconfig/rsyslog

把SYSLOGD_OPTIONS="-m 0"
改成 SYSLOGD_OPTIONS="-r -m 0 -c 2"
```

相关解释说明:

- `-r`: 打开接受外来日志消息的功能,其监控514 UDP端口;
- `-x`: 关闭自动解析对方日志服务器的FQDN信息,这能避免DNS不完整所带来的麻烦;
- `-m`: 修改syslog的内部mark消息写入间隔时间(0为关闭),例如240为每隔240分钟写入一次"--MARK--"信息;
- `-h`: 默认情况下,syslog不会发送从远端接受过来的消息到其他主机,而使用该选项,则把该开关打开,所有

接受到的信息都可根据syslog.conf中定义的@主机转发过去

配置完毕后重启 rsyslogd 和 haproxy 即可.

## haproxy-dconv logging

[doc](https://cbonte.github.io/haproxy-dconv/1.8/management.html#8)

```log
For logging, HAProxy always relies on a syslog server since it does not perform
any file-system access. The standard way of using it is to send logs over UDP
to the log server (by default on port 514). Very commonly this is configured to
127.0.0.1 where the local syslog daemon is running, but it's also used over the
network to log to a central server. The central server provides additional
benefits especially in active-active scenarios where it is desirable to keep
the logs merged in arrival order. HAProxy may also make use of a UNIX socket to
send its logs to the local syslog daemon, but it is not recommended at all,
because if the syslog server is restarted while haproxy runs, the socket will
be replaced and new logs will be lost. Since HAProxy will be isolated inside a
chroot jail, it will not have the ability to reconnect to the new socket. It
has also been observed in field that the log buffers in use on UNIX sockets are
very small and lead to lost messages even at very light loads. But this can be
fine for testing however.

It is recommended to add the following directive to the "global" section to
make HAProxy log to the local daemon using facility "local0" :

      log 127.0.0.1:514 local0

and then to add the following one to each "defaults" section or to each frontend
and backend section :

      log global

This way, all logs will be centralized through the global definition of where
the log server is.

Some syslog daemons do not listen to UDP traffic by default, so depending on
the daemon being used, the syntax to enable this will vary :

  - on sysklogd, you need to pass argument "-r" on the daemon's command line
    so that it listens to a UDP socket for "remote" logs ; note that there is
    no way to limit it to address 127.0.0.1 so it will also receive logs from
    remote systems ;

  - on rsyslogd, the following lines must be added to the configuration file :

      $ModLoad imudp
      $UDPServerAddress *
      $UDPServerRun 514

  - on syslog-ng, a new source can be created the following way, it then needs
    to be added as a valid source in one of the "log" directives :

      source s_udp {
        udp(ip(127.0.0.1) port(514));
      };
```
