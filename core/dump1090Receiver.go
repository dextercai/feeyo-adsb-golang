package core

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

type Dump1090Receiver struct {
	conn             net.Conn
	host             string
	port             int
	packageBatchSize int
	logEntry         *logrus.Entry
	c                chan []byte
	ctx              context.Context
	ctxCancel        context.CancelFunc
}

func NewDump1090Receiver(conn net.Conn, host string, port int, packageBatchSize int, logEntry *logrus.Entry, c chan []byte, ctx context.Context) *Dump1090Receiver {
	childCtx, cancelFunc := context.WithCancel(ctx)

	return &Dump1090Receiver{
		conn:             conn,
		host:             host,
		port:             port,
		packageBatchSize: packageBatchSize,
		logEntry:         logEntry,
		c:                c,
		ctx:              childCtx,
		ctxCancel:        cancelFunc,
	}
}

func (r *Dump1090Receiver) Run() {
	var err error
	defer r.logEntry.WithError(err).Infof("收到退出信号")
	for {
		if r.ctx.Err() != nil {
			return
		}
		r.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", r.host, r.port))
		if err != nil {
			r.logEntry.WithError(err).Error("连接到dump1090失败，将在10秒后重试。")
			time.Sleep(10 * time.Second)
			continue
		} else {
			r.logEntry.Warn("连接到Dump1090成功")
		}
		var buf = make([]byte, r.packageBatchSize)
		for {
			select {
			case <-r.ctx.Done():
				_ = r.conn.Close()
				return
			default:
				readLen, err := r.conn.Read(buf[0:])
				if err != nil {
					r.logEntry.WithError(err).Error("读取数据错误")
					err = r.conn.Close()
					if err != nil {
						r.logEntry.WithError(err).Error("关闭上一次dump1090链接时出错")
					}
					r.logEntry.Warn("断开连接，尝试重连")
					break
				} else {
					if buf[readLen-1] == 0x0A {
						r.logEntry.Debugf("收到数据长度: %d", len(buf))
						r.logEntry.Tracef("Hex: %x", buf)
						r.logEntry.Tracef("Str: %s", buf)
						r.c <- buf[0:readLen]
					}
				}
			}
		}
	}
}

func (r *Dump1090Receiver) Stop() {
	r.ctxCancel()
}
