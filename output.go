package vlog

import (
	"fmt"
	"os"
	"path"
	"time"
)

type LogFormat int

const (
	LogFormatText = iota
	LogFormatJson
)

type Output interface {
	print(line *Line)
	open() error
	close() error
}

func WithOutputStd() func(l *Logger) {
	return func(l *Logger) {
		*l.outputs = append(*l.outputs, OutputStd{
			stdout: os.Stdout,
			stderr: os.Stderr,
		})
	}
}

func WithOutputFile(format LogFormat, dir string) func(l *Logger) {
	return func(l *Logger) {
		switch format {
		case LogFormatText:
			output := &OutputTextFile{}
			output.dir = dir
			output.ext = ".txt"
			err := output.open()
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "failed to open log file: %v\n", err)
				return
			}
			*l.outputs = append(*l.outputs, output)
		case LogFormatJson:
			output := &OutputJsonFile{}
			output.dir = dir
			output.ext = ".jsonl"
			err := output.open()
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

func (out *OutputFile) open() error {
	err := os.MkdirAll(out.dir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	filename := fmt.Sprintf("vertex_logs_%s%s", time.Now().Format(time.DateOnly), out.ext)

	out.file, err = os.OpenFile(path.Join(out.dir, filename), os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	return err
}

func (out *OutputFile) close() error {
	return out.file.Close()
}

type OutputJsonFile struct {
	OutputFile
}

func (o *OutputJsonFile) print(l *Line) {
	j, err := l.ToJson()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v", err)
		return
	}

	_, err = fmt.Fprintln(o.file, j)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to write to log file: %v\n", err)
	}
}

type OutputTextFile struct {
	OutputFile
}

func (o *OutputTextFile) print(l *Line) {
	_, err := fmt.Fprintln(o.file, l.ToText())
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to write to log file: %v\n", err)
	}
}

type OutputStd struct {
	stdout *os.File
	stderr *os.File
}

func (o OutputStd) print(l *Line) {
	var file *os.File
	if l.tag == LogTagError {
		file = o.stderr
	} else {
		file = o.stdout
	}

	_, err := fmt.Fprintln(file, l.ToColoredText())
	if err != nil {
		_, _ = fmt.Fprintf(file, "failed to write to log file: %v\n", err)
	}
}

func (o OutputStd) open() error {
	return nil
}

func (o OutputStd) close() error {
	return nil
}
