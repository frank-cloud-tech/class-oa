// Package utils
// @Description:
package utils

import (
	"github.com/frank-cloud-tech/class-oa/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

var Log *logrus.Logger

func init() {
	Log = NewLoggerWithRotate()
}

//NewLoggerWithRotate 返回日志对象, 日志路径为access.log, 且会按照日期切割, 单例模式
func NewLoggerWithRotate() *logrus.Logger {
	if Log != nil {
		return Log
	}

	pathMap := lfshook.WriterMap{}

	logPath := path.Join(config.Config.GetLogDir(), "access.log")
	writer, _ := rotatelogs.New(
		logPath+".%Y-%m-%d",
		rotatelogs.WithLinkName(logPath),             // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(24*3600*time.Second),   // 文件最大保存时间
		rotatelogs.WithRotationTime(600*time.Second), // 日志切割时间间隔
	)

	levels := logrus.AllLevels
	for _, level := range levels {
		pathMap[level] = writer
	}

	Log = logrus.New()
	Log.SetOutput(os.Stdout)
	//Log.SetOutput(ioutil.Discard)
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{TimestampFormat: "2022-08-07 15:04:05"},
	))
	return Log
}
