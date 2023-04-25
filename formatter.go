package logrus_vertical_formatter

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

type (
	Option interface {
		Configure(*VerticalFormatter)
	}

	optionFunc func(*VerticalFormatter)

	VerticalFormatter struct {
		PadRight func(string) string
	}
)

func (of optionFunc) Configure(f *VerticalFormatter) {
	of(f)
}

func (f VerticalFormatter) Format(e *logrus.Entry) ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	buf.WriteString(f.PadRight("ts") + e.Time.String())
	buf.WriteByte('\n')

	for k, v := range e.Data {

		buf.WriteString(f.PadRight(k) + fmt.Sprintf("%v", v))
		buf.WriteByte('\n')

	}

	buf.WriteString(f.PadRight("message") + e.Message)
	buf.WriteString("\n" + strings.Repeat("_", 100) + "\n")

	return buf.Bytes(), nil
}

func New(opts ...Option) logrus.Formatter {
	f := &VerticalFormatter{
		PadRight: func(s string) string {
			return padRight(s, " ", 40) // defaults
		},
	}

	for _, opt := range opts {
		opt.Configure(f)
	}

	return f
}

func padRight(str string, fill string, length int) string {
	dt := length - len(str)
	if dt <= 0 {
		return str
	}
	return str + strings.Repeat(fill, dt)
}

func WithPadRight(fill string, length int) Option {
	return optionFunc(func(formatter *VerticalFormatter) {
		formatter.PadRight = func(left string) string {
			return padRight(left, fill, length)
		}
	})
}
