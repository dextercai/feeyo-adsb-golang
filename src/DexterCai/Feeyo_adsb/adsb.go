package main

import (
	"DexterCai/Feeyo_adsb/config"
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/textproto"
	"net/url"
	"strings"
	"time"
)
var UUID, IpDump1090, PortDump1090,FeeyoUrl string
func main() {
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
		fmt.Println("连接到Dump1090失败", err.Error())
		return
	}
	pipeline := textproto.NewConn(dump1090Conn)
	for {
		message, err := pipeline.ReadLineBytes()
		if err != nil {
			println("读取错误", err.Error())
			_ = dump1090Conn.Close()
			return
		}
		go sendMessage(message)

	}
}

func sendMessage(line []byte){
	sourceData := base64.StdEncoding.EncodeToString(DoZlibCompress(line))
	postValue := url.Values{}
	postValue.Set("from", UUID)
	postValue.Set("code", sourceData)
	resp, err := http.Post(FeeyoUrl,"application/x-www-form-urlencoded",strings.NewReader(postValue.Encode()))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if err != nil{
		println(time.Now().String(), "上传错误", err.Error())
	}else{
		println(time.Now().String(), "上传成功", string(body))
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