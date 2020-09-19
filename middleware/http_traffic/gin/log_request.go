package gin

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/middleware/http_traffic/internal/writer"
	"io/ioutil"
	"net/http/httputil"
)

func format(c *gin.Context)  {
	//log.Format(true, c.Request.Method, c.Request.URL.String(), c.Request.Header, )
}

func (g Gin) LogRequest(write writer.Wrapper) gin.HandlerFunc {
	return func(context *gin.Context) {
		if body, err := ioutil.ReadAll(context.Request.Body); err != nil {
			awesome_error.CheckErr(err)
		} else {
			if len(body) > 0 {
				context.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

				requestDump, err := httputil.DumpRequest(context.Request, true)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(string(requestDump))

				if _, err := write.Write(body); err != nil {
					awesome_error.CheckErr(err)
				}
			}
		}
		context.Next()
	}
}
