package main

import (
	"github.com/google/uuid"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	fileObj, _ := os.OpenFile(filepath.Join(filepath.Dir(os.Args[0]),"UUID.txt"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	UUID, _ := uuid.NewRandom()
	UuidString := strings.ReplaceAll(UUID.String(),"-","")[:16]
	println("生成的UUID为：", UuidString)
	_, _ = fileObj.Write([]byte(UuidString))
	println("UUID已保存至：", filepath.Join(filepath.Dir(os.Args[0]),"UUID.txt"))
	println("5秒后程序自动退出")
	time.Sleep(time.Second * 5)
}
