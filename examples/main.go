package main

import (
	"errors"

	"github.com/vertex-center/vlog"
)

func main() {
	log := vlog.New(
		vlog.WithOutputStd(),
		vlog.WithOutputFile(vlog.LogFormatText, "logs"),
		vlog.WithOutputFile(vlog.LogFormatJson, "logs"),
		vlog.WithOutputFunc(vlog.LogFormatJson, func(line string) {
			println(line)
		}),
	)
	defer log.Close()

	log.Debug("message", vlog.String("name", "abc"))
	log.Info("message", vlog.String("name", "abc"))
	log.Warn("message", vlog.String("name", "abc"))
	log.Error(errors.New("message"), vlog.String("name", "abc"))
	log.Request("message", vlog.String("name", "abc"))
}
