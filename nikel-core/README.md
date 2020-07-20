# Nikel Core

This README will provide a high-level overview of Nikel-Core's architecture.

### Web Framework

Nikel-Core is powered by [Gin](https://github.com/gin-gonic/gin), a relatively performant and featureful web framework.

### Database

All data is stored in memory through variables for low latency and fast querying. (with thread-safe access via immutability and read-only access)

### Querying

The querying system is based on [Gojsonq](https://github.com/thedevsaddam/gojsonq). It maps the REST API interface to Gojsonq query calls.

Using [guregu/Null](https://github.com/guregu/null), the queried data is passed through a struct that formats the data in the proper JSON structure.

All JSON is marshalled and unmarshalled via [jsoniter](https://github.com/json-iterator/go) for performance.

### Caching

Data is cached via an implementation of [LevelDB in Go](https://github.com/syndtr/goleveldb).

### CORS

Nikel-Core by default allows all origins so that the API is accessible to everyone.

### Rate Limits

Rate limits are handled in memory via [ulule/limiter](https://github.com/ulule/limiter).