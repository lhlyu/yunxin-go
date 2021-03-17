package main

import (
	"io/ioutil"
	"os"
	"path"
)

func clean(dir string) {
	rd, _ := ioutil.ReadDir(dir)
	for _, fi := range rd {
		if fi.IsDir() {
			continue
		} else {
			os.RemoveAll(path.Join(dir, fi.Name()))
		}
	}
}
