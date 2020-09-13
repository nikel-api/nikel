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

	req, _ := http.NewRequest(
		"GET",
		"/",
		nil,
	)

	t.Run("WithoutHeader", func(t *testing.T) {
		w := httptest.NewRecorder()
		r.Engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		// Content-Encoding header should be empty
		assert.Empty(t, w.Header().Get("Content-Encoding"))
	})

	t.Run("WithHeader", func(t *testing.T) {
		// send gzip accept request
		req.Header.Set("Accept-Encoding", "gzip")

		w := httptest.NewRecorder()
		r.Engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		// Content-Encoding header should be gzip
		assert.Equal(t, "gzip", w.Header().Get("Content-Encoding"))
	})
}
