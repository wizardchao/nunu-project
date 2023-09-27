package main

import (
	"fmt"
	"go.uber.org/zap"
	"nunu-project/pkg/config"
	"nunu-project/pkg/http"
	"nunu-project/pkg/log"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	app, cleanup, err := newApp(conf, logger)
	if err != nil {
		panic(err)
	}
	logger.Info("server start", zap.String("host", "http://127.0.0.1:"+conf.GetString("http.port")))

	http.Run(app, fmt.Sprintf(":%d", conf.GetInt("http.port")))
	defer cleanup()

}
