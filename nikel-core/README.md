# Nikel Core

This README provides a high-level overview of Nikel-Core.

### Benchmarks

Benchmark code: [GitHub Gist](https://gist.github.com/darenliang/2caaf2816908d3d95f9e112db1e02929)

Date: 07/25/2020

Commit ID: d20962

AMD Ryzen 7 3800X on localhost with standard output disabled

| Cache           | Concurrent Users | Hatch Rate  | Throughput    |
|-----------------|------------------|-------------|---------------|
| In-Memory       | 10000            | 100 users/s | 2474.4 reqs/s |
| LevelDB   (SSD) | 10000            | 100 users/s | 2382.5 reqs/s |
| BadgerDB  (SSD) | 10000            | 100 users/s | 2372.5 reqs/s |
| None            | 10000            | 100 users/s | 1639.6 reqs/s |

Note that benchmark results don't reflect real world performance so take the results with a grain of salt.

### Web Framework

Nikel-Core is powered by [Gin](https://github.com/gin-gonic/gin), a performant and featureful web framework.

### Database

All data is stored in memory through variables for low latency and fast querying. (with thread-safe access via immutability and read-only access)

### Querying

The querying system is based on [Gojsonq](https://github.com/thedevsaddam/gojsonq). It maps the REST API interface to Gojsonq query calls.

Using [guregu/Null](https://github.com/guregu/null), the queried data is passed through a struct that formats the data in the proper JSON structure.

### Caching

Data is cached via a Gin cache middleware backed by GoLevelDB. [(nikel-cache)](https://github.com/nikel-api/nikel-cache)

The cache can be optionally backed by memory or BadgerDB.

### CORS

Nikel-Core by default allows all origins so that the API is accessible to everyone.

### Gzip

Nikel-Core gzips all responses via gzip for smaller payload sizes.

### Rate Limits

Rate limits are handled in memory via [ulule/limiter](https://github.com/ulule/limiter).

### Process Flow

Request -> Ratelimit* -> Cache* -> Cors* -> Gzip* -> Route handler -> Query engine -> Response

*Optional