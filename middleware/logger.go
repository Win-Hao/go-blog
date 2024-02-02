package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		logger := logrus.New()
		filePath := "log/log"
		src, err1 := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err1 == nil {
			logger.Out = src
		} else {
			logger.Info("Fail to log to file,using default stderr")
		}
		logger.SetLevel(logrus.DebugLevel)
		logger.Out = os.Stdout
		logWriter, _ := rotatelogs.New(
			filePath+"%Y%m%d.log",
			rotatelogs.WithMaxAge(7*24*time.Hour),
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.FatalLevel: logWriter,
			logrus.PanicLevel: logWriter,
			logrus.TraceLevel: logWriter,
			logrus.WarnLevel:  logWriter,
		}
		hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:03",
		})
		logger.AddHook(hook)
		logger.Formatter = &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:03",
		}

		logger.SetReportCaller(true)
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))
		hostName, err2 := os.Hostname()
		if err2 != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		pathName := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"SpendTime":  spendTime,
			"HostName":   hostName,
			"StatusCode": statusCode,
			"ClientIp":   clientIp,
			"UserAgent":  userAgent,
			"DataSize":   dataSize,
			"Method":     method,
			"Path":       pathName,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
