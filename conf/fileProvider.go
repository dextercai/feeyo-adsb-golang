package conf

import (
	"dextercai.com/feeyo-adsb-golang/log"
	"github.com/Unknwon/goconfig"
)

type TConf struct {
	IpDump1090   string
	PortDump1090 string
	FeeyoUrl     string
	UUID         string
}

var GlobalConfig TConf
var Config *goconfig.ConfigFile

func reloadConfig(confLoc string) {
	var err error
	Config, err = goconfig.LoadConfigFile(confLoc)
	if err != nil {
		log.Logger.Fatalf("[Fatal]:conf.ini配置文件不存在，请检查.")
	}
	GlobalConfig.UUID, err = Config.GetValue("config", "UUID")
	GlobalConfig.IpDump1090, err = Config.GetValue("config", "ip")
	GlobalConfig.PortDump1090, err = Config.GetValue("config", "port")
	GlobalConfig.FeeyoUrl, err = Config.GetValue("config", "url")

	if err != nil {
		log.Logger.Fatalf("[Fatal]:解析配置时出现错误")
	}
}
