package main

import (
	"RRA/SecureDeleteFiles/SecureDeleteFile"
	"archive/zip"
	"helpers"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

var nameOfLogFile string

func zipDir(dir string) {
	var files []string

	errFileWalk := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if errFileWalk != nil {
		os.Exit(2)
		helpers.WriteLog(nameOfLogFile, "Error: "+errFileWalk.Error())
	}

	archive, err := os.Create("archive.zip")
	if err != nil {
		os.Exit(3)
		helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
	}
	defer archive.Close()
	zipWriting := zip.NewWriter(archive)

	for _, file := range files {
		archivedFile, errOfOpeningFile := os.Open(file)
		if errOfOpeningFile != nil {
			os.Exit(4)
			helpers.WriteLog(nameOfLogFile, "Error: "+errOfOpeningFile.Error())
		}

		fileInfo, _ := os.Stat(file)

		if !fileInfo.IsDir() {
			w1, err := zipWriting.Create(file)
			if err != nil {
				os.Exit(5)
				helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
			}
			if _, err := io.Copy(w1, archivedFile); err != nil {
				os.Exit(6)
				helpers.WriteLog(nameOfLogFile, "Error: "+err.Error())
			}
		}

		archivedFile.Close()
		SecureDeleteFile.SecureDelete(file, nameOfLogFile)
	}

	zipWriting.Close()

	os.RemoveAll(dir)
}
func main() {
	nameOfLogFile := helpers.CreateLogFileIfItDoesNotExist("./", "archive")
	helpers.CreateTestFiles("./", nameOfLogFile)

	helpers.WriteLog(nameOfLogFile, "Starting test: ArchiveFiles")
	_, err := os.Stat("./archive.zip")
	if err == nil {
		os.Remove("./archive.zip")
	}
	zipDir("./testfiles")
	helpers.WriteLog(nameOfLogFile, "Ending test: ArchiveFiles")
}

//
