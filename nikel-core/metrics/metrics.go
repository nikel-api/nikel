package metrics

import (
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/nikel-api/nikel/nikel-core/config"
	"github.com/nikel-api/nikel/nikel-core/response"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

// Metrics forces deserialization of metrics
// to be ordered a certain way
type Metrics struct {
	Memory          uint64 `json:"memory"`
	MemoryHumanized string `json:"memory_humanized"`
	Sys             uint64 `json:"sys"`
	SysHumanized    string `json:"sys_humanized"`
	Cache           uint64 `json:"cache"`
	CacheHumanized  string `json:"cache_humanized"`
	Pause           uint64 `json:"pause"`
	PauseHumanized  string `json:"pause_humanized"`
	Goroutines      int    `json:"goroutines"`
	LogicalCores    int    `json:"logical_cores"`
	StartTime       string `json:"start_time"`
}

// startTime to track boot time
var startTime time.Time

func init() {
	// initialize startTime
	startTime = time.Now()
}

func dirSize(path string) (size uint64, err error) {
	err = filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// info.Size should return a non-negative int
			size += uint64(info.Size())
		}
		return err
	})
	return size, err
}

// GetMetrics returns runtime metrics for app health monitoring
func GetMetrics(c *gin.Context) {
	var memStats runtime.MemStats

	// get runtime memory stats
	runtime.ReadMemStats(&memStats)

	// calculate cache size
	var size uint64
	if config.CacheFlag.Load() {
		var err error
		size, err = dirSize(config.CachePath)

		if err != nil {
			size = 0
		}
	} else {
		size = 0
	}

	// send successful response
	// humanize some values because humans are bad at math
	response.SendSuccess(c, Metrics{
		Memory:          memStats.Alloc,
		MemoryHumanized: humanize.Bytes(memStats.Alloc),
		Sys:             memStats.Sys,
		SysHumanized:    humanize.Bytes(memStats.Sys),
		Cache:           size,
		CacheHumanized:  humanize.Bytes(size),
		Pause:           memStats.PauseTotalNs,
		PauseHumanized:  time.Duration(memStats.PauseTotalNs).String(),
		Goroutines:      runtime.NumGoroutine(),
		LogicalCores:    runtime.NumCPU(),
		StartTime:       humanize.Time(startTime),
	})
}
