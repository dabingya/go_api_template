package config

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

func InitLog(filePath string) error {
	//logName := viper.GetString("log_file")
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if logf, err := rotatelogs.New(
		filePath+"/logs/info.%Y%m%d.log",
		rotatelogs.WithLinkName(filePath+"/logs/info.log"),
		rotatelogs.WithMaxAge(time.Hour*24*365),
		rotatelogs.WithRotationTime(time.Hour*24),
	); err != nil {
		return err
	} else {
		logrus.SetOutput(logf)
		logrus.Info("日志初始化完成")
	}

	if logr, err := rotatelogs.New(
		filePath+"/logs/run.%Y%m%d.log",
		rotatelogs.WithLinkName(filePath+"/logs/run.log"),
		rotatelogs.WithMaxAge(time.Hour*24*365),
		rotatelogs.WithRotationTime(time.Hour*24),
	); err != nil {
		return err
	} else {
		gin.DefaultWriter = io.MultiWriter(logr, os.Stdout)
		gin.DefaultErrorWriter = io.MultiWriter(logr, os.Stderr)
		gin.DisableConsoleColor()
	}

	return nil
}
