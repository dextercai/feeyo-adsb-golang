package main

import (
	"DexterCai/Feeyo_adsb/config"
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)
var UUID, IpDump1090, PortDump1090,FeeyoUrl string
func main() {
	println("敬告：请不要尝试将相关电波数据传送至FR24，RadarBox，FA等境外平台，这将严重违反无线电管理条例以及国家安全法！")
	var err error
	UUID, err = config.Config.GetValue("config", "UUID")
	IpDump1090, err = config.Config.GetValue("config", "ip")
	PortDump1090, err = config.Config.GetValue("config", "port")
	FeeyoUrl, err = config.Config.GetValue("config", "url")
	if UUID == "" || len(UUID) != 16 || IpDump1090 == "" || PortDump1090 == "" || FeeyoUrl == ""|| err != nil{
		println("配置错误")
		return
	}
	dump1090Conn, err := net.Dial("tcp", IpDump1090 + ":" + PortDump1090)
	if err != nil {
		println(time.Now().Format("2006-01-02 15:04:05"), "\t连接到Dump1090失败\t", err.Error())
		return
	}else{
		println(time.Now().Format("2006-01-02 15:04:05"), "\t连接成功\t")
	}
	var buf [1024]byte
	for {
		read, err := dump1090Conn.Read(buf[0:])
		if err != nil {
			println(time.Now().Format("2006-01-02 15:04:05"), "\t读取错误\t", err.Error())
			_ = dump1090Conn.Close()
			println(time.Now().Format("2006-01-02 15:04:05"), "\t断开链接\t")
			dump1090Conn, err = net.Dial("tcp", IpDump1090 + ":" + PortDump1090)
			println(time.Now().Format("2006-01-02 15:04:05"), "\t尝试重连\t")
			continue
		}else{
			if buf[read-1] == 10 {
				sendMessage(buf[0:read])
			}
		}
	}
}

func sendMessage(line []byte){
	sourceData := base64.StdEncoding.EncodeToString(DoZlibCompress(line))
	postValue := url.Values{}
	postValue.Set("from", UUID)
	postValue.Set("code", sourceData)
	resp, err := http.Post(FeeyoUrl,"application/x-www-form-urlencoded",strings.NewReader(postValue.Encode()))
	if err != nil {
		println(time.Now().Format("2006-01-02 15:04:05"), "\t上传错误\t", err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		println(time.Now().Format("2006-01-02 15:04:05"), "\t上传错误\t", err.Error())
	}else{
		print("\r",time.Now().Format("2006-01-02 15:04:05"), "\t上传成功\t", string(body))
	}
}

func DoZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(src)
	w.Close()
	return in.Bytes()
}


func DoZlibUnCompress(compressSrc []byte) []byte {
	b := bytes.NewReader(compressSrc)
	var out bytes.Buffer
	r, _ := zlib.NewReader(b)
	io.Copy(&out, r)
	return out.Bytes()
}