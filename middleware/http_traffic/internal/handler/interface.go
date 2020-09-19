package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/middleware/http_traffic/internal/writer"
)

type Interface interface {
	LogRequest(write writer.Wrapper) gin.HandlerFunc
	LogResponse(write writer.Wrapper) gin.HandlerFunc
}
