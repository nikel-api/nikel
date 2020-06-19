package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/gojsonq/v2"
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
