package main

import (
	"RRA/EncryptDecryptDirRecursive/decrypt"
	"RRA/EncryptDecryptDirRecursive/encrypt"
	"helpers"
	"os"
	"strings"
)

var nameOfLogFile string

func main() {
<<<<<<< HEAD
	nameOfLogFile = helpers.CreateLogFileIfItDoesNotExist("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursive/", "EncryptDecryptDirRecursive")

	helpers.WriteLog(nameOfLogFile, "Strating test: EncryptDecryptDir", 2)

	if strings.ToLower(os.Args[1]) == "true" {
		helpers.RemoveTestFilesIfExists("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursive/")
		helpers.CreateTestFiles("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursive/", nameOfLogFile)
		nameOfEncryptionInfoFile := helpers.CreateLogFileIfItDoesNotExist("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursive/", "EncryptionInfo")

		helpers.WriteLog(nameOfLogFile, "Starting encryption", 2)

		encrypt.EncryptDir("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursive/testfiles", nameOfLogFile, nameOfEncryptionInfoFile)

		helpers.WriteLog(nameOfLogFile, "Ending encryption", 2)
=======
	nameOfLogFile = helpers.CreateLogFileIfItDoesNotExist("./", "EncryptDirRecursive")

	helpers.WriteLog(nameOfLogFile, "Strating test: EncryptDecryptDir")

	if strings.ToLower(os.Args[1]) == "true" {
		helpers.RemoveTestFilesIfExists("./")
		helpers.CreateTestFiles("./", nameOfLogFile)
		nameOfEncryptionInfoFile := helpers.CreateLogFileIfItDoesNotExist("./", "EncryptionInfo")

		helpers.WriteLog(nameOfLogFile, "Starting encryption")

		encrypt.EncryptDir("./testfiles", nameOfLogFile, nameOfEncryptionInfoFile)

		helpers.WriteLog(nameOfLogFile, "Ending encryption")
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
	}

	if strings.ToLower(os.Args[2]) == "true" {

<<<<<<< HEAD
		helpers.WriteLog(nameOfLogFile, "Starting encryption", 2)

		decrypt.DecryptDir("C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursive/testfiles", "C:/Users/achon/OneDrive/Desktop/diplomna1/RRA/EncryptDecryptDirRecursive/EncryptionInfo.log", nameOfLogFile)

		helpers.WriteLog(nameOfLogFile, "Ending encryption", 2)
	}

	helpers.WriteLog(nameOfLogFile, "Endinging test: EncryptDecryptDir", 2)
	os.Exit(0)
=======
		helpers.WriteLog(nameOfLogFile, "Starting encryption")

		decrypt.DecryptDir("./testfiles", "./EncryptionInfo.log")

		helpers.WriteLog(nameOfLogFile, "Ending encryption")
	}

	helpers.WriteLog(nameOfLogFile, "Endinging test: EncryptDecryptDir")
>>>>>>> 7fc80a6796b63203067b2e2e7540726a00337350
}

//
