package router

import (
	"fmt"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"os"
	"strconv"
	"time"
)

// SetRateLimiter sets the rate limiter
func (r *Router) SetRateLimiter(opt ...int) *Router {
	ratelimitValueStr := os.Getenv("RATELIMIT")

	if len(ratelimitValueStr) == 0 && len(opt) == 0 {
		fmt.Println("[NIKEL-CORE] Ratelimit not set.")
		return r
	}

	ratelimitValueInt := 0

	if len(opt) != 0 {
		ratelimitValueInt = opt[0]
	} else {
		var err error
		ratelimitValueInt, err = strconv.Atoi(ratelimitValueStr)

		if err != nil {
			panic(err)
		}

		if ratelimitValueInt < 1 {
			panic(fmt.Errorf("nikel-core: invalid ratelimit value %d", ratelimitValueInt))
		}
	}

	ratelimit := limiter.Rate{
		Period: 1 * time.Second,
		Limit:  int64(ratelimitValueInt),
	}

	rateStore := memory.NewStore()
	rateInstance := limiter.New(rateStore, ratelimit)
	rateMiddleware := mgin.NewMiddleware(rateInstance)
	r.Engine.ForwardedByClientIP = true
	r.Engine.Use(rateMiddleware)

	fmt.Printf("[NIKEL-CORE] Ratelimit set to %d reqs/s.\n", ratelimitValueInt)
	return r
}
