package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/gojsonq/v2"
	"gopkg.in/guregu/null.v4"
	"net/http"
	"net/url"
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

// parseFloat is a light wrapper around ParseFloat with bound checks
func parseFloat(query string, low, high, def int) int {
	if query == "" {
		return def
	} else {
		i, err := strconv.ParseFloat(query, 10)
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

// filterQueryArr filters based on an string query on an array
func filterQueryArr(query string, value []null.String) bool {
	if query == "" {
		return true
	}
	for _, v := range value {
		if !v.IsZero() && strings.Contains(
			strings.ToLower(v.ValueOrZero()),
			strings.ToLower(query)) {
			return true
		}
	}
	return false
}

// filterBoolQuery filters based on an bool query
func filterBoolQuery(query string, value null.Bool, flip ...bool) bool {
	if value.IsZero() {
		return false
	}
	if query == "" {
		return true
	}
	boolean, err := strconv.ParseBool(query)
	if err != nil {
		return false
	}
	if len(flip) == 0 {
		return boolean == value.ValueOrZero()
	} else {
		return !(boolean == value.ValueOrZero())
	}
}

// filterIntQuery light wrapper around filterValueQuery
func filterIntQuery(query string, value null.Int, low, high int) bool {
	newValue := null.NewFloat(float64(value.Int64), value.Valid)
	return filterValueQuery(query, newValue, low, high)
}

// filterValueQuery filters based on an number value query
func filterValueQuery(query string, value null.Float, low, high int) bool {
	if query == "" {
		return true
	}
	valueParsed := int(value.ValueOrZero())
	if strings.HasPrefix(query, "!") {
		queryFloat := parseFloat(query[1:], low, high, -1)
		if queryFloat < 0 {
			return false
		}
		if value.IsZero() {
			return false
		}
		return valueParsed != queryFloat
	} else if strings.HasPrefix(query, "<=") {
		queryFloat := parseFloat(query[2:], low, high, -1)
		if queryFloat < 0 {
			return false
		}
		if value.IsZero() {
			return false
		}
		return valueParsed <= queryFloat
	} else if strings.HasPrefix(query, ">=") {
		queryFloat := parseFloat(query[2:], low, high, -1)
		if queryFloat < 0 {
			return false
		}
		if value.IsZero() {
			return false
		}
		return valueParsed >= queryFloat
	} else if strings.HasPrefix(query, ">") {
		queryFloat := parseFloat(query[1:], low, high, -1)
		if queryFloat < 0 {
			return false
		}
		if value.IsZero() {
			return false
		}
		return valueParsed > queryFloat
	} else if strings.HasPrefix(query, "<") {
		queryFloat := parseFloat(query[1:], low, high, -1)
		if queryFloat < 0 {
			return false
		}
		if value.IsZero() {
			return false
		}
		return valueParsed < queryFloat
	} else {
		queryFloat := parseFloat(query[1:], low, high, -1)
		if queryFloat < 0 {
			return false
		}
		if value.IsZero() {
			return false
		}
		return valueParsed == queryFloat
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

func sendSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response":       data},
	)
}

func sendEmptySuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response":       []struct{}{}},
	)
}

func sendNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status_code":    http.StatusNotFound,
		"status_message": "not found",
		"response":       nil},
	)
}

func sendServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"status_code":    http.StatusInternalServerError,
		"status_message": "server error",
		"response":       nil},
	)
}

func prefixHandler(query string) (string, string) {
	if strings.HasPrefix(query, "<=") ||
		strings.HasPrefix(query, ">=") {
		return query[:2], query[2:]
	} else if strings.HasPrefix(query, "<") ||
		strings.HasPrefix(query, ">") {
		return query[:1], query[1:]
	} else if strings.HasPrefix(query, "(") {
		return "startsWith", query[1:]
	} else if strings.HasPrefix(query, ")") {
		return "endsWith", query[1:]
	} else if strings.HasPrefix(query, "=") {
		return "=", query[1:]
	} else if strings.HasPrefix(query, "!") {
		return "!=", query[1:]
	} else {
		return "default", query
	}
}

func typeToOp(valueType string, op string) string {
	switch valueType {
	case "string":
		if op == "default" {
			return "contains"
		}
	default:
		if op == "default" {
			return "="
		}
	}
	return op
}

func reflectPreType(query, value string, op *string) interface{} {
	if strings.HasSuffix(query, "id") {
		*op = typeToOp("string", *op)
		return value
	}
	if v, err := strconv.ParseInt(value, 10, 64); err == nil {
		*op = typeToOp("notString", *op)
		return v
	} else if v, err := strconv.ParseFloat(value, 64); err == nil {
		*op = typeToOp("notString", *op)
		return v
	} else if v, err := strconv.ParseBool(value); err == nil {
		*op = typeToOp("notString", *op)
		return v
	} else {
		*op = typeToOp("string", *op)
		return value
	}
}

func autoQuery(jsonq *gojsonq.JSONQ, values url.Values, limit, offset int) interface{} {
	data := jsonq.Copy()
	data.Limit(limit).Offset(offset)
	for query, value := range values {
		if query == "limit" || query == "offset" {
			continue
		}
		for _, el := range value {
			op, cleanValue := prefixHandler(el)
			reflectedValue := reflectPreType(query, cleanValue, &op)
			data.Where(query, op, reflectedValue)
		}
	}
	return data.Get()
}
