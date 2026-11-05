package core

import (
	"blog_server/global"
	"bytes"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	red    = 31
	yellow = 33
	blue   = 34
	gray   = 37
)

type LogFormatter struct{}

type FileDateHook struct {
	file     *os.File
	logPath  string
	fileDate string //判断日期切换目录
	appName  string
}

func (hook FileDateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
func (hook FileDateHook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format("2006-01-02")
	line, _ := entry.String()
	if hook.fileDate == timer {
		hook.file.Write([]byte(line))
		return nil
	}
	// 时间不等
	hook.file.Close()
	os.MkdirAll(fmt.Sprintf("%s/%s", hook.logPath, timer), os.ModePerm)
	filename := fmt.Sprintf("%s/%s/%s.log", hook.logPath, timer, hook.appName)

	hook.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	hook.fileDate = timer
	hook.file.Write([]byte(line))
	return nil
}

func InitFile(logPath, appName string) {
	fileDate := time.Now().Format("2006-01-02")
	//创建目录
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), os.ModePerm)
	if err != nil {
		logrus.Error(err)
		return
	}

	filename := fmt.Sprintf("%s/%s/%s.log", logPath, fileDate, appName)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		logrus.Error(err)
		return
	}
	fileHook := FileDateHook{file, logPath, fileDate, appName}
	logrus.AddHook(&fileHook)
}

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 获取日志级别的颜色
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

	// 格式化日志输出
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.Caller != nil {
		// 获取文件名和行号
		funcName := entry.Caller.Function
		fileName := path.Base(entry.Caller.File)
		line := entry.Caller.Line
		level := entry.Level.String()
		fmt.Fprintf(b, "\x1b[%dm[%s] [%s][%s] [%s:%d] %s\x1b[0m\n", levelColor, level, timestamp, funcName, fileName, line, entry.Message)
	} else {
		fmt.Fprintf(b, "\x1b[%dm[%s]  %s\x1b[0m\n", levelColor, timestamp, entry.Message)
	}

	return b.Bytes(), nil
}

func InitLogger() *logrus.Logger {
	// 设置输出到标准输出
	logrus.SetOutput(os.Stdout)
	// 设置日志格式为自定义的LogFormatter
	logrus.SetFormatter(&LogFormatter{})
	// 开启返回调用者信息的功能（文件名和行号）
	logrus.SetReportCaller(true)
	// 设置日志级别为DebugLevel
	logrus.SetLevel(logrus.DebugLevel)
	logConfig := global.Config.Log
	InitFile(logConfig.Dir, logConfig.App)
	return logrus.StandardLogger()
}
