package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func clean(dir string) {
	rd, _ := ioutil.ReadDir(dir)
	for _, fi := range rd {
		if fi.IsDir() {
			continue
		} else {
			if isApiFile(fi.Name()) {
				os.RemoveAll(path.Join(dir, fi.Name()))
			}
			if fi.Name() == "README.md" {
				os.RemoveAll(path.Join(dir, fi.Name()))
			}
		}
	}
}

func isApiFile(file string) bool {
	return strings.Index(file, "api") == 0
}
