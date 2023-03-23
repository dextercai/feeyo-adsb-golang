package main

import (
	"bytes"
	"compress/zlib"
	"dextercai.com/feeyo-adsb-golang/conf"
	"dextercai.com/feeyo-adsb-golang/constant"
	"dextercai.com/feeyo-adsb-golang/log"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var err error

const eachPackage = 8192
const thisLogFlag = "main"

var confCheckFunc = map[string]func(tConf conf.TConf) bool{
	"UUID为空": func(tConf conf.TConf) bool {
		return tConf.UUID == ""
	},
	"UUID不满足16位标准长度": func(tConf conf.TConf) bool {
		return len(tConf.UUID) != 16
	},
	"Dump1090地址配置为空": func(tConf conf.TConf) bool {
		return tConf.IpDump1090 == ""
	},
	"Dump1090端口配置为空": func(tConf conf.TConf) bool {
		return tConf.PortDump1090 == ""
	},
	"FeeyoUrl配置未空": func(tConf conf.TConf) bool {
		return tConf.FeeyoUrl == ""
	},
}

func main() {
	fmt.Println("项目地址: https://github.com/dextercai/feeyo-adsb-golang")
	fmt.Printf("当前版本：%s，编译时间：%s", constant.Version, constant.BuildTime)
	fmt.Println("")
	fmt.Println("敬告: 请不要尝试将相关电波数据传送至FR24, RadarBox, FA等境外平台, 这将严重违反无线电管理条例以及国家安全法!")
	fmt.Println("=============================================================================================")
	conf.ParseConf()

	for item := range confCheckFunc {
		fun := confCheckFunc[item]
		if fun(conf.GlobalConfig) {
			log.Logger.Fatalf("配置检查失败：%s", item)
		}
	}
	for {
		dump1090Conn, err := net.Dial("tcp", conf.GlobalConfig.IpDump1090+":"+conf.GlobalConfig.PortDump1090)
		if err != nil {
			log.Logger.Printf("[%s]:%s\t%s", thisLogFlag, "连接到Dump1090失败", err.Error())
			log.Logger.Printf("[%s]:%s", thisLogFlag, "15秒后重试")
			time.Sleep(15 * time.Second)
			continue
		} else {
			log.Logger.Printf("[%s]:%s", thisLogFlag, "连接到Dump1090成功")
		}
		var buf [eachPackage]byte
		for {
			read, err := dump1090Conn.Read(buf[0:])
			if err != nil {
				log.Logger.Printf("[%s]:%s\t%s", thisLogFlag, "读取数据错误", err.Error())
				_ = dump1090Conn.Close()
				log.Logger.Printf("[%s]:%s", thisLogFlag, "已断开连接，正尝试重连")
				break
			} else {
				if buf[read-1] == 10 {
					sendMessage(buf[0:read])
				}
			}
		}
	}

}

func sendMessage(line []byte) {
	sourceData := base64.StdEncoding.EncodeToString(DoZlibCompress(line))
	postValue := url.Values{}
	postValue.Set("from", conf.GlobalConfig.UUID)
	postValue.Set("code", sourceData)
	resp, err := http.Post(conf.GlobalConfig.FeeyoUrl, "application/x-www-form-urlencoded", strings.NewReader(postValue.Encode()))
	if err != nil {
		log.Logger.Printf("[%s]:%s\t%s", thisLogFlag, "上传错误", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Logger.Printf("[%s]:%s\t%s", thisLogFlag, "上传错误", err.Error())
	} else {
		log.Logger.Printf("[%s]:%s\t%s", thisLogFlag, "上传成功", string(body))
	}
}

func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	_, _ = w.Write(src)
	_ = w.Close()
	return in.Bytes()
}

/*
func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	_, _ = io.Copy(&out, r)
	return out.Bytes()
}

*/
