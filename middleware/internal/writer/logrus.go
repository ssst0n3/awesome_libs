package writer

import "github.com/sirupsen/logrus"

type Logrus struct {
	Logger *logrus.Logger
}

func (l Logrus) Write(content []byte) (int, error) {
	l.Logger.Log(l.Logger.Level, string(content))
	return 0, nil
}
