package query

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/thedevsaddam/gojsonq/v2"
	"net/url"
	"strconv"
	"strings"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

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
	} else if strings.HasPrefix(query, "~") {
		return "interface", query[1:]
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

// AutoQuery queries url data along with limit and offset information
func AutoQuery(jsonq *gojsonq.JSONQ, values url.Values, limit, offset int) interface{} {
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

// InterfaceMacro matches by interface value,
// it's slow but it is slightly faster than the regex implementation
func InterfaceMacro(value interface{}, key interface{}) (bool, error) {
	keyString, ok := key.(string)

	if !ok {
		return false, fmt.Errorf("%v must be a string", key)
	}

	rawBytes, err := json.Marshal(value)

	if err != nil {
		return false, err
	}

	return strings.Contains(strings.ToLower(string(rawBytes)), strings.ToLower(keyString)), nil
}
