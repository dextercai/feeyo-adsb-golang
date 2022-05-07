## 综述

使用Golang编写的飞常准ADSB上传程序，适用于各个支持Golang编译的平台。
如果此项目您你有帮助，请给我一个Star，谢谢。
如果您有任何建议，欢迎提出一个issue。

## 使用说明

由于本项目不包括Dump1090，也不限制SBS服务是否运行在本机，因此你可能需要首先安装Dump1090，具体细节可自行搜索，当然你也可以在本项目提一个Issue，我将很乐意为你解答。

如果你不具备编译条件，可以直接前往[本项目发布页](https://github.com/dextercai/feeyo-adsb-golang/releases)下载使用。

具有两种配置方式

### 一般文件模式（默认）

你需要在程序**同目录**创建conf.ini文件，内容如下。

```
[config]
UUID=你的UUID（16位）
ip=127.0.0.1
port=30003
url=http://adsb.feeyo.com/adsb/ReceiveCompressADSB.php
```

以上展现的是dump1090运行在本机的情况，你也可以按照实际情况进行填写。

### 命令行模式（进阶）

若对终端操作较为熟悉，可使用该方式。

```
Usage of ./adsb:
  -conf string
        conf文件位置 (default "./conf.ini")
  -feeyo-url string
        飞常准接口地址 (default "https://adsb.feeyo.com/adsb/ReceiveCompressADSB.php")
  -ip string
        dump1090服务IP (default "127.0.0.1")
  -port string
        dump1090服务端口 (default "30003")
  -use-file
        是否使用conf文件作为配置来源 (default true)
  -uuid string
        UUID 16位
```

## TODO
- 统计、集成地图
- 集成部分dump1090功能
- webhook


## 其他

如果使用树莓派加RTL2832为主控的电视棒，建议您前往（[飞常准ADSB](https://flightadsb.variflight.com/)）

使用官方ADSB脚本，或者前往（[FEEYO-Adsb](https://github.com/dextercai/FEEYO-Adsb)），项目内有一份相同的官方脚本。

如果使用其他Linux发行版，可参考下面列出的相关资料，手动移植，或者使用本项目。

飞常准自建 ADS-B Windows 上传方案
https://blog.dextercai.com/2020-03-20-78.html

在 Arch Linux 下安装飞常准上传套件
https://blog.dextercai.com/2018-06-24-45.html

