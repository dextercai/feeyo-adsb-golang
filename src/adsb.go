package main

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"github.com/Unknwon/goconfig"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	ipDump1090   string
	portDump1090 string
	feeyoUrl     string
	UUID         string
)
var Config *goconfig.ConfigFile

func main() {
	initConfig()
	fmt.Println("本项目地址：https://github.com/dextercai/feeyo-adsb-golang")
	fmt.Println("温馨提示：二次分发时请遵守GPL3.0协议")
	fmt.Println("=============================================================================================")
	fmt.Println("敬告：请不要尝试将相关电波数据传送至FR24，RadarBox，FA等境外平台，这将严重违反无线电管理条例以及国家安全法！")

	var err error
	UUID, err = Config.GetValue("config", "UUID")
	ipDump1090, err = Config.GetValue("config", "ip")
	portDump1090, err = Config.GetValue("config", "port")
	feeyoUrl, err = Config.GetValue("config", "url")
	if UUID == "" || len(UUID) != 16 || ipDump1090 == "" || portDump1090 == "" || feeyoUrl == "" || err != nil {
		println("配置错误")
		os.Exit(0)
	}
	for {
		dump1090Conn, err := net.Dial("tcp", ipDump1090+":"+portDump1090)
		if err != nil {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "\t连接到Dump1090失败\t", err.Error())
			time.Sleep(15 * time.Second)
			continue
		} else {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "\t连接到Dump1090成功\t")
		}
		var buf [1024]byte
		for {
			read, err := dump1090Conn.Read(buf[0:])
			if err != nil {
				fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "\t读取数据错误\t", err.Error())
				_ = dump1090Conn.Close()
				fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "\t已断开连接，正尝试重连\t")
				break
			} else {
				if buf[read-1] == 10 {
					sendMessage(buf[0:read])
				}
			}
		}
	}

}
func initConfig() {
	var err error
	Config, err = goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		fmt.Println("conf.ini配置文件不存在，请检查.")
		os.Exit(0)
	}
}
func sendMessage(line []byte) {
	sourceData := base64.StdEncoding.EncodeToString(DoZlibCompress(line))
	postValue := url.Values{}
	postValue.Set("from", UUID)
	postValue.Set("code", sourceData)
	resp, err := http.Post(feeyoUrl, "application/x-www-form-urlencoded", strings.NewReader(postValue.Encode()))
	if err != nil {
		println(time.Now().Format("2006-01-02 15:04:05"), "\t上传错误\t", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(time.Now().Format("2006-01-02 15:04:05"), "\t上传错误\t", err.Error())
	} else {
		print("\r", time.Now().Format("2006-01-02 15:04:05"), "\t上传成功\t", string(body))
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
