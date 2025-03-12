/**
 * @Author: lixiang
 * @Date: 2025/3/11 11:02
 * @Description: logrus.go
 */

package core

import (
	"bytes"
	"fmt"
	"gin-admin-api/config"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

const (
	red    = 31
	yellow = 33
	blue   = 34
	gray   = 37
)

type LogFormatter struct{}

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	//根据不同的level 展示不同的颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel:
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

	log := config.Config.Logger
	//自定义日期格式
	format := entry.Time.Format(entry.Time.Format("2006-01-02 15:04:05"))
	location := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)

	if entry.HasCaller() {
		//自定义文件路径
		function := entry.Caller.Function
		_, err := fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s: %s\n", log.Prefix, format, levelColor, entry.Level, location, function, entry.Message)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := fmt.Fprintf(b, "%s[%s] \x1b[%dm[%s]\x1b[0m: %s\n]", log.Prefix, format, levelColor, entry.Level, entry.Message)
		if err != nil {
			return nil, err
		}
	}
	return b.Bytes(), nil

}

func InitLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(config.Config.Logger.ShowLine)
	logger.SetFormatter(&LogFormatter{})
	level, err := logrus.ParseLevel(config.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)
	InitDefaultLogger()
	return logger
}
func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(config.Config.Logger.ShowLine)
	logrus.SetFormatter(&LogFormatter{})
	level, err := logrus.ParseLevel(config.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}
