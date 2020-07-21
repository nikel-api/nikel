# Nikel Core

This README will provide a high-level overview of Nikel-Core's architecture.

### Benchmarks

Benchmark code: [GitHub Gist](https://gist.github.com/darenliang/2caaf2816908d3d95f9e112db1e02929)

Date: 07/20/2020

Commit ID: 8e960a

AMD Ryzen 7 3800X on localhost with standard output disabled

| Local Cache | Concurrent Users | Hatch Rate  | Throughput    |
|-------------|------------------|-------------|---------------|
| Yes         | 10000            | 100 users/s | 2382.5 reqs/s |
| No          | 10000            | 100 users/s | 1639.6 reqs/s |

Note that benchmark results don't reflect real world performance so take the results with a grain of salt.

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
