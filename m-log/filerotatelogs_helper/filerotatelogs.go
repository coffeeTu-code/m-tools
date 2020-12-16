package filerotatelogs_helper

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"log"
	"time"
)

func NewFileRotatelogs(filename string, maxAge, rotationTime, rotationCount int) io.Writer {
	if maxAge == 0 {
		maxAge = 24
	}
	if rotationTime == 0 {
		rotationTime = 60
	}
	if rotationCount == 0 {
		rotationCount = 10
	}
	writer, err := rotatelogs.New(
		filename+".%Y-%m-%d-%H",
		rotatelogs.WithLinkName(filename),                                    // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Duration(maxAge)*time.Hour),               // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Duration(rotationTime)*time.Minute), // 日志切割时间间隔
		rotatelogs.WithRotationCount(uint(rotationCount)),                    // 保存日志个数，默认 7，不能与 MaxAge 同时设置
	)
	if err != nil {
		log.Println("rotatelogs.New ", filename, "error=", err)
	}
	return writer
}
