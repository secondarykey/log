package log

import (
	"os"
	"path/filepath"
	"time"

	"golang.org/x/xerrors"
)

type RollingFileWriter struct {
	prefix string
	path   string

	format string

	targetValue string
	target      *os.File
}

type Interval int

const (
	Second Interval = iota
	Minite
	Hour
	Day
	Month
	Year
)

func (i Interval) getFormat() string {
	f := ""
	switch i {
	case Second:
		f = "20060102150405"
	case Minite:
		f = "200601021504"
	case Hour:
		f = "2006010215"
	case Day:
		f = "20060102"
	case Month:
		f = "200601"
	case Year:
		f = "2006"
	}
	return f
}

func NewRollingFileWriter(path string, prefix string, i Interval) (*RollingFileWriter, error) {

	w := RollingFileWriter{}

	w.format = i.getFormat()

	w.path = path
	w.prefix = prefix

	w.targetValue = ""

	return &w, nil
}

func (w *RollingFileWriter) Write(p []byte) (int, error) {

	f := time.Now().Format(w.format)
	if f != w.targetValue {
		w.targetValue = f
		err := w.setTarget()
		if err != nil {
			return -1, xerrors.Errorf("setTarget error : %w", err)
		}
	}

	return w.target.Write(p)
}

func (w *RollingFileWriter) Close() error {
	if w.target != nil {
		return w.target.Close()
	}
	return nil
}

func (w *RollingFileWriter) setTarget() error {

	w.Close()

	path := filepath.Join(w.path, w.prefix+"_"+w.targetValue+".log")

	var err error
	w.target, err = os.Create(path)
	if err != nil {
		return xerrors.Errorf("open error: %w", err)
	}
	return nil
}
