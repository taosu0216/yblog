package main

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumber "gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"runtime"
)

var _ log.Logger = (*ZapLogger)(nil)

type ZapLogger struct {
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
}

func initLog() *ZapLogger {
	var coreArr []zapcore.Core
	var logger *zap.Logger
	// 获取编码器
	// NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig := zap.NewProductionEncoderConfig()
	// 指定时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//显示完整文件路径
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 日志级别
	// error级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	// info和debug级别,debug级别是最低的
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	_, currentFile, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(currentFile)
	currentDir = filepath.Dir(currentDir)
	currentDir = filepath.Dir(currentDir)

	// info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumber.Logger{
		Filename:   currentDir + "/log/info.log", // 日志文件存放目录，
		MaxSize:    200,                          // 文件大小限制,单位MB
		MaxBackups: 30,                           // 最大保留日志文件数量
		MaxAge:     7,                            // 日志文件保留天数
		Compress:   false,
	})
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority)

	// error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumber.Logger{
		Filename:   currentDir + "/log/error.log", // 日志文件存放目录，
		MaxSize:    200,                           // 文件大小限制,单位MB
		MaxBackups: 30,                            // 最大保留日志文件数量
		MaxAge:     7,                             // 日志文件保留天数
		Compress:   false,
	})
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority)

	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)

	// zap.AddCaller()为显示文件名和行号，可省略
	logger = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller(), zap.AddCallerSkip(2))
	sugar := logger.Sugar()
	logger.Info("init logger success")

	return &ZapLogger{Logger: logger, Sugar: sugar}

}

func (zapL *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		zapL.Logger.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}

	switch level {
	case log.LevelDebug:
		zapL.Logger.Debug("", data...)
	case log.LevelInfo:
		zapL.Logger.Info("", data...)
	case log.LevelWarn:
		zapL.Logger.Warn("", data...)
	case log.LevelError:
		zapL.Logger.Error("", data...)
	case log.LevelFatal:
		zapL.Logger.Fatal("", data...)
	}
	return nil
}

func (zapL *ZapLogger) Infof(s string, v ...interface{}) {
	zapL.Sugar.Infof(s, v...)
}

func (zapL *ZapLogger) Infow(s string, v ...interface{}) {
	zapL.Sugar.Infow(s, v...)
}

func (zapL *ZapLogger) Info(v ...interface{}) {
	zapL.Sugar.Info(v...)
}

func (zapL *ZapLogger) Debugf(s string, v ...interface{}) {
	zapL.Sugar.Debugf(s, v...)
}

func (zapL *ZapLogger) Debugw(s string, v ...interface{}) {
	zapL.Sugar.Debugw(s, v...)
}

func (zapL *ZapLogger) Debug(v ...interface{}) {
	zapL.Sugar.Debug(v...)
}

func (zapL *ZapLogger) Errorf(s string, v ...interface{}) {
	zapL.Sugar.Errorf(s, v...)
}
func (zapL *ZapLogger) Errorw(s string, v ...interface{}) {
	zapL.Sugar.Errorw(s, v...)
}

func (zapL *ZapLogger) Error(v ...interface{}) {
	zapL.Sugar.Error(v...)
}

func (zapL *ZapLogger) Fatalf(s string, v ...interface{}) {
	zapL.Sugar.Fatalf(s, v...)
}

func (zapL *ZapLogger) Fatalw(s string, v ...interface{}) {
	zapL.Sugar.Fatalw(s, v...)
}

func (zapL *ZapLogger) Fatal(v ...interface{}) {
	zapL.Sugar.Error(v...)
}
func (zapL *ZapLogger) Sync() {
	_ = zapL.Logger.Sync()
}
