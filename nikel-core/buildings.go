package main

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func getBuildings(c *gin.Context) {
	limitQuery := parseInt(c.Query("limit"), 0, 100, 10)
	skipQuery := parseInt(c.Query("skip"), 0, math.MaxInt64, 0)
	if skipQuery >= len(buildingsOrder) {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       []Building{}},
		)
		return
	}
	upperLimit := skipQuery + limitQuery
	var buildings []Building
	for idx := skipQuery; idx < upperLimit && idx < len(buildingsOrder); idx++ {
		buildings = append(buildings, buildingsMap[buildingsOrder[idx]])
	}
	c.JSON(http.StatusNotFound, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response":       buildings},
	)
}

func buildingHandler(c *gin.Context) {
	p1 := c.Param("p1")
	if p1 == "search" {
		getBuildingBySearch(c)
	} else {
		getBuildingByID(c)
	}
}

func getBuildingByID(c *gin.Context) {
	id := c.Param("p1")
	if val, ok := buildingsMap[id]; ok {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       val},
		)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code":    http.StatusNotFound,
			"status_message": "building not found",
			"response":       nil},
		)
	}
}

func getBuildingBySearch(c *gin.Context) {
	limitQuery := parseInt(c.Query("limit"), 0, 100, 10)
	skipQuery := parseInt(c.Query("skip"), 0, math.MaxInt64, 0)
	if skipQuery >= len(buildingsOrder) {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       []Building{}},
		)
		return
	}

	// TODO: Optimize performance later
	var resBuildings []Building
	for _, v := range buildingsOrder {
		building := buildingsMap[v]
		if filterQuery(c.Query("id"), building.ID) &&
			filterQuery(c.Query("code"), building.Code) &&
			filterQuery(c.Query("tags"), building.Tags) &&
			filterQuery(c.Query("name"), building.Name) &&
			filterQuery(c.Query("short_name"), building.ShortName) &&
			filterQuery(c.Query("street"), building.Address.Street) &&
			filterQuery(c.Query("city"), building.Address.City) &&
			filterQuery(c.Query("province"), building.Address.Province) &&
			filterQuery(c.Query("country"), building.Address.Country) &&
			filterQuery(c.Query("postal"), building.Address.Postal) &&
			filterValueQuery(c.Query("latitude"), building.Coordinates.Latitude, -90, 90) &&
			filterValueQuery(c.Query("longitude"), building.Coordinates.Longitude, -180, 180) {
			resBuildings = append(resBuildings, building)
		}
	}

	var buildings []Building
	upperLimit := skipQuery + limitQuery
	for idx := skipQuery; idx < upperLimit && idx < len(resBuildings); idx++ {
		buildings = append(buildings, resBuildings[idx])
	}
	c.JSON(http.StatusNotFound, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response":       buildings},
	)
}
