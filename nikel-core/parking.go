package main

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func getParking(c *gin.Context) {
	limitQuery := parseInt(c.Query("limit"), 0, 100, 10)
	skipQuery := parseInt(c.Query("skip"), 0, math.MaxInt64, 0)
	if skipQuery >= len(parkingOrder) {
		c.JSON(http.StatusOK, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       []Parking{}},
		)
		return
	}
	upperLimit := skipQuery + limitQuery
	var parking []Parking
	for idx := skipQuery; idx < upperLimit && idx < len(parkingOrder); idx++ {
		parking = append(parking, parkingMap[parkingOrder[idx]])
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response":       parking},
	)
}

func parkingHandler(c *gin.Context) {
	p1 := c.Param("p1")
	if p1 == "search" {
		getParkingBySearch(c)
	} else {
		getParkingByID(c)
	}
}

func getParkingByID(c *gin.Context) {
	id := c.Param("p1")
	if val, ok := parkingMap[id]; ok {
		c.JSON(http.StatusOK, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       val},
		)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code":    http.StatusNotFound,
			"status_message": "parking not found",
			"response":       nil},
		)
	}
}

func getParkingBySearch(c *gin.Context) {
	limitQuery := parseInt(c.Query("limit"), 0, 100, 10)
	skipQuery := parseInt(c.Query("skip"), 0, math.MaxInt64, 0)
	if skipQuery >= len(parkingOrder) {
		c.JSON(http.StatusOK, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       []Parking{}},
		)
		return
	}

	// TODO: Optimize performance later
	var resParking []Parking
	for _, v := range parkingOrder {
		parking := parkingMap[v]
		if filterQuery(c.Query("id"), parking.ID) &&
			filterQuery(c.Query("name"), parking.Name) &&
			filterQuery(c.Query("alias"), parking.Alias) &&
			filterQuery(c.Query("building_id"), parking.BuildingID) &&
			filterQuery(c.Query("description"), parking.Description) &&
			filterQuery(c.Query("campus"), parking.Campus) &&
			filterQuery(c.Query("address"), parking.Address) &&
			filterValueQuery(c.Query("latitude"), parking.Coordinates.Latitude, -90, 90) &&
			filterValueQuery(c.Query("longitude"), parking.Coordinates.Longitude, -180, 180) {
			resParking = append(resParking, parking)
		}
	}

	var payload []Parking
	upperLimit := skipQuery + limitQuery
	for idx := skipQuery; idx < upperLimit && idx < len(resParking); idx++ {
		payload = append(payload, resParking[idx])
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response":       payload},
	)
}
