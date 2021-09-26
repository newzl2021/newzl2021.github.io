package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetZHFiles() []string {
	var files []string
	root := "./zh"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

func GetZHComponents() (header []byte, footer []byte) {
	root := "./zh"
	{
		fileFd, err := os.Open(root + "/components/header.html")
		if err != nil {
			panic(err)
		}
		defer fileFd.Close()
		header, err = ioutil.ReadAll(fileFd)
		if err != nil {
			panic(err)
		}
	}

	{
		fileFd, err := os.Open(root + "/components/footer.html")
		if err != nil {
			panic(err)
		}
		defer fileFd.Close()
		footer, err = ioutil.ReadAll(fileFd)
		if err != nil {
			panic(err)
		}
	}

	return
}

func WriteZHContent(fileName string, content string) {
	if err := ioutil.WriteFile(fileName, []byte(content), 0644); err != nil {
		fmt.Println("err: ", err)
	}
}

func main() {
	files := GetZHFiles()
	header, footer := GetZHComponents()
	for _, file := range files {
		// read file content
		fileFd, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		// merge file content
		if content, err := ioutil.ReadAll(fileFd); err == nil {
			WriteZHContent("../"+file, string(header)+string(content)+string(footer))
		}
		fileFd.Close()
	}
}
