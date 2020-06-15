package main

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func getAccessibility(c *gin.Context) {
	limitQuery := parseInt(c.Query("limit"), 0, 100, 10)
	skipQuery := parseInt(c.Query("skip"), 0, math.MaxInt64, 0)
	if skipQuery >= len(accessibilityOrder) {
		c.JSON(http.StatusOK, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       []Accessibility{}},
		)
		return
	}
	upperLimit := skipQuery + limitQuery
	var accessibility []Accessibility
	for idx := skipQuery; idx < upperLimit && idx < len(accessibilityOrder); idx++ {
		accessibility = append(accessibility, accessibilityMap[accessibilityOrder[idx]])
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response":       accessibility},
	)
}

func accessibilityHandler(c *gin.Context) {
	p1 := c.Param("p1")
	if p1 == "search" {
		getAccessibilityBySearch(c)
	} else {
		getAccessibilityByID(c)
	}
}

func getAccessibilityByID(c *gin.Context) {
	id := c.Param("p1")
	if val, ok := accessibilityMap[id]; ok {
		c.JSON(http.StatusOK, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       val},
		)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code":    http.StatusNotFound,
			"status_message": "accessibility not found",
			"response":       nil},
		)
	}
}

func getAccessibilityBySearch(c *gin.Context) {
	limitQuery := parseInt(c.Query("limit"), 0, 100, 10)
	skipQuery := parseInt(c.Query("skip"), 0, math.MaxInt64, 0)
	if skipQuery >= len(accessibilityOrder) {
		c.JSON(http.StatusOK, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       []Accessibility{}},
		)
		return
	}

	// TODO: Optimize performance later
	var resAccessibility []Accessibility
	for _, v := range accessibilityOrder {
		accessbility := accessibilityMap[v]
		if filterQuery(c.Query("id"), accessbility.ID) &&
			filterQuery(c.Query("name"), accessbility.Name) &&
			filterQuery(c.Query("description"), accessbility.Description) &&
			filterQuery(c.Query("building_id"), accessbility.BuildingID) &&
			filterQuery(c.Query("campus"), accessbility.Campus) &&
			filterValueQuery(c.Query("latitude"), accessbility.Coordinates.Latitude, -90, 90) &&
			filterValueQuery(c.Query("longitude"), accessbility.Coordinates.Longitude, -180, 180) &&
			filterQueryArr(c.Query("attributes"), accessbility.Attributes) {
			resAccessibility = append(resAccessibility, accessbility)
		}
	}

	var payload []Accessibility
	upperLimit := skipQuery + limitQuery
	for idx := skipQuery; idx < upperLimit && idx < len(resAccessibility); idx++ {
		payload = append(payload, resAccessibility[idx])
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response":       payload},
	)
}
