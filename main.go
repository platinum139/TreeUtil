package main

import (
	"github.com/fatih/color"
	"os"
)

func main() {
	rootDir, err := getRootDir()
	if err != nil {
		color.Red(err.Error())
		color.Green(getCliHelp())
		os.Exit(1)
	}
	walkDirectory(rootDir, "")
}
