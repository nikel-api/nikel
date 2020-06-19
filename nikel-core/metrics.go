package main

import (
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"runtime"
)

// getMetrics returns runtime metrics for app health monitoring
func getMetrics(c *gin.Context) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	sendSuccess(c, gin.H{
		"memory":           memStats.Alloc,
		"memory_humanized": humanize.Bytes(memStats.Alloc),
		"sys":              memStats.Sys,
		"sys_humanized":    humanize.Bytes(memStats.Sys),
		"routines":         runtime.NumGoroutine(),
	})
}
