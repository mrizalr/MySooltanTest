package utils

import (
	"log"
	"os"
)

func ReadLogFile(logFileDir string) string {
	bytes, err := os.ReadFile(logFileDir)
	if err != nil {
		log.Fatalf("Error when read input file | err : %v", err.Error())
	}

	return string(bytes)
}

func CheckDirExists(logFileDir string) error {
	_, err := os.Stat(logFileDir)
	return err
}
