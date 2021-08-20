package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"royallist/icons"
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

func printContentsFromPath(path string) {
	contents, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

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

func royallist(c *cli.Context) error {
	if c.Bool("recursive") {
		fmt.Println("recursive not supported yet")
		return nil
	}

	if c.Args().Len() == 1 {
		printContentsFromPath(c.Args().First())
		return nil
	}

	if c.Args().Len() > 1 {
		for _, path := range c.Args().Slice() {
			fmt.Printf("\n%s\n------------\n\n", path)
			printContentsFromPath(path)
		}
		return nil
	}

	printContentsFromPath("./")

	return nil
}
