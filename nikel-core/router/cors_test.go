package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nikel-api/nikel/nikel-core/database"
	"github.com/nikel-api/nikel/nikel-core/handlers"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCors tests cors middleware
func TestCors(t *testing.T) {
	// get rid of all router output
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	// get router and only attach courses
	r := NewRouter().SetAllowCors()
	r.Uncached.GET("/", func(c *gin.Context) {
		handlers.Get[database.Course](c, database.DB.CoursesData)
	})

	req, _ := http.NewRequest(
		"GET",
		"/",
		nil,
	)

	t.Run("WithoutOrigin", func(t *testing.T) {
		w := httptest.NewRecorder()
		r.Engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		// Access-Control-Allow-Origin header should empty
		assert.Empty(t, w.Header().Get("Access-Control-Allow-Origin"))
	})

	t.Run("WithOrigin", func(t *testing.T) {
		// send request from foo.com
		req.Header.Set("Origin", "foo.com")

		w := httptest.NewRecorder()
		r.Engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		// Access-Control-Allow-Origin header should be *
		assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	})
}
