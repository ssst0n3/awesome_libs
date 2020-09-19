package writer

import "io"

type IoWriter struct {
	W io.Writer
}

func (w IoWriter) Write(content []byte) (int, error) {
	return w.W.Write(content)
}
