package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func zipDir(dir string) {
	var files []string

	errFileWalk := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	if errFileWalk != nil {
		os.Exit(2)
		fmt.Println("From ArchiveFilesTest: " + errFileWalk.Error())
	}

	archive, err := os.Create("archive.zip")
	if err != nil {
		os.Exit(3)
		fmt.Println("From ArchiveFilesTest: " + err.Error())
	}
	defer archive.Close()
	zipWriting := zip.NewWriter(archive)

	for _, file := range files {

		archivedFile, errOfOpeningFile := os.Open(file)
		if errOfOpeningFile != nil {
			os.Exit(4)
			fmt.Println("From ArchiveFilesTest: " + errOfOpeningFile.Error())
		}

		fileInfo, _ := os.Stat(file)

		if !fileInfo.IsDir() {
			w1, err := zipWriting.Create(file)
			if err != nil {
				os.Exit(5)
				fmt.Println("From ArchiveFilesTest: " + err.Error())
			}
			if _, err := io.Copy(w1, archivedFile); err != nil {
				os.Exit(6)
				fmt.Println("From ArchiveFilesTest: " + err.Error())
			}
		} else {

		}

		archivedFile.Close()
		os.Remove(file)
	}
	zipWriting.Close()
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("invalid number of arguments")
		os.Exit(1)
	}

}
