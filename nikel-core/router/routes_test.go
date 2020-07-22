package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/nikel-api/nikel/nikel-core/config"
	"github.com/nikel-api/nikel/nikel-core/database"
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

var json = jsoniter.ConfigFastest

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
	r.Uncached.GET("/", handlers.GetCourses)

	// Random seed
	rand.Seed(time.Now().UnixNano())

	// Generate shuffled limits from 1-100
	limits := makeRange(1, 100)
	rand.Shuffle(len(limits), func(i, j int) { limits[i], limits[j] = limits[j], limits[i] })

	for _, limit := range limits {
		go func(limit int) {
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

			// Check if length matches limit value
			assert.Equal(t, limit, len(resp["response"].([]interface{})))
		}(limit)
	}
}

// TestOffset tests the offset option
func TestOffset(t *testing.T) {
	// Get rid of all router output
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	// Get router and only attach courses
	r := NewRouter()
	r.Uncached.GET("/", handlers.GetCourses)

	// Random seed
	rand.Seed(time.Now().UnixNano())

	// Load courses database
	coursesData := database.LoadFile(config.CoursesPath)

	// Generate shuffled offsets for all course elements
	offsets := makeRange(0, coursesData.Count()-1)
	rand.Shuffle(len(offsets), func(i, j int) { offsets[i], offsets[j] = offsets[j], offsets[i] })

	for _, offset := range offsets {
		go func(offset int) {
			// Make thread safe copy
			coursesDataCopy := coursesData.Copy()

			w := httptest.NewRecorder()
			params := url.Values{"offset": []string{strconv.Itoa(offset)}}
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

			// Check if offset matches offset in raw data
			assert.Equal(t,
				coursesDataCopy.From(fmt.Sprintf("[%d]", offset)).Get().(map[string]interface{})["id"],
				resp["response"].([]interface{})[0].(map[string]interface{})["id"],
			)
		}(offset)
	}
}
