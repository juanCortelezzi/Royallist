package prints

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"royallist/icons"
)

func CommonPrint(path string) {
	contents := getPathContent(path)
	var b bytes.Buffer

	for _, f := range contents {
		if f.IsDir() {
			b.WriteString(
				fmt.Sprintf("%s %s\n", icons.Names["directory"], f.Name()),
			)
		}
	}

	for _, f := range contents {
		if f.IsDir() {
			continue
		}
		extension := filepath.Ext(f.Name())

		if icon, iconExists := icons.Filetypes[extension]; iconExists {

			b.WriteString(fmt.Sprintf("%s %s\n", icon, f.Name()))
			continue
		}

		if icon, iconExists := icons.Names[f.Name()]; iconExists {

			b.WriteString(fmt.Sprintf("%s %s\n", icon, f.Name()))
			continue
		}

		b.WriteString(fmt.Sprintf("%s %s\n", icons.Names["file"], f.Name()))
	}

	_, err := b.WriteTo(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
