# goInception

[![travis-ci](https://img.shields.io/travis/hanchuanchuan/goInception.svg)](https://travis-ci.org/hanchuanchuan/goInception)
[![CircleCI Status](https://circleci.com/gh/hanchuanchuan/goInception.svg?style=shield)](https://circleci.com/gh/hanchuanchuan/goInception)
[![GitHub release](https://img.shields.io/github/release-pre/hanchuanchuan/goInception.svg?style=brightgreen)](https://github.com/hanchuanchuan/goInception/releases)
[![codecov](https://codecov.io/gh/hanchuanchuan/goInception/branch/master/graph/badge.svg)](https://codecov.io/gh/hanchuanchuan/goInception)
[![](https://img.shields.io/badge/go-1.12-brightgreen.svg)](https://golang.org/dl/)
[![TiDB](https://img.shields.io/badge/TiDB-v2.1.1-brightgreen.svg)](https://github.com/pingcap/tidb)
![](https://img.shields.io/github/downloads/hanchuanchuan/goInception/total.svg)
![](https://img.shields.io/github/license/hanchuanchuan/goInception.svg)


**[[English]](README.md)**
**[[Chinese]](README.zh.md)**


goInception is a MySQL maintenance tool, which can be used to review, implement, backup, and generate SQL statements for rollback. It parses SQL syntax and returns the result of the review based on custom rules.

**Documentation:**
**[[Document]](https://hanchuanchuan.github.io/goInception/)**
**[[中文文档]](https://hanchuanchuan.github.io/goInception/zh/)**

**[[Changelog]](https://hanchuanchuan.github.io/goInception/changelog.html)**


----

### Quick start


#### Binary

[goInception Download](https://github.com/hanchuanchuan/goInception/releases)


#### Docker Image
```
docker pull hanchuanchuan/goinception
```


#### Source code compilation

***go version 1.14+ (go mod)***

```bash
git clone https://github.com/hanchuanchuan/goInception.git
cd goInception
go build -o goInception tidb-server/main.go

./goInception -config=config/config.toml


---注意---
端口4000 密码为空
server 端登录可在线修改参数或是在config.toml 中修改重启
mysql -hlocalhost -uroot -P4000
注： inception get variables;  查看参数
    inception set enable_set_collation='true';  设置参数

| osc_bin_dir                            | /usr/bin   此参数与pt工具目录一致


```


----

***pt 工具安装***

```bash
官网https://www.percona.com/downloads/percona-toolkit

wget https://www.percona.com/downloads/percona-toolkit/3.1.0/binary/redhat/7/x86_64/percona-toolkit-3.5.0-2.el7.x86_64.rpm
yum install perl-DBI perl-DBD-MySQL perl-Digest-MD5 perl-IO-Socket-SSL perl-TermReadKey
yum install -y percona-toolkit-3.5.0-2.el7.x86_64.rpm
yum install percona-xtrabackup-80

报错 Transaction check error: file /etc/my.cnf from install of Percona-Server-shared-56-5.6.51-rel91.0.1.el7.x86_64 conflicts with file from package mysql-community-server-8.0.32-1.el7.x86_64
处理 yum install mysql-community-libs-compat -y  [说明：重新再安装一下]


```


#### Associated SQL audit platform

* [Archery](https://github.com/hhyo/Archery) `Query support (MySQL/MsSQL/Redis/PostgreSQL), MySQL optimization (SQLAdvisor|SOAR|SQLTuning), slow log management, table structure comparison, session management, Alibaba Cloud RDS management, etc.`


#### Acknowledgments
    GoInception reconstructs from the Inception which is a well-known MySQL auditing tool and uses TiDB SQL parser.

- [Inception](https://github.com/hanchuanchuan/inception)
- [TiDB](https://github.com/pingcap/tidb)

#### Sponsorship and support
- [Sponsorship and support](https://hanchuanchuan.github.io/goInception/support.html)

#### Contact

QQ group talk: **499262190**
e-mail: `chuanchuanhan@gmail.com`

### Contributing

Welcome and thank you very much for your contribution. For the process of submitting PR, please refer to [CONTRIBUTING.md](CONTRIBUTING.md)。


## Contributors

### Code Contributors

This project exists thanks to all the people who contribute. [[Contribute](CONTRIBUTING.md)].
<a href="https://github.com/hanchuanchuan/goInception/graphs/contributors"><img src="https://opencollective.com/goInception/contributors.svg?width=890&button=false" /></a>

### Financial Contributors

Become a financial contributor and help us sustain our community. [[Contribute](https://opencollective.com/goInception/contribute)]

#### Individuals

<a href="https://opencollective.com/goInception"><img src="https://opencollective.com/goInception/individuals.svg?width=890"></a>

#### Organizations

Support this project with your organization. Your logo will show up here with a link to your website. [[Contribute](https://opencollective.com/goInception/contribute)]

<a href="https://opencollective.com/goInception/organization/0/website"><img src="https://opencollective.com/goInception/organization/0/avatar.svg"></a>
<a href="https://opencollective.com/goInception/organization/1/website"><img src="https://opencollective.com/goInception/organization/1/avatar.svg"></a>
<a href="https://opencollective.com/goInception/organization/2/website"><img src="https://opencollective.com/goInception/organization/2/avatar.svg"></a>
<a href="https://opencollective.com/goInception/organization/3/website"><img src="https://opencollective.com/goInception/organization/3/avatar.svg"></a>
<a href="https://opencollective.com/goInception/organization/4/website"><img src="https://opencollective.com/goInception/organization/4/avatar.svg"></a>
<a href="https://opencollective.com/goInception/organization/5/website"><img src="https://opencollective.com/goInception/organization/5/avatar.svg"></a>
<a href="https://opencollective.com/goInception/organization/6/website"><img src="https://opencollective.com/goInception/organization/6/avatar.svg"></a>
<a href="https://opencollective.com/goInception/organization/7/website"><img src="https://opencollective.com/goInception/organization/7/avatar.svg"></a>
<a href="https://opencollective.com/goInception/organization/8/website"><img src="https://opencollective.com/goInception/organization/8/avatar.svg"></a>
<a href="https://opencollective.com/goInception/organization/9/website"><img src="https://opencollective.com/goInception/organization/9/avatar.svg"></a>
