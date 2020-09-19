package log_response

import (
	"bufio"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
)

type bufferedWriter struct {
	gin.ResponseWriter
	out    *bufio.Writer
	Buffer bytes.Buffer
}

func (g *bufferedWriter) Write(data []byte) (int, error) {
	g.Buffer.Write(data)
	return g.out.Write(data)
}

func LogResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := bufio.NewWriter(c.Writer)
		buff := bytes.Buffer{}
		newWriter := &bufferedWriter{c.Writer, w, buff}

		c.Writer = newWriter

		// You have to manually flush the buffer at the end
		defer func() {
			log.Logger.Info(newWriter.Buffer.String())
			awesome_error.CheckWarning(w.Flush())
		}()

		// Run the execution chain
		c.Next()
	}
}
