package metrics

import (
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/nikel-api/nikel/nikel-core/response"
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

// GetMetrics returns runtime metrics for app health monitoring
func GetMetrics(c *gin.Context) {
	var memStats runtime.MemStats

	// get runtime memory stats
	runtime.ReadMemStats(&memStats)

	// send successful response
	// humanize some values because humans are bad at math
	response.SendSuccess(c, Metrics{
		Memory:          memStats.Alloc,
		MemoryHumanized: humanize.Bytes(memStats.Alloc),
		Sys:             memStats.Sys,
		SysHumanized:    humanize.Bytes(memStats.Sys),
		Pause:           memStats.PauseTotalNs,
		PauseHumanized:  time.Duration(memStats.PauseTotalNs).String(),
		Goroutines:      runtime.NumGoroutine(),
		LogicalCores:    runtime.NumCPU(),
		StartTime:       humanize.Time(startTime),
	})
}
