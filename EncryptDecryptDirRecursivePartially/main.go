package main

import (
	"RRA/EncryptDecryptDirRecursivePartially/decrypt"
	"RRA/EncryptDecryptDirRecursivePartially/encrypt"
	"RRA/SecureDeleteFiles/SecureDeleteFile"
	"helpers"
	"os"
	"strings"
)

var nameOfLogFile string

func main() {
	nameOfLogFile = helpers.CreateLogFileIfItDoesNotExist("./", "EncryptDirRecursivePartially")
	helpers.CreateLogFileIfItDoesNotExist("./", "EncryptionInfo")

	if strings.ToLower(os.Args[1]) == "true" {

		_, err := os.Stat("./EncryptionInfo.log")
		if err == nil {
			SecureDeleteFile.SecureDelete("./EncryptionInfo.log", nameOfLogFile)
		}

		nameOfInfoFile := helpers.CreateLogFileIfItDoesNotExist("./", "EncryptionInfo")
		helpers.RemoveTestFilesIfExists("./")
		helpers.CreateTestFiles("./", nameOfLogFile)

		helpers.WriteLog(nameOfLogFile, "Strating test: EncryptDirPartially")

		encrypt.EncryptFilesInDir("./testfiles", nameOfLogFile, nameOfInfoFile)

		helpers.WriteLog(nameOfLogFile, "Ending test: EncryptDirRecursivePartially")
	}
	if strings.ToLower(os.Args[2]) == "true" {
		helpers.WriteLog(nameOfLogFile, "Strating test: DecryptDirRecursivePartially")

		decrypt.DecryptDir("./testfiles", nameOfLogFile, "./EncryptionInfo.log")

		helpers.WriteLog(nameOfLogFile, "Ending test: DecryptDirRecursivePartially")
	}
}

//
