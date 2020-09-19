package log_response

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/middleware/http_traffic/internal/handler"
	"github.com/ssst0n3/awesome_libs/middleware/http_traffic/internal/writer"
)

func NewGinDefaultWriter(i handler.Interface) gin.HandlerFunc {
	return i.LogResponse(gin.DefaultWriter)
}

func NewLogrus(i handler.Interface, logger *logrus.Logger) gin.HandlerFunc {
	return i.LogResponse(writer.Logrus{
		Logger: logger,
	})
}
