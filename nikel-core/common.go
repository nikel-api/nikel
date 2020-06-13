package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gopkg.in/guregu/null.v4"
	"strconv"
	"strings"
)

// parseInt is a light wrapper around ParseInt with bound checks
func parseInt(query string, low, high, def int) int {
	if query == "" {
		return def
	} else {
		i, err := strconv.ParseInt(query, 10, 64)
		val := int(i)
		if err != nil || val < low || val > high {
			return def
		} else {
			return val
		}
	}
}

// filterQuery filters based on an string query
func filterQuery(query string, value null.String) bool {
	if query == "" {
		return true
	}
	if !value.IsZero() && strings.Contains(
		strings.ToLower(value.ValueOrZero()),
		strings.ToLower(query)) {
		return true
	}
	return false
}

// filterBoolQuery filters based on an bool query
func filterBoolQuery(query string, value null.Bool) bool {
	if query == "" || value.IsZero() {
		return true
	}
	return value.ValueOrZero()
}

// filterIntQuery filters based on an int query
func filterIntQuery(query string, value null.Int, low, high int) bool {
	if query == "" {
		return true
	}
	valueParsed := int(value.ValueOrZero())
	if strings.HasPrefix(query, "!") {
		queryInt := parseInt(query[1:], low, high, -1)
		if queryInt < 0 {
			return false
		}
		if value.IsZero() {
			return false
		}
		return valueParsed != queryInt
	} else if strings.HasPrefix(query, "<=") {
		queryInt := parseInt(query[2:], low, high, -1)
		if queryInt < 0 {
			return false
		}
		if value.IsZero() {
			return false
		}
		return valueParsed <= queryInt
	} else if strings.HasPrefix(query, ">=") {
		queryInt := parseInt(query[2:], low, high, -1)
		if queryInt < 0 {
			return false
		}
		if value.IsZero() {
			return false
		}
		return valueParsed >= queryInt
	} else if strings.HasPrefix(query, ">") {
		queryInt := parseInt(query[1:], low, high, -1)
		if queryInt < 0 {
			return false
		}
		if value.IsZero() {
			return false
		}
		return valueParsed > queryInt
	} else if strings.HasPrefix(query, "<") {
		queryInt := parseInt(query[1:], low, high, -1)
		if queryInt < 0 {
			return false
		}
		if value.IsZero() {
			return false
		}
		return valueParsed < queryInt
	} else {
		queryInt := parseInt(query[1:], low, high, -1)
		if queryInt < 0 {
			return false
		}
		if value.IsZero() {
			return false
		}
		return valueParsed == queryInt
	}
}

// runCounter runs counter for usage metrics
func runCounter() {
	resp, _ := HttpClient.Get(MetricApi)
	if resp != nil {
		defer resp.Body.Close()
	}
}

// usageMetrics handles usage metrics
func usageMetrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		go runCounter()
	}
}

// urlToStruct loads url response into target
func urlToStruct(url string, target interface{}) error {
	resp, err := HttpClient.Get(url)

	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(target)

	errClose := resp.Body.Close()

	if errClose != nil {
		return errClose
	}

	if err != nil {
		return err
	}

	return nil
}

// Metric struct for usage metrics
type Metric struct {
	Value int `json:"value"`
}
