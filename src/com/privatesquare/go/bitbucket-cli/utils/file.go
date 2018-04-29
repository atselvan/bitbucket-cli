package utils

import (
	"log"
	"os"
)

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateFile(name string) os.File {
	var f os.File
	if !FileExists(name) {
		fo, err := os.Create(name)
		Error(err, "There was a error creating a new file")
		defer fo.Close()
	} else {
		log.Println("File created : ", name)
	}
	return f
}

func RemoveFile(name string) {
	if FileExists(name) {
		err := os.Remove(name)
		Error(err, "There was an error deleting the file")
	} else {
		log.Println("File does not exist hence not deleting the file : ", name)
	}
}
