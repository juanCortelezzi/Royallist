package prints

import (
	"fmt"
	"path/filepath"
	"royallist/icons"
)

func RecursivePrint(path string) {
	privateRecursivePrint(path, 0)
}

func privateRecursivePrint(path string, level int) {
	contents := getPathContent(path)

	var dirArr []string
	var fileArr []string

	for _, f := range contents {
		if f.IsDir() {
			dirArr = append(dirArr, f.Name())
		} else {
			fileArr = append(fileArr, f.Name())
		}
	}

	separator := ""
	for i := 0; i < level; i++ {
		separator += "│ "
	}

	for i, dir := range dirArr {
		if i == len(dirArr)-1 && len(fileArr) == 0 {
			fmt.Printf("%s", separator+"└─")
		} else {
			fmt.Printf("%s", separator+"├─")
		}
		fmt.Printf("%s %s\n", icons.Names["directory"], dir)
		privateRecursivePrint(path+"/"+dir, level+1)
	}

	for i, file := range fileArr {
		extension := filepath.Ext(file)

		if i == len(fileArr)-1 {
			fmt.Printf("%s", separator+"└─")
		} else {
			fmt.Printf("%s", separator+"├─")
		}

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
