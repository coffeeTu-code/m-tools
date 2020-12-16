package lumberjack_helper

import (
	"github.com/natefinch/lumberjack"
	"io"
)

func NewLumberjack(filename string, maxAge, rotationCount int) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,      //日志文件的位置
		MaxAge:     maxAge,        //保留旧文件的最大天数
		MaxSize:    100,           //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: rotationCount, //保留旧文件的最大个数
		Compress:   false,         //是否压缩/归档旧文件
	}
}
