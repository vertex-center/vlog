package vlog

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/fatih/color"
)

type LogFormat int

const (
	LogFormatText = iota
	LogFormatJson
)

type Output interface {
	print(line *Line)
}

func WithOutputStd() func(l *Logger) {
	return func(l *Logger) {
		*l.outputs = append(*l.outputs, OutputStd{})
	}
}

func WithOutputFile(dir string, format LogFormat) func(l *Logger) {
	return func(l *Logger) {
		switch format {
		case LogFormatText:
			output := &OutputTextFile{}
			output.dir = dir
			output.ext = ".txt"
			err := output.Open()
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to open log file: %v\n", err)
				return
			}
			*l.outputs = append(*l.outputs, output)
		case LogFormatJson:
			output := &OutputJsonFile{}
			output.dir = dir
			output.ext = ".jsonl"
			err := output.Open()
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to open log file: %v\n", err)
				return
			}
			*l.outputs = append(*l.outputs, output)
		}
	}
}

type OutputFile struct {
	dir  string
	ext  string
	file *os.File
}

func (out *OutputFile) Open() error {
	err := os.Mkdir(out.dir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	filename := fmt.Sprintf("vertex_logs_%s.%s", time.Now().Format(time.DateOnly), out.ext)

	out.file, err = os.OpenFile(path.Join(out.dir, filename), os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	return err
}

func (out *OutputFile) Close() error {
	return out.file.Close()
}

type OutputJsonFile struct {
	OutputFile
}

func (o *OutputJsonFile) print(l *Line) {
	m := map[string]any{
		"seconds":     l.now.Unix(),
		"nanoseconds": l.now.UnixNano(),
		"kind":        l.tag,
		"msg":         l.msg,
	}
	for _, field := range l.fields {
		m[field.Key] = field.Value
	}

	j, err := json.Marshal(m)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to marshal json: %v\n", err)
		return
	}

	_, err = fmt.Fprintln(o.file, string(j))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to write to log file: %v\n", err)
	}
}

type OutputTextFile struct {
	OutputFile
}

func (o *OutputTextFile) print(l *Line) {
	msg := fmt.Sprintf("%s %s %s", l.now.Format(time.DateTime), l.tag, l.msg)
	for _, field := range l.fields {
		msg += fmt.Sprintf(" %s=%s", field.Key, field.Value)
	}
	_, err := fmt.Fprintln(o.file, msg)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to write to log file: %v\n", err)
	}
}

type OutputStd struct{}

func (o OutputStd) print(l *Line) {
	var file *os.File
	if l.tag == LogTagError {
		file = os.Stderr
	} else {
		file = os.Stdout
	}

	msg := fmt.Sprintf("%s %s",
		color.New(color.FgHiWhite).Sprintf(l.now.Format(time.DateTime)),
		color.New(l.color).Sprint(l.tag),
	)
	msg += color.New(l.color).Sprintf(" msg=")
	msg += fmt.Sprint(l.msg)
	for _, field := range l.fields {
		msg += color.New(l.color).Sprintf(" %s=", field.Key)
		msg += fmt.Sprint(field.Key)
	}

	_, err := fmt.Fprintln(file, msg)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to write to log file: %v\n", err)
	}
}
