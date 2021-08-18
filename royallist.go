package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"royallist/icons"
)

func main() {
	path := "./"
	if len(os.Args[1:]) > 0 {
		path = os.Args[1]
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	directoryArr := make([]string, 0, len(files))
	fileArr := make([]string, 0, len(files))

	for _, f := range files {
		fmt.Println("raw: ", f.Name())
		if f.IsDir() {
			directoryArr = append(directoryArr, f.Name())
		} else {
			fileArr = append(fileArr, f.Name())
		}
	}

	for _, directory := range directoryArr {
		fmt.Printf("%s %s\n", icons.Names["directory"], directory)
	}

	for _, file := range fileArr {
		extension := filepath.Ext(file)

		if icon, iconExists := icons.Filetypes[extension]; iconExists {
			fmt.Printf("%s %s\n", icon, file)
			continue
		}

		if icon, iconExists := icons.Names[file]; iconExists {
			fmt.Printf("%s %s\n", icon, file)
			continue
		}

		fmt.Printf("%s %s\n", icons.Names["file"], file)
	}
}
