package main

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"flag"
)

func init()  {
	profile := flag.String("profile", "test", "Environment profile")
	flag.Parse()

	if *profile == "dev" {
		log.SetFormatter(&log.TextFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000",
			FullTimestamp:   true,
		})
	} else {
		log.SetFormatter(&log.JSONFormatter{})
	}
}

func main() {
	// 生成一个router 返回处理结果
	log.Infof("start http server")
	http.ListenAndServe(":6767", NewHHTTPRouter())
}
