package main

import (
	"errors"
	"fmt"
	"os"
)

/*
getRootDir function takes command-line arguments and
checks if there is a necessary one - path of a root directory.
If user doesn't specify any params or specifies two or more, function returns error.
*/
func getRootDir() (string, error) {
	var rootDir string
	argsWithoutProg := os.Args[1:]
	switch len(argsWithoutProg) {
	case 0:
		path, err := os.Getwd()
		if err != nil {
			err := fmt.Sprintf("can't get current dir: %s", err.Error())
			return rootDir, errors.New(err)
		}
		rootDir = path
	case 1:
		rootDir = argsWithoutProg[0]
	default:
		return rootDir, errors.New("invalid cli parameters")
	}
	return rootDir, nil
}

/*
getCliHelp function prints to console info how to use this program.
*/
func getCliHelp() string {
	var help string
	help += fmt.Sprintln("Usage:")
	help += fmt.Sprintln("\t./tree.exe\t- current dir will be root dir")
	help += fmt.Sprintln("\t./tree.exe path\t- path will be root dir")
	return help
}
