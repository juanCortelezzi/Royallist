package main

import (
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
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

	filDirectories := make([]string, 0, len(files))
	filFiles := make([]string, 0, len(files))

	for _, f := range files {
		if f.IsDir() {
			filDirectories = append(filDirectories, f.Name())
		} else {
			filFiles = append(filFiles, f.Name())
		}
	}

	green := color.FgGreen.Render
	red := color.FgRed.Render
	yellow := color.FgYellow.Render
	cyan := color.FgCyan.Render
	magenta := color.FgMagenta.Render
	blue := color.FgBlue.Render
	white := color.FgWhite.Render

	for _, d := range filDirectories {
		fmt.Printf("%s %s\n", cyan(""), d)
	}

	icons := map[string]string{
		".h":    magenta(""),
		".c":    magenta(""),
		".vim":  green(""),
		".scss": red(""),
		".html": red(""),
		".json": yellow(""),
		".yml":  yellow(""),
		".js":   yellow(""),
		".jsx":  yellow(""),
		".cpp":  blue(""),
		".lua":  blue(""),
		".py":   blue(""),
		".go":   blue("ﳑ"),
		".css":  blue(""),
		".ts":   blue(""),
		".tsx":  blue(""),
		".sh":   white(""),
		".txt":  white(""),
		".rs":   white(""),
		".iso":  white(""),
		".md":   white(""),
		".mdx":  white(""),
		".wiki": white(""),
		".mod":  white(""),
		".sum":  white(""),
		".toml": white(""),
	}

	for _, f := range filFiles {
		extension := filepath.Ext(f)
		icon, iconExists := icons[extension]
		if !iconExists {
			switch f {
			case ".gitignore":
				icon = magenta("")
			default:
				icon = green("")
			}
		}
		fmt.Printf("%s %s\n", icon, f)
	}
}
