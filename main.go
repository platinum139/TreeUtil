package main

import (
	"fmt"
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
	fmt.Println(rootDir)
	// walkDirectory(rootDir, "")

	tree := NewFileTree(rootDir)
	err = tree.Build()
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
	json, err := tree.ToJSON()
	if err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
	color.Cyan(json)
}
