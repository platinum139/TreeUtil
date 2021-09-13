package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type FileTree struct {
	root *File
}

func NewFileTree(root string) *FileTree {
	return &FileTree{root: &File{
		Path:	  root,
		Name:     filepath.Base(root),
		Type:     Directory,
	}}
}

func (tree *FileTree) Build() error {
	return tree.walkDir(tree.root)
}

func (tree *FileTree) walkDir(root *File) error {
	files, err := ioutil.ReadDir(root.Path)
	if err != nil {
		return err
	}
	dirs := getOnlyDirs(files)
	regFiles := getOnlyRegularFiles(files)

	for _, dir := range dirs {
		file := NewFile(root.Path + "\\" + dir.Name(), Directory)
		root.Add(file)
		err = tree.walkDir(file)
		if err != nil {
			return err
		}
	}
	for _, file := range regFiles {
		file := NewFile(root.Path + "\\" + file.Name(), RegFile)
		root.Add(file)
	}
	return nil
}

func (tree *FileTree) ToJSON() (string, error) {
	res, err := json.MarshalIndent(tree.root.Children, "", "    ")
	if err != nil {
		return "", err
	}
	return string(res), nil
}

