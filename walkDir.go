package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/fs"
	"io/ioutil"
)

const (
	Regular = "├───"
	Last = "└───"
	Between = "│"
)

/*
walkDirectory function iterates over all files in a root directory and
prints their names. If the file is a directory, it calls itself recursively
with that directory as root.
*/
func walkDirectory(rootDir string, indent string) {
	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		color.Red(err.Error())
		return
	}
	dirs := getOnlyDirs(files)
	regFiles := getOnlyRegularFiles(files)

	var indentSymbol string
	for i, dir := range dirs {
		if i == len(regFiles) - 1 && len(regFiles) == 0 {
			indentSymbol = Last
		} else {
			indentSymbol = Regular
		}
		fmt.Print(indent + indentSymbol)
		color.Cyan(dir.Name())
		walkDirectory(fmt.Sprintf("%s\\%s", rootDir, dir.Name()), indent + Between + "    ")
	}
	for i, file := range regFiles {
		if i == len(regFiles) - 1 {
			indentSymbol = Last
		} else {
			indentSymbol = Regular
		}
		fmt.Print(indent + indentSymbol)
		color.Magenta(file.Name())
	}
}

/*
getOnlyDirs function iterates over all files in a slice and
appends only directories to a result slice.
*/
func getOnlyDirs(files []fs.FileInfo) []fs.FileInfo {
	var dirs []fs.FileInfo
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file)
		}
	}
	return dirs
}

/*
getOnlyRegularFiles function iterates over all files in a slice and
appends only regular files to a result slice.
*/
func getOnlyRegularFiles(files []fs.FileInfo) []fs.FileInfo {
	var regFiles []fs.FileInfo
	for _, file := range files {
		if !file.IsDir() {
			regFiles = append(regFiles, file)
		}
	}
	return regFiles
}
