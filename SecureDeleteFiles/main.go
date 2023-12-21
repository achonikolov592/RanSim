package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func SecureDelete(filename string) {
	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(5)
	}

	stats, err := file.Stat()

	if err != nil {
		fmt.Println(err)
		os.Exit(6)
	}

	sizeOfFile := stats.Size()
	chunk := int64(1 * (1 << 18))

	numberOfParts := (sizeOfFile / chunk) + 1
	position := int64(0)

	for i := int64(0); i < numberOfParts; i++ {
		var howMuchToChange int64
		if sizeOfFile-(i+1)*chunk < 0 {
			howMuchToChange = sizeOfFile - i*chunk
		} else {
			howMuchToChange = chunk
		}

		zeros := make([]byte, howMuchToChange)

		_, err = file.WriteAt([]byte(zeros), position)
		if err != nil {
			fmt.Println(err)
			os.Exit(7)
		}

		position += howMuchToChange
	}

	file.Close()

	if err = os.Remove(stats.Name()); err != nil {
		fmt.Println(err)
		os.Exit(8)
	}

}

func deleteFilesInDir(dir string) {
	var filesInDir []string
	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if path != os.Args[1] {
			filesInDir = append(filesInDir, path)
		}
		return nil
	})

	for _, file := range filesInDir {
		info, err := os.Stat(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		if !(info.IsDir()) {
			SecureDelete(file)
		} else {
			deleteFilesInDir(file)
		}
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	err = os.Remove(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("invalid number of arguments")
		os.Exit(1)
	}

	fmt.Println("Starting test: SecureDeleteFiles")

	deleteFilesInDir(os.Args[1])
}
