package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"path/filepath"
)

type ZipEntry struct {
	absDir string
}

func newZipEntry(path string) *ZipEntry {
	abs, err := filepath.Abs(path)
	if err != nil {
		Error(fmt.Sprintf("parse path error,error:%v", err))
	}
	return &ZipEntry{abs}
}

func (e *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(e.absDir)
	if err != nil {
		Error(fmt.Sprintf("zip open read file error,className:%s, dir:%s, error:%v", className, e.absDir, err))
		return nil, nil, err
	}

	defer func() {
		innerErr := reader.Close()
		if innerErr != nil {
			Error(fmt.Sprintf("zip reader close error,className:%s, dir:%s, error:%v", className, e.absDir, innerErr))
		}
	}()

	for _, file := range reader.File {
		if className == file.Name {
			rc, innerErr := file.Open()
			if innerErr != nil {
				Error(fmt.Sprintf("file close error,className:%s, dir:%s, error:%v", className, e.absDir, innerErr))
				return nil, nil, innerErr
			}
			defer func() {
				innerErr2 := reader.Close()
				if innerErr2 != nil {
					Error(fmt.Sprintf("zip reader close error,className:%s, dir:%s, error:%v", className, e.absDir, innerErr2))
				}
			}()

			bytes, innerErr := io.ReadAll(rc)
			if innerErr != nil {
				Error(fmt.Sprintf("io readAll error,className:%s, dir:%s, error:%v", className, e.absDir, innerErr))
				return nil, nil, innerErr
			}
			return bytes, e, nil
		}
	}

	return nil, nil, errors.New(fmt.Sprintf("cannot find this class, className:%s", className))
}

func (e *ZipEntry) String() string {
	return e.absDir
}
