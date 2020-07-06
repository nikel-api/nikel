<h1 align="center">
  <br>
  <a href="https://docs.nikel.ml"><img src="https://docs.nikel.ml/img/nikel-api-circle.png" alt="Nikel API" width="200"></a>
  <br>
  Nikel API
  <br>
</h1>

<h4 align="center">A collection of data APIs for the University of Toronto.</h4>

<p align="center">
  <a href="https://travis-ci.com/nikel-api/nikel">
    <img alt="Build Status" src="https://img.shields.io/travis/nikel-api/nikel">
  </a>
  <a href="https://status.nikel.ml/">
    <img alt="API Status" src="https://img.shields.io/uptimerobot/status/m785379986-9f61400de9d1a64fff1b0b51">
  </a>
  <a href="https://goreportcard.com/report/github.com/nikel-api/nikel">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/nikel-api/nikel">
  </a>
  <a href="https://github.com/nikel-api/nikel/blob/master/LICENSE">
    <img alt="License" src="https://img.shields.io/github/license/nikel-api/nikel">
  </a>
</p>

<p align="center">
  <a href="#documentation">Documentation</a> •
  <a href="#api-wrappers">API Wrappers</a> •
  <a href="#self-hosting">Self Hosting</a> •
  <a href="#contributing">Contributing</a> •
  <a href="#license">License</a>
</p>

## Documentation

[Nikel API Documentation](https://docs.nikel.ml)

## API Wrappers

#### Official

* [nikel-ts (Node.js)](https://www.npmjs.com/package/nikel)

#### Unofficial

* [nikel-rs (Rust)](https://crates.io/crates/nikel-rs)

Please feel free to submit a pull request to add your own API wrapper to this list!

## Self Hosting

Please make sure you have the same go version displayed in the `go.mod` file. It should usually be the latest stable release. If you are unsure which go version you have, use `go version` to find out.

Nikel should work on any 32/64 bit system with go installed.

1. git clone
```
git clone https://github.com/nikel-api/nikel.git
```
2. cd into nikel-parser submodule
```
cd nikel/nikel-parser
```
3. Update nikel-parser submodule to latest
```
git pull origin master
```
4. cd into nikel-core
```
cd ../nikel-core
```
5. Build nikel-core
```
go build
```
6. Run nikel-core
```
Windows
./nikel-core.exe

Linux and macOS
./nikel-core
```

7. Optional configuration

* By default, nikel-core should be listening and serving on port 8080. To change the port, modify the `PORT` environment variable.
* To suppress debug logs, add the environment variable `GIN_MODE` with the value `release`.
* To add optional application metrics via New Relic APM, add the environment variable `NEW_RELIC_LICENSE_KEY` along with a license key.

## Contributing

For contributing, there are a few things to look out for:

* Always use `go fmt` to format code
* Consult the article [Godoc: documenting Go code](https://blog.golang.org/godoc) on how to write docstrings if you aren't 100% sure

If you find any inconsistencies or parts of code that can be reworked, any pull requests are greatly appreciated.

## License

[MIT](https://github.com/nikel-api/nikel/blob/master/LICENSE)