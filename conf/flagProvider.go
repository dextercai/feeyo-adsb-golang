package conf

import (
	"dextercai.com/feeyo-adsb-golang/log"
	"flag"
)

var useFile bool
var confFile string

func ParseConf() {
	flag.StringVar(&GlobalConfig.UUID, "uuid", "", "UUID 16位")
	flag.StringVar(&GlobalConfig.IpDump1090, "ip", "127.0.0.1", "dump1090服务IP")
	flag.StringVar(&GlobalConfig.PortDump1090, "port", "30003", "dump1090服务端口")
	flag.StringVar(&GlobalConfig.FeeyoUrl, "feeyo-url", "https://adsb.feeyo.com/adsb/ReceiveCompressADSB.php", "飞常准接口地址")
	flag.BoolVar(&useFile, "use-file", true, "是否使用conf文件作为配置来源")
	flag.StringVar(&confFile, "conf", "./conf.ini", "conf文件位置")

	flag.Parse()

	if useFile {
		reloadConfig(confFile)
		log.Logger.Printf("[%s]:%s%s", "CONF", "使用文件载入配置，文件位置:", confFile)
	} else {
		log.Logger.Printf("[%s]:%s", "CONF", "使用命令行参数载入配置")
	}
}
