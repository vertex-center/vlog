package vlog

import (
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
