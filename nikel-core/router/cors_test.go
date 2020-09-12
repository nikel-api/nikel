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

// TestCors tests cors middleware
func TestCors(t *testing.T) {
	// get rid of all router output
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	// get router and only attach courses
	r := NewRouter().SetAllowCors()
	r.Uncached.GET("/", handlers.GetCourses)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		"GET",
		"/",
		nil,
	)

	// send request from foo.com
	req.Header.Set("Origin", "foo.com")

	r.Engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	// Access-Control-Allow-Origin header should be *
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
}
