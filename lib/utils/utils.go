package utils

import (
	"fmt"
	"log"
	"os"
)

const SessionFilePath = ".session"

func CheckError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err.Error())
	}
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func CreateFileIfNotExists(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
	CheckError(err, fmt.Sprintf("Error opening file: %s", filePath))

	return file
}
