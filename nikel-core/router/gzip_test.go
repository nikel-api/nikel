package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nikel-api/nikel/nikel-core/handlers"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGzip tests gzip middleware
func TestGzip(t *testing.T) {
	// get rid of all router output
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	// get router and only attach courses
	r := NewRouter().SetGzip()
	r.Uncached.GET("/", handlers.GetCourses)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"GET",
		"/",
		nil,
	)

	// send gzip accept request
	req.Header.Set("Accept-Encoding", "gzip")

	r.Engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// Content-Encoding header should be gzip
	assert.Equal(t, "gzip", w.Header().Get("Content-Encoding"))
}
