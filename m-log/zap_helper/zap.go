package zap_helper

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

//zap提供了几个快速创建logger的方法，zap.NewExample()、zap.NewDevelopment()、zap.NewProduction()，还有高度定制化的创建方法zap.New()。
//创建前 3 个logger时，zap会使用一些预定义的设置，它们的使用场景也有所不同。
//Example适合用在测试代码中，Development在开发环境中使用，Production用在生成环境。

func NewZapExample() *zap.Logger {
	return zap.NewExample()
}

func NewZapDevelopment() (*zap.Logger, error) {
	return zap.NewDevelopment()
}

func NewZapProduction() (*zap.Logger, error) {
	return zap.NewProduction()
}

func NewZap(writer io.Writer) *zap.SugaredLogger {
	writeSyncer := zapcore.AddSync(writer)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	//zap提供了便捷的方法SugarLogger，可以使用printf格式符的方式。
	//SugaredLogger的使用比Logger简单，只是性能比Logger低 50% 左右，可以用在非热点函数中。
	sugarLogger := logger.Sugar()

	return sugarLogger
}
