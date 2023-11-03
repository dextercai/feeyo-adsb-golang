package main

import (
	"dextercai.com/feeyo-adsb-golang/conf"
	"fmt"
)

func main() {
	conf.InitConfig()

	fmt.Printf("%#v", conf.ReadConfig())
}
