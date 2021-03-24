package main

import (
	"fmt"
	"github.com/gookit/color"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	files, err := ioutil.ReadDir("./")

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
	blue := color.FgBlue.Render
	white := color.FgWhite.Render

	for _, d := range filDirectories {
		fmt.Printf("%s %s\n", cyan(""), d)
	}

	for _, f := range filFiles {
		extension := filepath.Ext(f)
		switch extension {
		case ".ts":
			fmt.Printf("%s %s\n", blue(""), f)
		case ".tsx":
			fmt.Printf("%s %s\n", blue(""), f)
		case ".json":
			fmt.Printf("%s %s\n", yellow(""), f)
		case ".go":
			fmt.Printf("%s %s\n", blue("ﳑ"), f)
		case ".mod":
			fmt.Printf("%s %s\n", white(""), f)
		case ".sum":
			fmt.Printf("%s %s\n", white(""), f)
		case ".js":
			fmt.Printf("%s %s\n", yellow(""), f)
		case ".jsx":
			fmt.Printf("%s %s\n", yellow(""), f)
		case ".css":
			fmt.Printf("%s %s\n", blue(""), f)
		case ".html":
			fmt.Printf("%s %s\n", red(""), f)
		default:
			fmt.Printf("%s %s\n", green(""), f)
		}
	}
}
