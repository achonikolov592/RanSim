package main

import (
	"RRA/SecureDeleteFiles/SecureDeleteFile"
	"helpers"
	"os"
	"path/filepath"
)

var nameOfLogFile string

func deleteFilesInDir(dir string) {
	var filesInDir []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if path != dir {
			filesInDir = append(filesInDir, path)
		}
		return nil
	})

	for _, file := range filesInDir {
		info, err := os.Stat(file)
		if err != nil {
			helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
			os.Exit(2)
		}

		if !(info.IsDir()) {
			SecureDeleteFile.SecureDelete(file, nameOfLogFile)
		} else {
			deleteFilesInDir(file)
		}
	}

	if err != nil {
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
		os.Exit(3)
	}

	err = os.Remove(dir)
	if err != nil {
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
		os.Exit(4)
	}
}

func main() {
	nameOfLogFile = helpers.CreateLogFileIfItDoesNotExist("./", "SecureDeleteFiles")
	helpers.CreateTestFiles("./", nameOfLogFile)

	helpers.WriteLog("Starting test: SecureDeleteFiles", nameOfLogFile)

	deleteFilesInDir("./testfiles")

	helpers.WriteLog("Ending test: SecureDeleteFiles", nameOfLogFile)
}

//
