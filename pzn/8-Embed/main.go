package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed test/version.txt
var version string

//go:embed test/files/download.jpeg
var image []byte

//go:embed test/files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := os.WriteFile("logo_new.jpeg", image, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := path.ReadDir("test/files")

	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())

			content, _ := path.ReadFile("test/files/" + entry.Name())

			fmt.Println(string(content))
		}
	}
}
