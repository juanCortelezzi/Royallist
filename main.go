package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"royallist/icons"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "royallist",
		Usage: "Colorfull and simple ls",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Usage:   "recursive view of the directory",
				Aliases: []string{"r"},
			},
		},
		Action: royallist,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func royallist(c *cli.Context) error {
	if c.Bool("recursive") {
		fmt.Println(c.Args().First())
		recursiveIconPrint(c.Args().First(), 0)
		return nil
	}

	if c.Args().Len() == 1 {
		iconPrint(c.Args().First())
		return nil
	}

	if c.Args().Len() > 1 {
		for _, path := range c.Args().Slice() {
			fmt.Printf("\n%s\n------------\n\n", path)
			iconPrint(path)
		}
		return nil
	}

	iconPrint("./")

	return nil
}

func getPathContent(path string) []fs.DirEntry {
	contents, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	return contents
}

func recursiveIconPrint(path string, level int) {
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
		recursiveIconPrint(path+"/"+dir, level+1)
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

func iconPrint(path string) {
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
