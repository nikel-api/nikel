package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// get exact location of the codegen folder
	pathPrefix := ""
	wd, _ := os.Getwd()
	if filepath.Base(wd) == "codegen" {
		pathPrefix = "../"
	}

	// read generic.tmpl
	rawData, err := ioutil.ReadFile(pathPrefix + "codegen/generic.tmpl")
	if err != nil {
		panic(err)
	}

	// get file infos of files in Nikel's dataset folder
	fileInfos, err := ioutil.ReadDir(pathPrefix + "nikel-datasets/data")
	if err != nil {
		panic(err)
	}

	for _, file := range fileInfos {
		// ignore non json files
		if filepath.Ext(file.Name()) != ".json" {
			continue
		}

		// remove suffix
		name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

		// replace templated data
		data := string(rawData)
		data = strings.ReplaceAll(data, "{{GenericType}}", strings.Title(name))
		data = strings.ReplaceAll(data, "{{genericType}}", name)

		// handle messed up linguistic exceptions with filenames ending with s
		nameSingle := strings.Title(name)
		if strings.HasSuffix(nameSingle, "s") {
			nameSingle = nameSingle[:len(nameSingle)-1]
		}

		// replace some more
		data = strings.ReplaceAll(data, "{{GenericTypeSingle}}", nameSingle)

		// write file
		_ = ioutil.WriteFile(
			pathPrefix+"nikel-core/handlers/gen"+strings.Title(name)+".go",
			[]byte(data),
			0644,
		)
	}
}
