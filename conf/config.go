package conf

import (
	"dextercai.com/feeyo-adsb-golang/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var ConfigFile string
var logEntry = log.Logger.WithField("scope", "conf")

type Config struct {
	Dump1090Host string
	Dump1090Port int

	LogLevel string
	LogPath  string
	LogFile  string

	LogRotationTime int64
	LogMaxAge       int64

	LogRotationSize  int64
	LogRotationCount uint

	FeeyoUrl  string
	FeeyoUUID string
}

func init() {
	pflag.String(Conf, defaultConfigMap[Conf].(string), "配置文件位置")
	pflag.String(Dump1090Host, defaultConfigMap[Dump1090Host].(string), "dump1090服务地址")
	pflag.Int(Dump1090Port, defaultConfigMap[Dump1090Port].(int), "dump1090服务端口")
	pflag.String(LogLevel, defaultConfigMap[LogLevel].(string), "日志等级")
	pflag.String(LogPath, defaultConfigMap[LogPath].(string), "dump1090服务端口")
	pflag.String(LogFile, defaultConfigMap[LogFile].(string), "dump1090服务端口")

	pflag.Int(LogRotationTime, defaultConfigMap[LogRotationTime].(int), "dump1090服务端口")
	pflag.Int(LogMaxAge, defaultConfigMap[LogMaxAge].(int), "dump1090服务端口")
	pflag.Int(LogRotationSize, defaultConfigMap[LogRotationSize].(int), "dump1090服务端口")
	pflag.Uint(LogRotationCount, defaultConfigMap[LogRotationCount].(uint), "dump1090服务端口")
	pflag.String(FeeyoUrl, defaultConfigMap[FeeyoUrl].(string), "dump1090服务端口")
	pflag.String(FeeyoUUID, defaultConfigMap[FeeyoUUID].(string), "dump1090服务端口")
	pflag.Parse()
}

const (
	Conf         = "conf"
	Dump1090Host = "dump1090.host"
	Dump1090Port = "dump1090.port"

	LogLevel         = "log.level"
	LogPath          = "log.path"
	LogFile          = "log.file"
	LogRotationTime  = "log.rotation_time"
	LogMaxAge        = "log.max_age"
	LogRotationSize  = "log.rotation_size"
	LogRotationCount = "log.rotation_count"

	FeeyoUrl  = "feeyo.url"
	FeeyoUUID = "feeyo.uuid"
)

var defaultConfigMap = map[string]any{
	Conf:             "./config.ini",
	Dump1090Host:     "127.0.0.1",
	Dump1090Port:     30003,
	LogLevel:         log.LogLevelInfo,
	LogPath:          "./logs/",
	LogFile:          "feeyo-adsb-golang.log",
	LogRotationTime:  86400,
	LogMaxAge:        604800,
	LogRotationSize:  10 << 20,
	LogRotationCount: uint(0),
	FeeyoUrl:         "http://adsb.feeyo.com/adsb/ReceiveCompressADSB.php",
	FeeyoUUID:        "",
}

func InitConfig() {
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		logEntry.WithError(err).Fatal("读取命令行参数时出错")
	}
	ConfigFile = viper.GetString(Conf)
	viper.SetConfigFile(ConfigFile)

	if err := viper.ReadInConfig(); err != nil {
		logEntry.Errorf("配置文件读取失败: %s", err)
	} else {
		logEntry.Infof("使用配置文件: %s", viper.ConfigFileUsed())
	}

	for s, a := range defaultConfigMap {
		viper.SetDefault(s, a)
	}
}

func ReadConfig() Config {
	cfg := Config{
		Dump1090Host:     viper.GetString(Dump1090Host),
		Dump1090Port:     viper.GetInt(Dump1090Port),
		LogLevel:         viper.GetString(LogLevel),
		LogFile:          viper.GetString(LogFile),
		LogPath:          viper.GetString(LogPath),
		LogRotationTime:  viper.GetInt64(LogRotationTime),
		LogMaxAge:        viper.GetInt64(LogMaxAge),
		LogRotationCount: viper.GetUint(LogRotationCount),
		LogRotationSize:  viper.GetInt64(LogRotationSize),

		FeeyoUrl:  viper.GetString(FeeyoUrl),
		FeeyoUUID: viper.GetString(FeeyoUUID),
	}
	if cfg.FeeyoUUID == "" {
		logEntry.Fatal("feeyo.uuid 当前配置为空串，不可启动。")
	} else if len(cfg.FeeyoUUID) != 16 {
		logEntry.Warnf("feeyo.uuid 长度不为16，可能不被服务器接受。")
	}
	if cfg.FeeyoUrl == "" {
		logEntry.Fatal("feeyo.url 当前配置为空串，不可启动。")
	}
	return cfg
}
