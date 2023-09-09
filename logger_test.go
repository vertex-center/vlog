package vlog

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogging(t *testing.T) {
	stdout, err := os.CreateTemp("", "stdout")
	assert.NoError(t, err)

	defer stdout.Close()
	defer os.Remove(stdout.Name())

	log := New()
	log.outputs = &[]Output{
		OutputStd{
			stdout: stdout,
		},
	}
	defer log.Close()

	log.Info("info message")

	stdoutContent, err := os.ReadFile(stdout.Name())
	assert.NoError(t, err)

	assert.Contains(t, string(stdoutContent), " INF msg=info message\n")
}
