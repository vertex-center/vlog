package main

import (
	"errors"

	"github.com/vertex-center/vlog"
)

func main() {
	log := vlog.New(
		vlog.WithOutputStd(),
		vlog.WithOutputFile("logs", vlog.LogFormatText),
		vlog.WithOutputFile("logs", vlog.LogFormatJson),
	)
	defer log.Close()

	log.Debug("message", vlog.String("name", "abc"))
	log.Info("message", vlog.String("name", "abc"))
	log.Warn("message", vlog.String("name", "abc"))
	log.Error(errors.New("message"), vlog.String("name", "abc"))
	log.Request("message", vlog.String("name", "abc"))
}
