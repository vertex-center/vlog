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
	LogFormatColoredText
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

func WithOutputFunc(format LogFormat, fc func(line string)) func(l *Logger) {
	return func(l *Logger) {
		*l.outputs = append(*l.outputs, &OutputFunc{
			format: format,
			fc:     fc,
		})
	}
}

func WithOutputFile(format LogFormat, dir string) func(l *Logger) {
	return func(l *Logger) {
		output := &OutputFile{
			dir:    dir,
			format: format,
		}

		switch format {
		case LogFormatText:
			output.ext = ".txt"
		case LogFormatJson:
			output.ext = ".jsonl"
		default:
			output.ext = ".txt"
		}

		err := output.open()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to open log file: %v\n", err)
			return
		}
		*l.outputs = append(*l.outputs, output)
	}
}

type OutputFile struct {
	dir    string
	ext    string
	format LogFormat
	file   *os.File
}

func (out *OutputFile) print(l *Line) {
	_, _ = fmt.Fprintln(out.file, l.ToFormat(out.format))
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

	_, _ = fmt.Fprintln(file, l.ToColoredText())
}

func (o OutputStd) open() error {
	return nil
}

func (o OutputStd) close() error {
	return nil
}

type OutputFunc struct {
	format LogFormat
	fc     func(line string)
}

func (o *OutputFunc) print(l *Line) {
	o.fc(l.ToFormat(o.format))
}

func (o *OutputFunc) open() error {
	return nil
}

func (o *OutputFunc) close() error {
	return nil
}
