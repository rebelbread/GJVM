package classpath

import (
	"path/filepath"
	"archive/zip"
	"errors"
	"io/ioutil"
)

type ZipEntry struct {
	absPath string
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		panic(err)
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rf, err := f.Open()
			if err != nil {
				panic(err)
			}
			defer rf.Close()
			data, err := ioutil.ReadAll(rf)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found : " + className)
}
func (self *ZipEntry) String() string {
	return self.absPath
}
func newZipEntry(path string) (*ZipEntry) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath: absPath}
}
