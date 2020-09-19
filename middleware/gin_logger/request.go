package gin_logger

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/awesome_libs/middleware/internal/writer"
	"io/ioutil"
)

// Instances a Logger middleware that will write the logs to gin.DefaultWriter
// By default gin.DefaultWriter = os.Stdout
func NewRequest() gin.HandlerFunc {
	return LogRequest(gin.DefaultWriter)
}

func NewLogrusRequest() gin.HandlerFunc {
	return LogRequest(log.Logger.Writer())
}

func LogRequest(write writer.Wrapper) gin.HandlerFunc {
	return func(context *gin.Context) {
		if body, err := ioutil.ReadAll(context.Request.Body); err != nil {
			awesome_error.CheckErr(err)
		} else {
			if len(body) > 0 {
				log.Logger.Info(string(body))
				context.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
			}
		}
		context.Next()
	}
}
