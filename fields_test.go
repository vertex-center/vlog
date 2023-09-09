package vlog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyValues(t *testing.T) {
	kvs := []struct {
		value    KeyValue
		expected string
	}{
		{String("key", "value"), "value"},

		{Int("key", -4), "-4"},
		{Int8("key", -4), "-4"},
		{Int16("key", -4), "-4"},
		{Int32("key", -4), "-4"},
		{Int64("key", -4), "-4"},

		{Uint("key", 3), "3"},
		{Uint8("key", 3), "3"},
		{Uint16("key", 3), "3"},
		{Uint32("key", 3), "3"},
		{Uint64("key", 3), "3"},

		{Float32("key", 3.14), "3.14"},
		{Float64("key", 3.14), "3.14"},

		{Bool("key", true), "true"},
		{Bool("key", false), "false"},

		{Any("key", "value"), "value"},
		{Any("key", 3), "3"},
		{Any("key", map[string]any{"a": "b"}), "map[a:b]"},
	}

	for _, kv := range kvs {
		assert.Equal(t, "key", kv.value.Key)
		assert.Equal(t, kv.expected, kv.value.Value)
	}
}
