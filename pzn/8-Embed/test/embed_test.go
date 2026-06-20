package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestStringEmbed(t *testing.T) {
	fmt.Println(version)
}

//go:embed files/download.jpeg
var image []byte

func TestFileEmebd(t *testing.T) {
	err := os.WriteFile("logo_new.jpeg", image, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFile(t *testing.T){
	a, _ := files.ReadFile("files/a.txt")
	b, _ := files.ReadFile("files/b.txt")
	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(a))
	fmt.Println(string(b))
	fmt.Println(string(c))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T){
	dir, _ := path.ReadDir("files")

	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())

			content, _ := path.ReadFile("files/" + entry.Name())

			fmt.Println(string(content))
		}
	}
}