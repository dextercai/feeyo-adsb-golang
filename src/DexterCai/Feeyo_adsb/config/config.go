package config

import "github.com/Unknwon/goconfig"

var Config *goconfig.ConfigFile
func init() {
	var err error
	Config, err = goconfig.LoadConfigFile("conf.ini")
	if err != nil{
		panic("Cant load config file from conf.ini")
	}

}