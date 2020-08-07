package query

import (
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq/v2"
	"net/url"
	"strconv"
	"strings"
)

// twoCharPrefixMap represents the map of two char prefixes
var twoCharPrefixMap = map[string]string{
	"<=": "<=",
	">=": ">=",
}

// oneCharPrefixMap represents the map of one char prefixes
var oneCharPrefixMap = map[string]string{
	"<": "<",
	">": ">",
	"(": "startsWith",
	")": "endsWith",
	"=": "=",
	"!": "!=",
	"~": "interface",
}

// opTypeMap maps the ops to a specific type.
// This improves the query performance by
// limiting the number of inferred ops on lookup.
// This mapping is defined in https://docs.nikel.ml/docs/query_guide.
var opTypeMap = map[string][]string{
	"default":    {"string", "numerical"},
	"=":          {"string", "numerical"},
	"!=":         {"string", "numerical"},
	"<":          {"numerical"},
	"<=":         {"numerical"},
	">":          {"numerical"},
	">=":         {"numerical"},
	"startsWith": {"string"},
	"endsWith":   {"string"},
	"interface":  {"string"},
}

// prefixHandler determines the prefix for each query
func prefixHandler(query string) (string, string) {

	// look up two char prefixes
	queryLen := len(query)
	if queryLen >= 2 {
		if val, ok := twoCharPrefixMap[query[:2]]; ok {
			return val, query[2:]
		}
	}

	// look up one char prefixes
	if queryLen >= 1 {
		if val, ok := oneCharPrefixMap[query[:1]]; ok {
			return val, query[1:]
		}
	}

	return "default", query
}

// whereWrapper wraps jsonq queries so it creates
// an initial query and orWhere's for subsequent queries.
// The initial flag is really messed, but its unavoidable
// unless a struct wrapper is used.
func whereWrapper(jsonq *gojsonq.JSONQ, initial *bool, query, op string, value interface{}) {
	if *initial {
		jsonq.Where(query, op, value)
		*initial = false
	} else {
		jsonq.OrWhere(query, op, value)
	}
}

// queryBuilder builds queries based on reflected type
func queryBuilder(jsonq *gojsonq.JSONQ, query, op, value string) {
	// flag for initial jsonq query
	initial := true

	// add queries
	if val, ok := opTypeMap[op]; ok {
		// loop through possible type mappings
		for _, t := range val {
			switch t {
			case "numerical":
				// handle default mapping for numerical
				newOp := op
				if newOp == "default" {
					newOp = "="
				}
				// handle floats
				if v, err := strconv.ParseFloat(value, 64); err == nil {
					whereWrapper(jsonq, &initial, query, newOp, v)
				}
				// handle booleans
				if v, err := strconv.ParseBool(value); err == nil {
					whereWrapper(jsonq, &initial, query, newOp, v)
				}
			case "string":
				// handle default mapping for string
				if op == "default" {
					whereWrapper(jsonq, &initial, query, "contains", value)
				} else {
					whereWrapper(jsonq, &initial, query, op, value)
				}
			}
		}
	}

	jsonq.More()
}

// AutoQuery queries url data along with limit and offset information
func AutoQuery(jsonq *gojsonq.JSONQ, values url.Values, limit, offset int) interface{} {
	// copy to make thread-safe
	data := jsonq.Copy()

	// loop through query params
	for query, value := range values {
		// filter out limit and offset
		if query == "limit" || query == "offset" {
			continue
		}

		// loop through duplicates
		for _, el := range value {
			// handle prefix and separate by op and value
			op, cleanValue := prefixHandler(el)
			// build query
			queryBuilder(data, query, op, cleanValue)
		}
	}

	// set limit and offset
	data.Limit(limit).Offset(offset)

	// get result
	return data.Get()
}

// InterfaceMacro matches by interface value,
func InterfaceMacro(value interface{}, key interface{}) (bool, error) {
	// check if string
	keyString, ok := key.(string)

	if !ok {
		return false, fmt.Errorf("%v must be a string", key)
	}

	// marshal interface value to bytes
	rawBytes, err := json.Marshal(value)

	if err != nil {
		return false, err
	}

	// check if substring contains in string (case insensitive)
	// it's slow but it is slightly faster than the regex implementation
	// this could be further optimized
	return strings.Contains(strings.ToLower(string(rawBytes)), strings.ToLower(keyString)), nil
}
