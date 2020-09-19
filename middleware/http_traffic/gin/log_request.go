package gin

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/middleware/http_traffic/internal/writer"
	"io/ioutil"
)

func (g Gin) LogRequest(write writer.Wrapper) gin.HandlerFunc {
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
