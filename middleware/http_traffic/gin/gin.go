package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/middleware/http_traffic/log_request"
	"github.com/ssst0n3/awesome_libs/middleware/http_traffic/log_response"
)

type Gin struct {
}

var G Gin

func NewGinDefaultWriter() gin.HandlerFunc {
	return log_request.NewGinDefaultWriter(G)
}

func NewLogrus(logger *logrus.Logger) gin.HandlerFunc {
	return log_response.NewLogrus(G, logger)
}
