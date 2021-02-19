## 综述

使用Golang编写的飞常准ADSB上传程序，适用于各个支持Golang编译的平台。
如果此项目您你有帮助，请给我一个Star，谢谢。
如果您有任何建议，欢迎提出一个issue。

## 使用说明

本项目已更新，修改为Golang 1.11起支持的Go modules。
主分支去除了UUID生成器的代码，但依旧保留二进制版本，请自定义UUID时确保不与其他人冲突。

由于本项目不包括Dump1090，也不限制SBS服务是否运行在本机，因此你可能需要首先安装Dump1090，具体细节可自行搜索，当然你也可以在本项目提一个Issue，我将很乐意为你解答。

你需要编辑conf.ini文件

```
[config]
UUID=你的UUID（16位）
ip=127.0.0.1
port=30003
url=http://adsb.feeyo.com/adsb/ReceiveCompressADSB.php
```

以上展现的是dump1090运行在本机的情况，你也可以按照实际情况进行填写。

如果你不具备编译条件，可以直接前往[本项目发布页](https://github.com/dextercai/feeyo-adsb-golang/releases)下载使用。

## 其他

如果使用树莓派加RTL2832为主控的电视棒，建议您前往（[飞常准ADSB](https://flightadsb.variflight.com/)）

使用官方ADSB脚本，或者前往（[FEEYO-Adsb](https://github.com/dextercai/FEEYO-Adsb)），项目内有一份相同的官方脚本。

如果使用其他Linux发行版，可参考下面列出的相关资料，手动移植，或者使用本项目。

飞常准自建 ADS-B Windows 上传方案
https://blog.dextercai.com/2020-03-20-78.html

在 Arch Linux 下安装飞常准上传套件
https://blog.dextercai.com/2018-06-24-45.html

