package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	pathPrefix := ""
	wd, _ := os.Getwd()
	if filepath.Base(wd) == "codegen" {
		pathPrefix = "../"
	}

	rawData, err := ioutil.ReadFile(pathPrefix + "codegen/generic.tmpl")
	if err != nil {
		panic(err)
	}

	fileInfos, err := ioutil.ReadDir(pathPrefix + "nikel-parser/data")
	if err != nil {
		panic(err)
	}

	for _, file := range fileInfos {
		name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

		data := string(rawData)
		data = strings.ReplaceAll(data, "{{GenericType}}", strings.Title(name))
		data = strings.ReplaceAll(data, "{{genericType}}", name)
		nameSingle := strings.Title(name)
		if strings.HasSuffix(nameSingle, "s") {
			nameSingle = nameSingle[:len(nameSingle)-1]
		}
		data = strings.ReplaceAll(data, "{{GenericTypeSingle}}", nameSingle)

		ioutil.WriteFile(
			pathPrefix+"nikel-core/handlers/gen"+strings.Title(name)+".go",
			[]byte(data),
			0644,
		)
	}
}
