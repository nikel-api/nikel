package router

import (
	"github.com/gin-gonic/gin"
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
	// Get rid of all router output
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	// Get router, set rate limit to 20 reqs/s and only attach courses
	r := NewRouter().SetRateLimiter(20)
	r.Uncached.GET("/", handlers.GetCourses)

	// Set ticker to tick at a rate of 15 reqs/s
	ticker := time.NewTicker(time.Second / 15)
	done := make(chan bool)

	// Send requests in another goroutine
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

				assert.Equal(t, w.Code, http.StatusOK)
			}
		}
	}()

	// Run the test for 1 seconds (a long test is annoying to deal with).
	// A 1 second test is long enough to catch a 25 reqs/s ratelimit.
	// Hopefully this is long enough for the requests to ramp up.
	// If this test time is too short, then maybe we are screwed.
	time.Sleep(1 * time.Second)
	done <- true
}

// TestRateLimiterBlocked tests if the rate limiter blocks
func TestRateLimiterBlocked(t *testing.T) {
	// Get rid of all router output
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard

	// Get router, set rate limit to 20 reqs/s and only attach courses
	r := NewRouter().SetRateLimiter(20)
	r.Uncached.GET("/", handlers.GetCourses)

	ratelimited := false
	numRequests := 0

	// Give max 5 seconds to run the test
	for start := time.Now(); time.Since(start) < time.Second*5; {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		r.Engine.ServeHTTP(w, req)

		// At some point there should be a 429.
		// If there isn't, then there's something
		// really wrong with the throughput
		if w.Code == http.StatusTooManyRequests {
			ratelimited = true
			break
		}

		assert.Equal(t, w.Code, http.StatusOK)
		numRequests += 1
	}

	assert.True(t, ratelimited, "rate limit should be reached")
	assert.GreaterOrEqual(t, numRequests, 20, "should be at least 20")
}
