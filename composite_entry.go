package main

import (
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(listPath string) CompositeEntry {
	compositeEntry := make([]Entry, 0)
	for _, path := range strings.Split(listPath, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}

	return compositeEntry
}

func (e CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range e {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, err
		}
	}
	return nil, nil, nil
}

func (e CompositeEntry) String() string {
	strs := make([]string, 0)
	for _, entry := range e {
		strs = append(strs, entry.String())
	}
	return strings.Join(strs, pathListSeparator)
}
