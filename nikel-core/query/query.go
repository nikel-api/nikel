package query

import (
	"github.com/thedevsaddam/gojsonq/v2"
	"net/url"
	"strconv"
	"strings"
)

// PrefixHandler determines the prefix for each query
func PrefixHandler(query string) (string, string) {
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

// TypeToOp handles differing behavior between strings and non-strings
func TypeToOp(valueType string, op string) string {
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

// QueryBuilder builds queries based on reflected type
func QueryBuilder(jsonq *gojsonq.JSONQ, query, op, value string) {
	jsonq.Where(query, TypeToOp("string", op), value)
	newOp := TypeToOp("notString", op)
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

// AutoQuery queries url data along with limit and offset information
func AutoQuery(jsonq *gojsonq.JSONQ, values url.Values, limit, offset int) interface{} {
	data := jsonq.Copy()
	for query, value := range values {
		if query == "limit" || query == "offset" {
			continue
		}
		for _, el := range value {
			op, cleanValue := PrefixHandler(el)
			QueryBuilder(data, query, op, cleanValue)
		}
	}
	data.Limit(limit).Offset(offset)
	return data.Get()
}
