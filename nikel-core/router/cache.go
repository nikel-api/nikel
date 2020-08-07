package router

import (
	"fmt"
	"github.com/nikel-api/nikel-cache"
	"os"
	"strconv"
	"time"
)

// SetLevelDBCache sets a LevelDB backed cache
func (r *Router) SetLevelDBCache(expires ...time.Duration) *Router {
	// get env variable
	cacheExpiryValueStr := os.Getenv("CACHE_EXPIRY")

	// check if env variable is set
	if len(cacheExpiryValueStr) == 0 && len(expires) == 0 {
		fmt.Println("[NIKEL-CORE] LevelDB cache not set.")
		return r
	}

	var cacheExpiryValue time.Duration

	if len(expires) != 0 {
		cacheExpiryValue = expires[0]
	} else {
		var err error
		tmp, err := strconv.Atoi(cacheExpiryValueStr)

		cacheExpiryValue = time.Second * time.Duration(tmp)

		if err != nil {
			panic(err)
		}

		if cacheExpiryValue < 0 {
			panic(fmt.Errorf("nikel-core: invalid cache expiry value %d", cacheExpiryValue))
		}
	}

	// attach only cached group
	r.Cached.Use(cache.New(cache.Options{
		Store: func() *cache.LevelDB {
			store, err := cache.NewLevelDB("cache")
			if err != nil {
				panic(err)
			}
			return store
		}(),
		Expire:        cacheExpiryValue,
		Headers:       []string{},
		StripHeaders:  []string{"X-Ratelimit-Limit", "X-Ratelimit-Remaining", "X-Ratelimit-Reset"},
		DoNotUseAbort: false,
	}))

	if cacheExpiryValue == 0 {
		fmt.Println("[NIKEL-CORE] Set LevelDB cache to never expire.")
	} else {
		fmt.Printf("[NIKEL-CORE] Set LevelDB cache to expire after %d seconds.\n", cacheExpiryValue/time.Second)
	}

	return r
}
