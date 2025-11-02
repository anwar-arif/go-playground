package log_aggregator

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Log struct {
	Service string
	Level   string
	Message string
	Time    time.Time
}

type Aggregator struct {
	mu                 sync.RWMutex
	logs               []Log
	LogsCountByService map[string]int64
	FilteredLogs       map[string][]Log
}

func NewAggregator() *Aggregator {
	return &Aggregator{
		logs:               make([]Log, 0),
		LogsCountByService: make(map[string]int64),
		FilteredLogs:       make(map[string][]Log),
	}
}

func Code(service, level string) string {
	return fmt.Sprintf("%s|%s", service, level)
}

func (a *Aggregator) AddLog(log Log) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if log.Level == "" || log.Service == "" {
		return errors.New("service name or level name cannot be empty")

	}
	a.logs = append(a.logs, log)
	a.LogsCountByService[log.Service]++
	a.FilteredLogs[Code(log.Service, log.Level)] = append(a.FilteredLogs[Code(log.Service, log.Level)], log)
	return nil
}

func (a *Aggregator) Filter(service, level string) []Log {
	code := Code(service, level)
	return a.FilteredLogs[code]
}

func (a *Aggregator) CountByService() map[string]int64 {
	return a.LogsCountByService
}
