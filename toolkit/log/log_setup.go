package log

import (
	"fmt"
	stdlog "log"
	kitconfig "gitlab.com/wit-id/service-hub-test/toolkit/config"
)

func NewFromConfig(cfg kitconfig.KVStore, path string) (*Logger, error) {
	_ = cfg.GetString(fmt.Sprintf("%s.level", path))
	return &Logger{StdLog: stdlog.Default()}, nil
}