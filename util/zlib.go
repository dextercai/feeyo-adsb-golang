package util

import (
	"bytes"
	"compress/zlib"
)

func DoZlibCompress(src []byte) (in bytes.Buffer) {
	w := zlib.NewWriter(&in)
	_, _ = w.Write(src)
	_ = w.Close()
	return
}
