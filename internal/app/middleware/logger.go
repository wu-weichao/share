package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"share/configs"
	"time"
)

func Log() gin.HandlerFunc {
	// log file path
	logPath := configs.Log.Path
	if _, err := os.Stat(logPath); err != nil {
		// mkdir log directory
		if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
			log.Printf("[fail] cannot make dir %s", logPath)
		}
	}
	logFileName := fmt.Sprintf("%s-%s.%s", "share", time.Now().Format("2006-01-02"), "log")
	// if log file not exists, make file
	logFile := path.Join(logPath, logFileName)
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 644)
	if err != nil {
		log.Printf("[fail] open logger file error: %v", err)
	}

	// logrus init
	logger := logrus.New()
	logger.Out = f
	logger.Level = logrus.Level(configs.Log.Level)
	logger.Formatter = &logrus.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05",
		DisableHTMLEscape: true,
	}

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime).Milliseconds()
		// 记录日志
		logger.WithFields(logrus.Fields{
			"code":         c.Writer.Status(),                // status code
			"latency_time": fmt.Sprintf("%dms", latencyTime), // execution time
			"method":       c.Request.Method,                 // request method
			"url":          c.Request.RequestURI,             // request uri
			"params":       c.Request.PostForm,               // request form
			"ip":           c.ClientIP(),                     // client ip
		}).Info()
	}
}
