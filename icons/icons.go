package icons

import (
	"github.com/gookit/color"
)

var green = color.FgGreen.Render
var red = color.FgRed.Render
var yellow = color.FgYellow.Render
var cyan = color.FgCyan.Render
var magenta = color.FgMagenta.Render
var blue = color.FgBlue.Render
var white = color.FgWhite.Render

var Filetypes = map[string]string{
	".h":    magenta(""),
	".c":    magenta(""),
	".jpg":  magenta(""),
	".png":  magenta(""),
	".vim":  green(""),
	".zip":  green(""),
	".scss": red(""),
	".html": red(""),
	".pdf":  red(""),
	".json": yellow(""),
	".yml":  yellow(""),
	".yaml": yellow(""),
	".js":   yellow(""),
	".jsx":  yellow(""),
	".mp4":  blue(""),
	".cpp":  blue(""),
	".lua":  blue(""),
	".py":   blue(""),
	".go":   blue("ﳑ"),
	".css":  blue(""),
	".ts":   blue(""),
	".tsx":  blue(""),
	".sh":   white(""),
	".txt":  white(""),
	".mk":   white(""),
	".rs":   white(""),
	".iso":  white(""),
	".md":   white(""),
	".mdx":  white(""),
	".wiki": white(""),
	".mod":  white(""),
	".sum":  white(""),
	".toml": white(""),
}

var Names = map[string]string{
	"directory":  cyan(""),
	"file":       green(""),
	"LICENSE":    white("ﲘ"),
	"Makefile":   white(""),
	".gitignore": magenta(""),
}
