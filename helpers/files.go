package helpers

import (
	"io/ioutil"
	"log"
)

/**
* return with []strings that contains
* files names inside path what you set
*/
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
