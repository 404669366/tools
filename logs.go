package tools

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
	"time"
)

type formatter struct{}

func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format(DatetimeFormat)
	var file string
	var line int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		line = entry.Caller.Line
	}
	msg := fmt.Sprintf("%s %s %s line %d\n%s\n", timestamp, strings.ToUpper(entry.Level.String()), file, line, entry.Message)
	return []byte(msg), nil
}

func InitLogs(writer *rotatelogs.RotateLogs, Level logrus.Level) {
	logrus.SetLevel(Level)
	logrus.SetOutput(writer)
	logrus.SetFormatter(&formatter{})
}

//LogsWriter 日志切割writer FileFormat %Y-年 %m-月 %d-日 %H-时 %M-分 %S-秒
func LogsWriter(FilePath, FileFormat string, MaxAge, RotationTime time.Duration) (*rotatelogs.RotateLogs, error) {
	return rotatelogs.New(
		//设置切割后的文件名
		FilePath+"/"+FileFormat,
		//设置生成运行时日志文件软链
		rotatelogs.WithLinkName(FilePath+"/runtime.log"),
		//设置文件最大保存时间
		rotatelogs.WithMaxAge(MaxAge),
		//设置日志切割时间间隔
		rotatelogs.WithRotationTime(RotationTime),
	)
}
