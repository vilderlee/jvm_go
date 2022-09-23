package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	abs, err := filepath.Abs(path)
	if err != nil {
		Error(fmt.Sprintf("parse path error,error:%v", err))
	}
	return &DirEntry{abs}
}

func (e *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(e.absDir, className)
	file, err := os.ReadFile(fileName)
	if err != nil {
		Error(fmt.Sprintf("os read file error,error:%v", err))
		return nil, nil, err
	}
	return file, e, nil
}

func (e *DirEntry) String() string {
	return e.absDir
}
