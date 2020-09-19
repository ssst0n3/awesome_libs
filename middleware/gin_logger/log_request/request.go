package log_request

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/middleware/internal/writer"
	"io/ioutil"
)

// Instances a Logger middleware that will write the logs to gin.DefaultWriter
// By default gin.DefaultWriter = os.Stdout
func NewGinDefaultWriter() gin.HandlerFunc {
	return LogRequest(gin.DefaultWriter)
}

func NewLogrus(logger *logrus.Logger) gin.HandlerFunc {
	return LogRequest(writer.Logrus{
		Logger: logger,
	})
}

func LogRequest(write writer.Wrapper) gin.HandlerFunc {
	return func(context *gin.Context) {
		if body, err := ioutil.ReadAll(context.Request.Body); err != nil {
			awesome_error.CheckErr(err)
		} else {
			if len(body) > 0 {
				context.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
				if _, err := write.Write(body); err != nil {
					awesome_error.CheckErr(err)
				}
			}
		}
		context.Next()
	}
}
