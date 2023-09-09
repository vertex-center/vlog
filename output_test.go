package vlog

import (
	"os"
	"testing"
	"time"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

var line = &Line{
	color: color.FgRed,
	now:   time.Unix(0, 0),
	tag:   LogTagInfo,
	msg:   "message",
	fields: []KeyValue{
		{Key: "key", Value: "value"},
	},
}

func TestOutputStd(t *testing.T) {
	stdout, err := os.CreateTemp("", "out")
	assert.NoError(t, err)

	defer stdout.Close()
	defer os.Remove(stdout.Name())

	stderr, err := os.CreateTemp("", "err")
	assert.NoError(t, err)

	defer stderr.Close()
	defer os.Remove(stderr.Name())

	output := &OutputStd{
		stdout: stdout,
		stderr: stderr,
	}
	output.print(line)

	file, err := os.ReadFile(stdout.Name())
	assert.NoError(t, err)

	assert.Contains(t, string(file), " INF msg=message key=value\n")
}

func TestOutputTextFile(t *testing.T) {
	file, err := os.CreateTemp("", "log")
	assert.NoError(t, err)

	defer file.Close()
	defer os.Remove(file.Name())

	output := &OutputTextFile{}
	output.file = file
	output.print(line)

	content, err := os.ReadFile(file.Name())
	assert.NoError(t, err)

	assert.Contains(t, string(content), " INF msg=message key=value\n")
}

func TestOutputJsonFile(t *testing.T) {
	file, err := os.CreateTemp("", "log")
	assert.NoError(t, err)

	defer file.Close()
	defer os.Remove(file.Name())

	output := &OutputJsonFile{}
	output.file = file
	output.print(line)

	content, err := os.ReadFile(file.Name())
	assert.NoError(t, err)

	expected := "{\"seconds\":0,\"nanoseconds\":0,\"kind\":\"INF\",\"msg\":\"message\",\"key\":\"value\"}\n"
	assert.JSONEq(t, expected, string(content))
}
