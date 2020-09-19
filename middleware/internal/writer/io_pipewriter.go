package writer

import "io"

type IoPipeWriter struct {
	Writer *io.PipeWriter
}

func (l IoPipeWriter) Write(content []byte) (int, error) {
	return l.Writer.Write(content)
}
