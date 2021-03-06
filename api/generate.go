// Code generated by grpc-kit-cli. DO NOT EDIT.

// +build ignore

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/shurcooL/vfsgen"
)

//go:generate go run generate.go
func main() {
	var fs http.FileSystem = http.Dir("./doc")

	err := vfsgen.Generate(fs, vfsgen.Options{
		Filename:        "assets_vfsdata.go",
		PackageName:     "api",
		VariableName:    "Assets",
		VariableComment: "Assets user for static swagger file",
	})

	if err != nil {
		fmt.Println("Generate api doc err:", err)
		os.Exit(1)
	}
}