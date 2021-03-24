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
		".ts":   blue(""),
		".tsx":  blue(""),
		".json": yellow(""),
		".go":   blue("ﳑ"),
		".mod":  white(""),
		".sum":  white(""),
		".toml":  white(""),
		".yml":  yellow(""),
		".js":   yellow(""),
		".jsx":  yellow(""),
		".css":  blue(""),
		".scss":  red(""),
		".html": red(""),
		".c": magenta(""),
		".cpp": blue(""),
		".py": blue(""),
		".sh": white(""),
		".txt": white(""),
		".rs": white(""),
		".vim": green(""),
		".lua": blue(""),
	}

	for _, f := range filFiles {
		extension := filepath.Ext(f)
		icon, iconExists := icons[extension]
		if !iconExists {
			icon = green("")
		}
		fmt.Printf("%s %s\n", icon, f)
	}
}
