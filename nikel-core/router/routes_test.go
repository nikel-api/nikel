package router

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/nikel-api/nikel/nikel-core/handlers"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// makeRange generates a sequence of numbers
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// TestLimit tests the limit option
func TestLimit(t *testing.T) {

	// Get rid of all router output
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	// Get router and only attach courses
	r := NewRouter()
	r.Engine.GET("/", handlers.GetCourses)

	// Random seed
	rand.Seed(time.Now().UnixNano())

	// Generate shuffled limits from 1-100
	limits := makeRange(1, 100)
	rand.Shuffle(len(limits), func(i, j int) { limits[i], limits[j] = limits[j], limits[i] })

	for _, limit := range limits {
		w := httptest.NewRecorder()
		params := url.Values{"limit": []string{strconv.Itoa(limit)}}
		req, _ := http.NewRequest(
			"GET",
			"/?"+params.Encode(),
			nil,
		)
		r.Engine.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		resp := map[string]interface{}{}
		err := json.Unmarshal([]byte(w.Body.String()), &resp)
		assert.Nil(t, err)

		assert.Equal(t, "success", resp["status_message"])
		assert.Equal(t, limit, len(resp["response"].([]interface{})))
	}
}
