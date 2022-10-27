package tools

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

func InitLogs(Level logrus.Level, FilePath, FileFormat string, MaxAge, RotationTime time.Duration) (*rotatelogs.RotateLogs, error) {
	writer, err := LogsWriter(FilePath, FileFormat, MaxAge, RotationTime)
	if err != nil {
		return nil, err
	}
	logrus.SetLevel(Level)
	logrus.SetOutput(writer)
	return writer, nil
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
