package vlog

import (
	"fmt"
	"strconv"
)

type KeyValue struct {
	Key   string
	Value string
}

func String(key string, value string) KeyValue {
	return KeyValue{Key: key, Value: value}
}

func Int(key string, value int) KeyValue {
	return KeyValue{Key: key, Value: strconv.Itoa(value)}
}

func Int8(key string, value int8) KeyValue {
	return KeyValue{Key: key, Value: strconv.Itoa(int(value))}
}

func Int16(key string, value int16) KeyValue {
	return KeyValue{Key: key, Value: strconv.Itoa(int(value))}
}

func Int32(key string, value int32) KeyValue {
	return KeyValue{Key: key, Value: strconv.Itoa(int(value))}
}

func Int64(key string, value int64) KeyValue {
	return KeyValue{Key: key, Value: strconv.FormatInt(value, 10)}
}

func Uint(key string, value uint) KeyValue {
	return KeyValue{Key: key, Value: strconv.FormatUint(uint64(value), 10)}
}

func Uint8(key string, value uint8) KeyValue {
	return KeyValue{Key: key, Value: strconv.FormatUint(uint64(value), 10)}
}

func Uint16(key string, value uint16) KeyValue {
	return KeyValue{Key: key, Value: strconv.FormatUint(uint64(value), 10)}
}

func Uint32(key string, value uint32) KeyValue {
	return KeyValue{Key: key, Value: strconv.FormatUint(uint64(value), 10)}
}

func Uint64(key string, value uint64) KeyValue {
	return KeyValue{Key: key, Value: strconv.FormatUint(value, 10)}
}

func Float32(key string, value float32) KeyValue {
	return KeyValue{Key: key, Value: strconv.FormatFloat(float64(value), 'f', -1, 32)}
}

func Float64(key string, value float64) KeyValue {
	return KeyValue{Key: key, Value: strconv.FormatFloat(value, 'f', -1, 64)}
}

func Bool(key string, value bool) KeyValue {
	return KeyValue{Key: key, Value: strconv.FormatBool(value)}
}

func Any(key string, value any) KeyValue {
	return KeyValue{Key: key, Value: fmt.Sprintf("%+v", value)}
}
