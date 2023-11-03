package main

import (
	"context"
	"dextercai.com/feeyo-adsb-golang/conf"
	"dextercai.com/feeyo-adsb-golang/constant"
	"dextercai.com/feeyo-adsb-golang/core"
	"dextercai.com/feeyo-adsb-golang/log"
	"os"
	"os/signal"
	"sync"
)

func main() {
	log.Logger.Warnf("github.com/dextercai/feeyo-adsb-golang")
	log.Logger.Warnf("Version: %s, BuildTime: %s", constant.Version, constant.BuildTime)
	log.Logger.Warnf("根据《中华人民共和国国家安全法》第七十七条；《中华人民共和国无线电管理条例》第五十五条、七十五条。")
	log.Logger.Warnf("任何单位或者个人不得向境外组织或者个人提供涉及国家安全的境内电波参数资料")

	conf.InitConfig()
	currentConfig := conf.ReadConfig()
	log.InitLog(currentConfig.LogLevel, currentConfig.LogPath, currentConfig.LogFile,
		currentConfig.LogRotationTime, currentConfig.LogMaxAge, currentConfig.LogRotationSize, currentConfig.LogRotationCount)
	log.Logger.Debugf("配置读取: %#v", currentConfig)

	log.Logger.Warnf("将使用UUID: %s", currentConfig.FeeyoUUID)

	wg := sync.WaitGroup{}
	ctx := context.Background()
	sender := core.NewFeeyoCompressSender(ctx, log.Logger.WithField("scope", "feeyo-sender"), currentConfig.FeeyoUUID, currentConfig.FeeyoUrl)

	wg.Add(1)
	go func() {
		sender.Run()
		wg.Done()
	}()

	receiver := core.NewDump1090Receiver(ctx, currentConfig.Dump1090Host,
		currentConfig.Dump1090Port,
		8192,
		log.Logger.WithField("scope", "dump1090-receiver"), sender.GetSendChan(),
	)
	wg.Add(1)
	go func() {
		receiver.Run()
		wg.Done()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Kill, os.Interrupt)

	sig := <-quit
	log.Logger.WithField("scope", "process").Infof("收到信号: %s", sig.String())

	receiver.Stop()
	sender.Stop()
	wg.Wait()
	log.Logger.WithField("scope", "process").Info("按预期关闭")
}
