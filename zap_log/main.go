package main

import (
	"go.uber.org/zap"
	"net/http"
	"zap_log/logger"
)

func main() {
	logger.InitLogger()
	defer logger.Logger.Sync()

	simpleHttpGet("www.baidu.com")
	simpleHttpGet("https://www.baidu.com")
}

func simpleHttpGet(url string) {
	response, err := http.Get(url)
	if err != nil {
		s := "mingcheng------------"
		logger.Error(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err),
		)
		logger.Errorf(
			"Error fetching url...",
			zap.String("url", url),
			zap.Error(err),
		)
		logger.Error("---------test")
		logger.Errorf("--------test %s", s)
	} else {
		s := "name-------"
		logger.Info(
			"Success---",
			zap.String("StatusCode", response.Status),
			zap.String("url", url),
		)
		logger.Infof(
			"Success---",
			zap.String("StatusCode", response.Status),
			zap.String("url", url),
		)
		logger.Infof("-------------test info")
		logger.Infof("-------------test info %s", s)
		response.Body.Close()
	}
}
