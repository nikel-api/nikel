package main

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

func getTextbooks(c *gin.Context) {
	limitQuery := parseInt(c.Query("limit"), 0, 100, 10)
	skipQuery := parseInt(c.Query("skip"), 0, math.MaxInt64, 0)
	if skipQuery >= len(textbooksOrder) {
		c.JSON(http.StatusOK, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       []Textbook{}},
		)
		return
	}
	upperLimit := skipQuery + limitQuery
	var textbooks []Textbook
	for idx := skipQuery; idx < upperLimit && idx < len(textbooksOrder); idx++ {
		textbooks = append(textbooks, textbooksMap[textbooksOrder[idx]])
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response":       textbooks},
	)
}

func textbookHandler(c *gin.Context) {
	p1 := c.Param("p1")
	if p1 == "search" {
		getTextbooksBySearch(c)
	} else {
		getTextbookByID(c)
	}
}

func getTextbookByID(c *gin.Context) {
	id := c.Param("p1")
	if val, ok := textbooksMap[id]; ok {
		c.JSON(http.StatusOK, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       val},
		)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code":    http.StatusNotFound,
			"status_message": "textbook not found",
			"response":       nil},
		)
	}
}

func getTextbooksBySearch(c *gin.Context) {
	limitQuery := parseInt(c.Query("limit"), 0, 100, 10)
	skipQuery := parseInt(c.Query("skip"), 0, math.MaxInt64, 0)
	if skipQuery >= len(textbooksOrder) {
		c.JSON(http.StatusOK, gin.H{
			"status_code":    http.StatusOK,
			"status_message": "success",
			"response":       []Textbook{}},
		)
		return
	}

	// TODO: Optimize performance later
	var resTextbooks []Textbook
	for _, v := range textbooksOrder {
		textbook := textbooksMap[v]
		if filterQuery(c.Query("id"), textbook.ID) &&
			filterQuery(c.Query("isbn"), textbook.Isbn) &&
			filterQuery(c.Query("title"), textbook.Title) &&
			filterIntQuery(c.Query("edition"), textbook.Edition, 0, math.MaxInt64) &&
			filterQuery(c.Query("author"), textbook.Author) &&
			filterValueQuery(c.Query("price"), textbook.Price, 0, math.MaxInt64) {
		coursesOut:
			for _, w := range textbook.Courses {
				if filterQuery(c.Query("course_id"), w.ID) &&
					filterQuery(c.Query("course_code"), w.Code) &&
					filterQuery(c.Query("requirement"), w.Requirement) {
					for _, x := range w.MeetingSections {
						if filterQuery(c.Query("section_code"), x.Code) &&
							filterQueryArr(c.Query("instructors"), x.Instructors) {
							resTextbooks = append(resTextbooks, textbook)
							break coursesOut
						}
					}
				}
			}
		}
	}

	var payload []Textbook
	upperLimit := skipQuery + limitQuery
	for idx := skipQuery; idx < upperLimit && idx < len(resTextbooks); idx++ {
		payload = append(payload, resTextbooks[idx])
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response":       payload},
	)
}
