package prints

import (
	"io/fs"
	"log"
	"os"
)

func getPathContent(path string) []fs.DirEntry {
	contents, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return contents
}
