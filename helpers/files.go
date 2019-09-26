package helpers

import (
	"io/ioutil"
	"log"
)

func ReadAllFiles(root string) []string {
	var filesList []string
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		filesList = append(filesList, f.Name())
	}
	return filesList
}
