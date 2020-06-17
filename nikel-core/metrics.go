package main

import (
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"runtime"
)

func getMetrics(c *gin.Context) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	sendSuccess(c, gin.H{
		"memory":           memStats.TotalAlloc,
		"memory_humanized": humanize.Bytes(memStats.TotalAlloc),
		"sys":              memStats.Sys,
		"sys_humanized":    humanize.Bytes(memStats.Sys),
		"routines":         runtime.NumGoroutine(),
	})
}
