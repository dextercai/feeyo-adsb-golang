package core

import (
	"context"
	"dextercai.com/feeyo-adsb-golang/util"
	"encoding/base64"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type FeeyoCompressSender struct {
	c        chan []byte
	logEntry *logrus.Entry

	ctx       context.Context
	ctxCancel context.CancelFunc

	UUID     string
	FeeyoUrl string
}

func NewFeeyoCompressSender(ctx context.Context, logEntry *logrus.Entry, UUID string, feeyoUrl string) *FeeyoCompressSender {
	childCtx, ctxCancel := context.WithCancel(ctx)
	ch := make(chan []byte)
	return &FeeyoCompressSender{
		c:         ch,
		logEntry:  logEntry,
		ctx:       childCtx,
		UUID:      UUID,
		FeeyoUrl:  feeyoUrl,
		ctxCancel: ctxCancel,
	}
}

func (f *FeeyoCompressSender) Run() {
	defer f.logEntry.Infof("收到退出信号")
	for {
		select {
		case <-f.ctx.Done():
			return
		case buf, ok := <-f.c:
			if !ok {
				return
			}
			f.sendMsg(buf)
		}
	}
}

func (f *FeeyoCompressSender) GetSendChan() chan []byte {
	return f.c
}

func (f *FeeyoCompressSender) sendMsg(line []byte) {
	postValue := url.Values{}
	in := util.DoZlibCompress(line)
	postValue.Set("from", f.UUID)
	postValue.Set("code", base64.StdEncoding.EncodeToString(in.Bytes()))
	resp, err := http.Post(f.FeeyoUrl, "application/x-www-form-urlencoded", strings.NewReader(postValue.Encode()))
	if err != nil {
		f.logEntry.WithError(err).Errorf("发送上传请求时错误")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		f.logEntry.WithError(err).Errorf("读取上传请求返回值时出现错误")
		return
	}
	f.logEntry.Debugf("上传成功：%s", body)
}

func (f *FeeyoCompressSender) Stop() {
	f.ctxCancel()
}
