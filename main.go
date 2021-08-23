package main

import (
	"fmt"
	"log"
	"os"
	"royallist/prints"

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
	path := "."

	if c.Args().Len() == 1 {
		path = c.Args().First()
	}

	if c.Bool("recursive") {
		fmt.Println(path)
		prints.RecursivePrint(path)
		return nil
	}

	if c.Args().Len() > 1 {
		for _, path := range c.Args().Slice() {
			fmt.Printf("\n%s\n------------\n\n", path)
			prints.CommonPrint(path)
		}
		return nil
	}

	prints.CommonPrint(path)
	return nil
}
