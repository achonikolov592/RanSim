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
	nameOfLogFile = helpers.CreateLogFileIfItDoesNotExist("./", "EncryptDirRecursive")

	helpers.WriteLog(nameOfLogFile, "Strating test: EncryptDecryptDir")

	if strings.ToLower(os.Args[1]) == "true" {
		helpers.RemoveTestFilesIfExists("./")
		helpers.CreateTestFiles("./", nameOfLogFile)
		nameOfEncryptionInfoFile := helpers.CreateLogFileIfItDoesNotExist("./", "EncryptionInfo")

		helpers.WriteLog(nameOfLogFile, "Starting encryption")

		encrypt.EncryptDir("./testfiles", nameOfLogFile, nameOfEncryptionInfoFile)

		helpers.WriteLog(nameOfLogFile, "Ending encryption")
	}

	if strings.ToLower(os.Args[2]) == "true" {

		helpers.WriteLog(nameOfLogFile, "Starting encryption")

		decrypt.DecryptDir("./testfiles", "./EncryptionInfo.log")

		helpers.WriteLog(nameOfLogFile, "Ending encryption")
	}

	helpers.WriteLog(nameOfLogFile, "Endinging test: EncryptDecryptDir")
}

//
