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
<<<<<<< HEAD
			helpers.WriteLog(nameOfLogFile, err.Error(), 1)
=======
			helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
			os.Exit(2)
		}

		if !(info.IsDir()) {
			SecureDeleteFile.SecureDelete(file, nameOfLogFile)
		} else {
			deleteFilesInDir(file)
		}
	}

	if err != nil {
<<<<<<< HEAD
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
=======
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
		os.Exit(3)
	}

	err = os.Remove(dir)
	if err != nil {
<<<<<<< HEAD
		helpers.WriteLog(nameOfLogFile, err.Error(), 1)
=======
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
		os.Exit(4)
	}
}

func main() {
<<<<<<< HEAD
	nameOfLogFile = helpers.CreateLogFileIfItDoesNotExist("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/SecureDeleteFiles/", "SecureDeleteFiles")
	helpers.CreateTestFiles("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/SecureDeleteFiles/", nameOfLogFile)

	helpers.WriteLog(nameOfLogFile, "Starting test: SecureDeleteFiles", 2)

	deleteFilesInDir("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/SecureDeleteFiles/testfiles")

	helpers.WriteLog(nameOfLogFile, "Ending test: SecureDeleteFiles", 2)

	os.Exit(0)
=======
	nameOfLogFile = helpers.CreateLogFileIfItDoesNotExist("./", "SecureDeleteFiles")
	helpers.CreateTestFiles("./", nameOfLogFile)

	helpers.WriteLog("Starting test: SecureDeleteFiles", nameOfLogFile)

	deleteFilesInDir("./testfiles")

	helpers.WriteLog("Ending test: SecureDeleteFiles", nameOfLogFile)
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
}

//
