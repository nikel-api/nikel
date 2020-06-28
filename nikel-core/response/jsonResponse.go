package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SendSuccess sends successful JSON payload
func SendSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response":       data},
	)
}

// SendEmptySuccess sends successful empty JSON payload
func SendEmptySuccess(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"status_code":    http.StatusNoContent,
		"status_message": "success: results not found",
		"response":       []struct{}{}},
	)
}

// SendNotFound sends 404 JSON payload
func SendNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status_code":    http.StatusNotFound,
		"status_message": "error: endpoint not found",
		"response":       []struct{}{}},
	)
}
