// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

var assetsFs http.FileSystem = http.Dir("../dist/frontend")

func main() {
	if err := vfsgen.Generate(assetsFs, vfsgen.Options{
		PackageName:  "main",
		Filename:     "assets.go",
		VariableName: "assetsFs",
		BuildTags:    "build",
	}); err != nil {
		log.Fatalln(err)
	}
}
