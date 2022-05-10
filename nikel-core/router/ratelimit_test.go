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
	"time"
)

// TestRateLimiterAllow tests if the rate limiter allows
func TestRateLimiterAllow(t *testing.T) {
	// get rid of all router output
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	// get router, set rate limit to 20 reqs/s and only attach courses
	r := NewRouter().SetRateLimiter(20)
	r.Uncached.GET("/", func(c *gin.Context) {
		handlers.Get[database.Course](c, database.DB.CoursesData)
	})

	// set ticker to tick at a rate of 15 reqs/s
	ticker := time.NewTicker(time.Second / 15)
	done := make(chan bool)

	// send requests in another goroutine
	go func() {
		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case <-ticker.C:
				w := httptest.NewRecorder()
				req, _ := http.NewRequest("GET", "/", nil)
				r.Engine.ServeHTTP(w, req)

				assert.Equal(t, http.StatusOK, w.Code)
			}
		}
	}()

	// run the test for 1 second (a long test is annoying to deal with)
	// a 1 second test is long enough to catch a 25 reqs/s ratelimit
	// hopefully this is long enough for the requests to ramp up
	// if this test time is too short, then maybe we are screwed
	time.Sleep(1 * time.Second)
	done <- true
}

// TestRateLimiterBlocked tests if the rate limiter blocks
func TestRateLimiterBlocked(t *testing.T) {
	// get rid of all router output
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	// get router, set rate limit to 20 reqs/s and only attach courses
	r := NewRouter().SetRateLimiter(20)
	r.Uncached.GET("/", func(c *gin.Context) {
		handlers.Get[database.Course](c, database.DB.CoursesData)
	})

	ratelimited := false
	numRequests := 0

	// give max 5 seconds to run the test
	for start := time.Now(); time.Since(start) < time.Second*5; {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		r.Engine.ServeHTTP(w, req)

		// at some point there should be a 429
		// if there isn't, then there's something
		// really wrong with the throughput
		if w.Code == http.StatusTooManyRequests {
			ratelimited = true
			break
		}

		// if not 429, the then response should be OK
		// if not, then there's something really bad going on
		assert.Equal(t, http.StatusOK, w.Code)
		numRequests++
	}

	assert.True(t, ratelimited, "rate limit should be reached")

	// should reach at least 20 requests because its ratelimited at 20reqs/s
	assert.GreaterOrEqual(t, numRequests, 20, "should be at least 20")
}
