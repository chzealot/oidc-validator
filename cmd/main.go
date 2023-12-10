package main

import (
	"github.com/chzealot/gobase/logger"
	"oidc-validator/internal/api"
)

func main() {
	if err := logger.InitWithConfig(logger.Config{
		AppName:   "oidc-validator",
		DebugMode: logger.DebugModeFromEnv,
	}); err != nil {
		panic(err)
	}
	logger.Infof("start DingTalk calendar ...")

	server := api.NewHttpServer()
	if err := server.Run(":3022"); err != nil {
		logger.Errorf("start server failed, err=%s", err.Error())
	}
	select {}
}
