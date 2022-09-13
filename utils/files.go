package utils

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"

	"github.com/gomarkdown/markdown"
)

// DiscoverFiles discovers all markdown files which are within a directory and
// its subdirectories recursively, if successfull, it returns err!=nil and a
// list of paths relative to rootDir
func DiscoverFiles(rootDir string) (error, []string) {
	var files []string
	err := filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(info.Name()) == ".md" {
			files = append(files, path[len(rootDir)+1:])
		}
		return nil
	})
	return err, files
}

// ProcessFile reads the file from the path and converts the markdown content
// into html, if successfull it returns error=nil and the markdown html as an array
// of bytes
func ProcessFile(path string) (error, []byte) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err, file
	}
	return nil, markdown.ToHTML(file, nil, nil)
}
