package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type WildCardEntry struct {
	absDir string
}

func newWildCardEntry(path string) *WildCardEntry {
	abs, err := filepath.Abs(path)
	if err != nil {
		Error(fmt.Sprintf("parse path error,error:%v", err))
	}
	return &WildCardEntry{abs}
}

func (e *WildCardEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(e.absDir, className)
	file, err := os.ReadFile(fileName)
	if err != nil {
		Error(fmt.Sprintf("os read file error,error:%v", err))
	}
	return file, e, nil
}

func (e *WildCardEntry) String() string {
	return e.absDir
}
