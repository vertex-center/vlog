package vlog

import (
	"time"

	"github.com/fatih/color"
)

type Tag string

const (
	LogTagInfo    Tag = "INF"
	LogTagWarn    Tag = "WRN"
	LogTagError   Tag = "ERR"
	LogTagRequest Tag = "REQ"
)

type Logger struct {
	outputs *[]Output
}

func New(opts ...func(l *Logger)) Logger {
	l := Logger{
		outputs: &[]Output{},
	}
	for _, opt := range opts {
		opt(&l)
	}
	return l
}

func (l Logger) Info(msg string, fields ...KeyValue) {
	l.print(color.FgBlue, LogTagInfo, msg, fields...)
}

func (l Logger) Warn(msg string, fields ...KeyValue) {
	l.print(color.FgYellow, LogTagWarn, msg, fields...)
}

func (l Logger) Error(err error, fields ...KeyValue) {
	l.print(color.FgRed, LogTagError, err.Error(), fields...)
}

func (l Logger) Request(msg string, fields ...KeyValue) {
	l.print(color.FgGreen, LogTagRequest, msg, fields...)
}

func (l Logger) print(color color.Attribute, tag Tag, msg string, fields ...KeyValue) {
	for _, output := range *l.outputs {
		output.print(&Line{
			color:  color,
			now:    time.Now(),
			tag:    tag,
			msg:    msg,
			fields: fields,
		})
	}
}
