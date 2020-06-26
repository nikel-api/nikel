# Nikel API

Nikel (pronunciation: `/'ni:kɛl/`) is a collection of data APIs on information about the University of Toronto.

### Documentation

[Nikel API Documentation](https://docs.nikel.ml/docs)

### Endpoints currently supported

* /courses
* /textbooks
* /exams
* /evals
* /food
* /services
* /buildings
* /parking

### Self Hosting

Please make sure you have the same go version displayed in the `go.mod` file. It should usually be the latest stable release. If you are unsure which go version you have, use `go version` to find out.

1. git clone
```
git clone https://github.com/nikel-api/nikel.git
```
2. cd into nikel-parser submodule
```
cd nikel/nikel-parser
```
3. update nikel-parser submodule to latest
```
git pull origin master
```
4. cd into nikel-core
```
cd ../nikel-core
```
5. build nikel-core
```
go build
```
6. run nikel-core
```
Windows
./nikel-core.exe

Linux and macOS
./nikel-core
```

### Contributing

For contributing, there are a few things to look out for:

* Always use `go fmt` to format code
* Consult the article [Godoc: documenting Go code](https://blog.golang.org/godoc) on how to write docstrings if you aren't 100% sure

If you find any inconsistencies or parts of code that can be reworked, any pull requests are greatly appreciated.
