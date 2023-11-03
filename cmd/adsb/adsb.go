package main

import (
	"dextercai.com/feeyo-adsb-golang/conf"
	"dextercai.com/feeyo-adsb-golang/log"
	"os"
	"os/signal"
)

func main() {
	conf.InitConfig()
	currentConfig := conf.ReadConfig()
	log.InitLog(currentConfig.LogLevel, currentConfig.LogPath, currentConfig.LogFile,
		currentConfig.LogRotationTime, currentConfig.LogMaxAge, currentConfig.LogRotationSize, currentConfig.LogRotationCount)
	log.Logger.Debugf("配置读取: %#v", currentConfig)

	log.Logger.Warnf("将使用UUID: %s", currentConfig.FeeyoUUID)

	//wg := sync.WaitGroup{}
	//wg.Add(1)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Kill, os.Interrupt)

	sig := <-quit
	log.Logger.WithField("scope", "process").Infof("收到信号: %s", sig.String())

	//grpcServer.GracefulStop()
	//wg.Wait()
	log.Logger.WithField("scope", "process").Info("按预期关闭")
}
