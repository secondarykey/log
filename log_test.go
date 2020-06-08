package log_test

import (
	own "github.com/secondarykey/log"

	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {

	logger := own.Get()
	if logger == nil {
		t.Errorf("logger is nil")
	}

	l := log.New(os.Stdout, "PREFIX:", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	own.Set(l, own.INFO)

	logger = own.Get()
	logger.Debug("test")
}
