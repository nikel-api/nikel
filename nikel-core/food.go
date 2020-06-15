package main

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func getFood(c *gin.Context) {
	limitQuery := parseInt(c.Query("limit"), 0, 100, 10)
	skipQuery := parseInt(c.Query("skip"), 0, math.MaxInt64, 0)
	if skipQuery >= len(foodOrder) {
		c.JSON(http.StatusOK, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       []Food{}},
		)
		return
	}
	upperLimit := skipQuery + limitQuery
	var food []Food
	for idx := skipQuery; idx < upperLimit && idx < len(foodOrder); idx++ {
		food = append(food, foodMap[foodOrder[idx]])
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response":       food},
	)
}

func foodHandler(c *gin.Context) {
	p1 := c.Param("p1")
	if p1 == "search" {
		getFoodBySearch(c)
	} else {
		getFoodByID(c)
	}
}

func getFoodByID(c *gin.Context) {
	id := c.Param("p1")
	if val, ok := foodMap[id]; ok {
		c.JSON(http.StatusOK, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       val},
		)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code":    http.StatusNotFound,
			"status_message": "food not found",
			"response":       nil},
		)
	}
}

func getFoodBySearch(c *gin.Context) {
	limitQuery := parseInt(c.Query("limit"), 0, 100, 10)
	skipQuery := parseInt(c.Query("skip"), 0, math.MaxInt64, 0)
	if skipQuery >= len(foodOrder) {
		c.JSON(http.StatusOK, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       []Food{}},
		)
		return
	}

	// TODO: Optimize performance later
	var resFood []Food
	for _, v := range foodOrder {
		food := foodMap[v]
		if filterQuery(c.Query("id"), food.ID) &&
			filterQuery(c.Query("name"), food.Name) &&
			filterQuery(c.Query("description"), food.Description) &&
			filterQuery(c.Query("tags"), food.Tags) &&
			filterQuery(c.Query("campus"), food.Campus) &&
			filterQuery(c.Query("address"), food.Address) &&
			filterValueQuery(c.Query("latitude"), food.Coordinates.Latitude, -90, 90) &&
			filterValueQuery(c.Query("longitude"), food.Coordinates.Longitude, -180, 180) &&
			filterBoolQuery(c.Query("sunday"), food.Hours.Sunday.Closed, true) &&
			filterBoolQuery(c.Query("monday"), food.Hours.Monday.Closed, true) &&
			filterBoolQuery(c.Query("tuesday"), food.Hours.Tuesday.Closed, true) &&
			filterBoolQuery(c.Query("wednesday"), food.Hours.Wednesday.Closed, true) &&
			filterBoolQuery(c.Query("thursday"), food.Hours.Thursday.Closed, true) &&
			filterBoolQuery(c.Query("friday"), food.Hours.Friday.Closed, true) &&
			filterBoolQuery(c.Query("saturday"), food.Hours.Saturday.Closed, true) &&
			filterQueryArr(c.Query("attributes"), food.Attributes) {
			resFood = append(resFood, food)
		}
	}

	var foods []Food
	upperLimit := skipQuery + limitQuery
	for idx := skipQuery; idx < upperLimit && idx < len(resFood); idx++ {
		foods = append(foods, resFood[idx])
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response":       foods},
	)
}
