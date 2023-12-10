package vlog

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/go-co-op/gocron"
)

type Tag string

const (
	LogTagDebug   Tag = "DBG"
	LogTagInfo    Tag = "INF"
	LogTagWarn    Tag = "WRN"
	LogTagError   Tag = "ERR"
	LogTagRequest Tag = "REQ"
)

type Logger struct {
	outputs   *[]Output
	scheduler *gocron.Scheduler
}

func New(opts ...func(l *Logger)) *Logger {
	l := &Logger{
		outputs: &[]Output{},
	}
	for _, opt := range opts {
		opt(l)
	}
	l.startCron()
	return l
}

func (l *Logger) Close() {
	for _, output := range *l.outputs {
		err := output.close()
		if err != nil {
			_, _ = fmt.Fprint(os.Stderr, err.Error())
		}
	}
	l.stopCron()
}

func (l *Logger) Debug(msg string, fields ...KeyValue) {
	if os.Getenv("DEBUG") == "" {
		return
	}
	l.print(color.FgWhite, LogTagDebug, msg, fields...)
}

func (l *Logger) Info(msg string, fields ...KeyValue) {
	l.print(color.FgBlue, LogTagInfo, msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...KeyValue) {
	l.print(color.FgYellow, LogTagWarn, msg, fields...)
}

func (l *Logger) Error(err error, fields ...KeyValue) {
	l.print(color.FgRed, LogTagError, err.Error(), fields...)
}

func (l *Logger) Request(msg string, fields ...KeyValue) {
	l.print(color.FgGreen, LogTagRequest, msg, fields...)
}

func (l *Logger) Raw(msg string) {
	for _, output := range *l.outputs {
		output.printRaw(msg)
	}
}

func (l *Logger) print(color color.Attribute, tag Tag, msg string, fields ...KeyValue) {
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

func (l *Logger) startCron() {
	l.scheduler = gocron.NewScheduler(time.Local)
	_, err := l.scheduler.Every(1).Day().At("00:00").Do(func() {
		for _, output := range *l.outputs {
			err := output.close()
			if err != nil {
				_, _ = fmt.Fprint(os.Stderr, err.Error())
				continue
			}

			err = output.open()
			if err != nil {
				_, _ = fmt.Fprint(os.Stderr, err.Error())
				continue
			}
		}
	})
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
		return
	}
	l.scheduler.StartAsync()
}

func (l *Logger) stopCron() {
	l.scheduler.Stop()
	l.scheduler = nil
}
