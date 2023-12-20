package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"os"
	"path"
)

// color
const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type LogFormatter struct{}

// Format instantiate Format(entry *logrus.Entry) ([]byte, error) interface
// 将 Formatter 方法实例化到 LogFormatter 结构体上
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Use different corresponding to different level
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	log := global.Config.Logger
	// Date format
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// file path
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		// log format
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", log.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s\n", log.Prefix, timestamp, levelColor, entry.Level, entry.Message)
	}

	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	mLog := logrus.New()                                // 新建一个实例
	mLog.SetOutput(os.Stdout)                           // 设置输出类型
	mLog.SetReportCaller(global.Config.Logger.ShowLine) // 开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{})                  // 设置自定义的Formatter
	mLog.SetLevel(logrus.DebugLevel)                    // 设置最低的log level
	InitDefaultLogger()
	return mLog
}

func InitDefaultLogger() {
	// global log
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(global.Config.Logger.ShowLine)
	logrus.SetFormatter(&LogFormatter{})

	level, err := logrus.ParseLevel(global.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}
