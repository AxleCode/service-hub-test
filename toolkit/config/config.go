package config

import "time"

type KVStore interface {
	IsSet(key string) bool
	Get(key string) interface{}
	AllKeys() []string
	GetBool(key string) bool
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetIntSlice(key string) []int
	GetString(key string) string
	GetStringSlice(key string) []string
	GetStringMap(key string) map[string]interface{}
	GetStringMapStringSlice(key string) map[string][]string
	GetFloat64(key string) float64
	GetUint(key string) uint
	GetUint32(key string) uint32
	GetUint64(key string) uint64
	GetSizeInBytes(key string) uint
}