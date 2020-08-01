package router

import (
	"github.com/gin-gonic/gin"
	cache "github.com/nikel-api/nikel-cache"
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

	// Get rid of all router output
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	// Get router and only attach courses
	r := NewRouter()

	store, err := cache.NewLevelDB("cache_test")
	assert.Nil(t, err)

	r.Cached.Use(cache.New(cache.Options{
		Store: store,
	}))
	r.Cached.GET("/", handlers.GetCourses)

	for i := 0; i < 2; i++ {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(
			"GET",
			"/",
			nil,
		)
		r.Engine.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		if i == 0 {
			assert.Equal(t, "", w.Header().Get("x-gin-cache"))
		} else {
			assert.NotEqual(t, "", w.Header().Get("x-gin-cache"))
		}
	}

	err = store.DB.Close()
	assert.Nil(t, err)

	err = os.RemoveAll("cache_test")
	assert.Nil(t, err)
}