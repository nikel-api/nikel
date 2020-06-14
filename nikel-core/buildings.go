package main

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func getBuildings(c *gin.Context) {
	limitQuery := parseInt(c.Query("limit"), 0, 100, 10)
	skipQuery := parseInt(c.Query("skip"), 0, math.MaxInt64, 0)
	if skipQuery >= len(buildingsMap) {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       []Building{}},
		)
		return
	}
	upperLimit := skipQuery + limitQuery
	var buildings []Building
	for idx := skipQuery; idx < upperLimit && idx < len(buildingsMap); idx++ {
		buildings = append(buildings, buildingsMap[strconv.Itoa(idx+1)])
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
	if skipQuery >= len(buildingsMap) {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       []Building{}},
		)
		return
	}

	// TODO: Optimize performance later
	var resBuildings []Building
	for idx := 0; idx < len(buildingsMap); idx++ {
		building := buildingsMap[strconv.Itoa(idx+1)]
		if filterIntQuery(c.Query("id"), building.ID, 1, len(buildingsMap)) &&
			filterQuery(c.Query("code"), building.Code) &&
			filterQuery(c.Query("name"), building.Name) &&
			filterQuery(c.Query("campus"), building.Campus) &&
			filterQuery(c.Query("street_number"), building.Address.StreetNumber) &&
			filterQuery(c.Query("street_name"), building.Address.StreetName) &&
			filterQuery(c.Query("city"), building.Address.City) &&
			filterQuery(c.Query("province"), building.Address.Province) &&
			filterQuery(c.Query("country"), building.Address.Country) &&
			filterQuery(c.Query("postal_code"), building.Address.PostalCode) {
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
