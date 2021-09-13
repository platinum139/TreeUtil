package main

import (
	"path/filepath"
)

type Type string
const (
	RegFile Type = "file"
	Directory Type = "directory"
)

type File struct {
	Path string			`json:"-"`
	Name string			`json:"name"`
	Type Type			`json:"type"`
	Children []*File	`json:"children,omitempty"`
}

func NewFile(path string, fileType Type) *File {
	return &File{
		Path:	  path,
		Name:     filepath.Base(path),
		Type:     fileType,
		Children: []*File{},
	}
}

func (f *File) Add(file *File) {
	f.Children = append(f.Children, file)
}
