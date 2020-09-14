package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nikel-api/nikel-cache"
	"github.com/nikel-api/nikel/nikel-core/handlers"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// TestLevelDBCache tests the LevelDB cache
func TestLevelDBCache(t *testing.T) {
	// get rid of all router output
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	// get router and only attach courses
	r := NewRouter()

	// setup cache store
	store, err := cache.NewLevelDB("cache_test")
	assert.Nil(t, err)

	// attach cache store
	r.Cached.Use(cache.New(cache.Options{
		Store: store,
	}))
	r.Cached.GET("/", handlers.GetCourses)

	// make two requests
	for i := 0; i < 2; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(
			"GET",
			"/",
			nil,
		)
		r.Engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		if i == 0 {
			// first should not have a x-gin-cache header defined
			assert.Empty(t, w.Header().Get("x-gin-cache"))
		} else {
			// second should have a x-gin-cache header defined
			assert.NotEmpty(t, w.Header().Get("x-gin-cache"))
		}
	}

	// do necessary cleanup
	err = store.DB.Close()
	assert.Nil(t, err)

	err = os.RemoveAll("cache_test")
	assert.Nil(t, err)
}
