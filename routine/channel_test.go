package routine

import (
	"github.com/ssst0n3/awesome_libs/log"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	var singleThreadFlag bool
	chanResult := make(chan Result, 10)
	mapResult := make(map[string]Result)
	Run(&singleThreadFlag, chanResult, mapResult)
	chanResult <- Result{
		Key:   "test",
		Value: true,
	}
	time.Sleep(time.Millisecond*100)
	log.Logger.Info(mapResult)
}
