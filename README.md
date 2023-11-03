## 综述

使用Golang编写的飞常准ADSB上传程序，适用于各个支持Golang编译的平台。
如果此项目您你有帮助，请给我一个Star，谢谢。
如果您有任何建议，欢迎提出一个issue。

## 关于2.0版本

2.0版本正在开发。请注意：与1.0在配置文件上不兼容

### 二进制发布大小膨胀提示

2.0由于使用了一些库的缘故，二进制大小将和1.0版本有较大差异

release编译模式下，在 Ubuntu 22.04 5.14.0-1059-oem 上，目前amd64编译产物大小为6.1MB

建议具有微小设备运行的用户尝试使用UPX进行压缩，上述平台经过UPX压缩后可降至2.1MB

## Docker 使用说明

docker仓库地址：https://hub.docker.com/r/dextercai/feeyo-adsb-golang

如您所见，我们为x86/64和ARM平台都提供了对应的docker镜像，并托管在了DockerHub仓库。
如果您需要其他架构版本的镜像，欢迎提出一个PR。

### 使用Docker与文件配置(conf.ini)
```bash
docker run --net host \
  -v /YOUR-PATH-OF/conf.ini:/app/conf.ini:r -d \
  dextercai/feeyo-adsb-golang:latest --conf=/app/conf.ini
```

### 使用Docker与命令行配置
```bash
docker run --net host \
  dextercai/feeyo-adsb-golang:latest /app/feeyo-adsb-golang \
  --log.path=/tmp --log.rotation_count=1 --log.rotation_size=2 --log.rotation_time=0\
  --dump1090.host --dump1090.port=30003 --feeyo.uuid=YOUR-UUID

# 日志会被导向到临时目录中，不占用磁盘。
```

## Binary 使用说明

由于本项目不包括Dump1090，也不限制SBS服务是否运行在本机，因此你可能需要首先安装Dump1090，具体细节可自行搜索，当然你也可以在本项目提一个Issue，我将很乐意为你解答。

如果你不具备编译条件，可以直接前往[本项目发布页](https://github.com/dextercai/feeyo-adsb-golang/releases)下载使用。

具有两种配置方式

UUID在线生成可访问：https://feeyo-uuid.dextercai.com

### 一般文件模式（默认）

你需要在程序**同目录**创建conf.ini文件，内容如下。

```
[dump1090]
host=127.0.0.1
port=30003

[log]
level=debug
file=feeyo-adsb-golang.log
path=./logs/
rotation_time=86400  # 单位秒
max_age=604800 # 单位秒
# max_age 与 rotation_count 不可同时配置
rotation_size=32 #单位 MB
rotation_count=0

[feeyo]
uuid=你的UUID（16位）
url=http://adsb.feeyo.com/adsb/ReceiveCompressADSB.php

```

以上展现的是dump1090运行在本机的情况，你也可以按照实际情况进行填写。

### 命令行模式（进阶）

若对终端操作较为熟悉，可使用该方式。

```
Usage of /adsb:
      --conf string               配置文件位置 (default "./config.ini")
      --dump1090.host string      dump1090服务地址 (default "127.0.0.1")
      --dump1090.port int         dump1090服务端口 (default 30003)
      --feeyo.url string          飞常准上传接口 (建议保留默认) (default "http://adsb.feeyo.com/adsb/ReceiveCompressADSB.php")
      --feeyo.uuid string         设备UUID
      --log.file string           日志存储文件 (default "feeyo-adsb-golang.log")
      --log.level string          日志等级 (default "info")
      --log.max_age int           日志最大保留时间 单位秒 (default 604800)
      --log.path string           日志存储路径 (default "./logs/")
      --log.rotation_count uint   日志轮转个数 (max_age 与 rotation_count 不可同时配置)
      --log.rotation_size int     日志轮转大小 单位MB (为嵌入式设备设计) (default 10)
      --log.rotation_time int     日志轮转时间 单位秒 (default 86400)
```

## 其他

如果使用树莓派加RTL2832为主控的电视棒，建议您前往（[飞常准ADSB](https://flightadsb.variflight.com/)）

使用官方ADSB脚本，或者前往（[FEEYO-Adsb](https://github.com/dextercai/FEEYO-Adsb)），项目内有一份相同的官方脚本。

## 以下内容不再适用于2.0版本

~~如果使用其他Linux发行版，可参考下面列出的相关资料，手动移植，或者使用本项目。~~

~~飞常准自建 ADS-B Windows 上传方案~~
~~https://blog.dextercai.com/archives/78.html~~

~~在 Arch Linux 下安装飞常准上传套件~~
~~https://blog.dextercai.com/archives/45.html~~

