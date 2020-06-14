package log_test

import (
	own "github.com/secondarykey/log"
	"log"
	"testing"
	"time"
)

func TestWriter(t *testing.T) {

	w, err := own.NewRollingFileWriter("./", own.Day)
	if err != nil {
		t.Errorf("NewRollingFileWriter error")
	}
	log.SetOutput(w)

	log.Println("test")
}

func TestRolling(t *testing.T) {
	w, err := own.NewRollingFileWriter("./", own.Second)
	if err != nil {
		t.Errorf("NewRollingFileWriter error")
	}
	log.SetOutput(w)

	limit := 5 * time.Second
	begin := time.Now()
	for now := range time.Tick(10 * time.Millisecond) {
		log.Println("write time is", now)
		if now.Sub(begin) >= limit {
			break
		}
	}
}
