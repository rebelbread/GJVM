package classpath

import (
	"strings"
	"errors"
)

type CompositeEntry []Entry

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, classEntry, err := entry.readClass(className)
		if err == nil {
			return data, classEntry, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
func newCompositeEntry(path string) CompositeEntry {
	paths := strings.Split(path, pathListSeparator)
	var entryArr []Entry
	for _, entryPath := range paths {
		entry := newEntry(entryPath)
		entryArr = append(entryArr, entry)
	}
	return entryArr
}
