package prints

import (
	"fmt"
	"path/filepath"
	"royallist/icons"
)

func CommonPrint(path string) {
	contents := getPathContent(path)

	var fileArr []string

	for _, f := range contents {
		if f.IsDir() {
			fmt.Printf("%s %s\n", icons.Names["directory"], f.Name())
		} else {
			fileArr = append(fileArr, f.Name())
		}
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
