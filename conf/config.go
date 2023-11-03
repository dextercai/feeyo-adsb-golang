package conf

import (
	"dextercai.com/feeyo-adsb-golang/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"time"
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
	pflag.String(Conf, defaultConfigMap[Conf].(string), "help message for flagname")
	pflag.String(Dump1090Host, "", "help message for flagname")
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
	Conf:             "./etc/config.ini",
	Dump1090Host:     "127.0.0.1",
	Dump1090Port:     "30003",
	LogLevel:         log.LogLevelInfo,
	LogPath:          "./logs/",
	LogFile:          "feeyo-adsb-golang.log",
	LogRotationTime:  int64(24 * time.Hour),
	LogMaxAge:        int64(7 * 24 * time.Hour),
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
