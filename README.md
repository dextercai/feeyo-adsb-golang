## 综述

使用Golang编写的飞常准ADSB上传程序，适用于各个支持Golang编译的平台。

## 使用说明

本项目分为两部分组成，UUID_Gen以及Feeyo-adsb。

注：UUID_Gen中的UUID生成器基于UUID Version 4。		

由于本项目不包括Dump1090，也不限制SBS服务是否运行在本机，因此你可能需要首先安装Dump1090，具体细节可自行搜索，当然你也可以在本项目提一个Issue，我将很乐意为你解答。



你需要编辑Feeyo-adsb下的conf.ini文件

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

如果使用其他Linux发行版，可参考我的[这篇博文](https://blog.dextercai.com/2018-06-a04d2416/)，手动移植，或者使用本项目。

