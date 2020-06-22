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

// sendSuccess sends successful JSON payload
func sendSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success",
		"response":       data},
	)
}

// sendEmptySuccess sends successful empty JSON payload
func sendEmptySuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status_code":    http.StatusOK,
		"status_message": "success: results not found",
		"response":       []struct{}{}},
	)
}

// sendNotFound sends 404 JSON payload
func sendNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status_code":    http.StatusNotFound,
		"status_message": "error: endpoint not found",
		"response":       []struct{}{}},
	)
}

// prefixHandler determines the prefix for each query
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

// typeToOp handles differing behavior between strings and non-strings
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

// queryBuilder builds queries based on reflected type
func queryBuilder(jsonq *gojsonq.JSONQ, query, op, value string) {
	jsonq.Where(query, typeToOp("string", op), value)
	newOp := typeToOp("notString", op)
	if v, err := strconv.ParseInt(value, 10, 64); err == nil {
		jsonq.OrWhere(query, newOp, v)
	}
	if v, err := strconv.ParseFloat(value, 64); err == nil {
		jsonq.OrWhere(query, newOp, v)
	}
	if v, err := strconv.ParseBool(value); err == nil {
		jsonq.OrWhere(query, newOp, v)
	}
	jsonq.More()
}

// autoQuery queries url data along with limit and offset information
func autoQuery(jsonq *gojsonq.JSONQ, values url.Values, limit, offset int) interface{} {
	data := jsonq.Copy()
	for query, value := range values {
		if query == "limit" || query == "offset" {
			continue
		}
		for _, el := range value {
			op, cleanValue := prefixHandler(el)
			queryBuilder(data, query, op, cleanValue)
		}
	}
	data.Limit(limit).Offset(offset)
	return data.Get()
}
