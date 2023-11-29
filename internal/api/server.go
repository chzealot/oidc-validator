package api

import (
	"github.com/chzealot/gobase/logger"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

func (s *HttpServer) Run(address string) error {
	logger.Infof("run http server on %s", address)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	return r.Run(address)
}
