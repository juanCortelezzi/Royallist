package prints

import (
	"fmt"
	"path/filepath"
	"royallist/icons"
)

// old 18344 ns/op
// vs
// new 17816 ns/op
func CommonPrint(path string) {
	contents := getPathContent(path)

	for _, f := range contents {
		if f.IsDir() {
			fmt.Printf("%s %s\n", icons.Names["directory"], f.Name())
		}
	}

	for _, f := range contents {
		if f.IsDir() {
			continue
		}
		extension := filepath.Ext(f.Name())

		if icon, iconExists := icons.Filetypes[extension]; iconExists {
			fmt.Printf("%s %s\n", icon, f.Name())
			continue
		}

		if icon, iconExists := icons.Names[f.Name()]; iconExists {
			fmt.Printf("%s %s\n", icon, f.Name())
			continue
		}

		fmt.Printf("%s %s\n", icons.Names["file"], f.Name())
	}
}
