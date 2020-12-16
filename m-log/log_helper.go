package m_log

import (
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"io"
	"m-tools/m-log/filerotatelogs_helper"
	"m-tools/m-log/logrus_helper"
	"m-tools/m-log/lumberjack_helper"
	"m-tools/m-log/zap_helper"
)

type logHelper struct {
	loghelper *logrus.Logger
	loghot    *zap.SugaredLogger
}

func NewLogHelper() *logHelper {
	var writer io.Writer
	{
		writer = lumberjack_helper.NewLumberjack("runtime.log", 24, 10)
		//or
		writer = filerotatelogs_helper.NewFileRotatelogs("runtime.log", 24, 60, 10)
	}

	return &logHelper{
		loghelper: logrus_helper.NewLogrus(writer),
		loghot:    zap_helper.NewZap(writer),
	}
}

func (lh *logHelper) Debug(args ...interface{}) {
	if lh != nil && lh.loghelper != nil {
		lh.loghelper.Debug(args)
	}
}

func (lh *logHelper) Info(args ...interface{}) {
	if lh != nil && lh.loghelper != nil {
		lh.loghelper.Info(args)
	}
}

func (lh *logHelper) Warn(args ...interface{}) {
	if lh != nil && lh.loghelper != nil {
		lh.loghelper.Warn(args)
	}
}

func (lh *logHelper) Error(args ...interface{}) {
	if lh != nil && lh.loghelper != nil {
		lh.loghelper.Error(args)
	}
}

func (lh *logHelper) Fatal(args ...interface{}) {
	if lh != nil && lh.loghelper != nil {
		lh.loghelper.Fatal(args)
	}
}

func (lh *logHelper) Hot(args ...interface{}) {
	if lh != nil && lh.loghot != nil {
		lh.loghot.Info(args)
	}
}
