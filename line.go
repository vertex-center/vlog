package vlog

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fatih/color"
)

type Line struct {
	color  color.Attribute
	now    time.Time
	tag    Tag
	msg    string
	fields []KeyValue
}

func (l *Line) ToFormat(format LogFormat) string {
	switch format {
	case LogFormatText:
		return l.ToText()
	case LogFormatJson:
		j, _ := l.ToJson()
		return j
	case LogFormatColoredText:
		return l.ToColoredText()
	default:
		return l.ToText()
	}
}

func (l *Line) ToColoredText() string {
	msg := fmt.Sprintf("%s %s",
		color.New(color.FgHiWhite).Sprintf(l.now.Format(time.DateTime)),
		color.New(l.color).Sprint(l.tag),
	)
	msg += color.New(l.color).Sprintf(" msg=")
	msg += fmt.Sprint(l.msg)
	for _, field := range l.fields {
		msg += color.New(l.color).Sprintf(" %s=", field.Key)
		msg += fmt.Sprint(field.Value)
	}
	return msg
}

func (l *Line) ToText() string {
	msg := fmt.Sprintf("%s %s msg=%s", l.now.Format(time.DateTime), l.tag, l.msg)
	for _, field := range l.fields {
		msg += fmt.Sprintf(" %s=%s", field.Key, field.Value)
	}
	return msg
}

func (l *Line) ToJson() (string, error) {
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
		return "", fmt.Errorf("failed to marshal json: %w", err)
	}
	return string(j), nil
}
